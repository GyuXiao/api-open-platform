package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	nurl "net/url"
)

// 入口函数

func Do(ctx context.Context, method, url string, jsonBody map[string]map[string]any) (*http.Response, error) {
	req, err := buildRequest(ctx, method, url, jsonBody)
	if err != nil {
		return nil, err
	}
	return DoRequest(req)
}

var interceptors = []Interceptor{
	LogInterceptor,
}

type (
	client interface {
		do(r *http.Request) (*http.Response, error)
	}

	defaultClient struct{}
)

func (c defaultClient) do(r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r)
}

// 发送一个 HTTP request 并返回一个 HTTP response

func DoRequest(r *http.Request) (*http.Response, error) {
	return request(r, defaultClient{})
}

// 主要请求逻辑

func request(r *http.Request, cli client) (*http.Response, error) {
	ctx := r.Context()

	tracer := trace.TracerFromContext(ctx)
	propagator := otel.GetTextMapPropagator()

	spanName := r.URL.Path
	ctx, span := tracer.Start(
		ctx,
		spanName,
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		oteltrace.WithAttributes(semconv.HTTPClientAttributesFromHTTPRequest(r)...),
	)
	defer span.End()

	respHandlers := make([]ResponseHandler, len(interceptors))
	for i, interceptor := range interceptors {
		var h ResponseHandler
		// http.Request, responseHandler
		r, h = interceptor(r)
		respHandlers[i] = h
	}

	r = r.WithContext(ctx)
	propagator.Inject(ctx, propagation.HeaderCarrier(r.Header))

	resp, err := cli.do(r)
	for i := len(respHandlers) - 1; i >= 0; i-- {
		respHandlers[i](resp, err)
	}
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(resp.StatusCode)...)
	span.SetStatus(semconv.SpanStatusFromHTTPStatusCodeAndSpanKind(resp.StatusCode, oteltrace.SpanKindClient))

	return resp, err
}

// 通过一个 map[string]map[string]any 类型的数据结构，构建请求

func buildRequest(ctx context.Context, method, url string, body map[string]map[string]any) (*http.Request, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil, err
	}

	var reader io.Reader
	jsonVars, hasJsonBody := body[jsonKey]
	if hasJsonBody {
		if method == http.MethodGet {
			return nil, ErrGetWithBody
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(jsonVars); err != nil {
			return nil, err
		}

		reader = &buf
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), reader)
	if err != nil {
		return nil, err
	}

	if hasJsonBody {
		req.Header.Set(ContentType, JsonContentType)
	}

	return req, nil
}
