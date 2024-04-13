package response

import (
	"encoding/json"
	"fmt"
	"gyu-api-sdk/sdk/errors"
	"io/ioutil"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

type BaseResponse struct {
	ErrorResponse
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) error {
	err := json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	if r.Code > 0 {
		return errors.NewSDKError(r.Code, r.Msg)
	}
	return nil
}

func ParseFromHttpResponse(rawResponse *http.Response, response Response) error {
	defer rawResponse.Body.Close()
	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return err
	}
	if rawResponse.StatusCode != 200 {
		return fmt.Errorf("request fail with status: %s, with body: %s", rawResponse.Status, body)
	}
	err = response.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &response)
}
