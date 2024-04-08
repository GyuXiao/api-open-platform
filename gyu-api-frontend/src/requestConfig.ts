import type { RequestOptions } from '@@/plugin-request/request';
import type { RequestConfig } from '@umijs/max';
import {getToken, setToken} from "@/tools/token";
import {message} from "antd";

// 与后端约定的响应数据格式
interface ResponseStructure {
  data?: any;
}

/**
 * @name 错误处理
 * pro 自带的错误处理， 可以在这里做自己的改动
 * @doc https://umijs.org/docs/max/request#配置
 */
export const requestConfig: RequestConfig = {
  baseURL:'http://localhost:8081' ,

  // 请求拦截器
  requestInterceptors: [
    (config: RequestOptions) => {
      // 拦截请求配置，进行个性化处理。
      const authHeader = { Authorization: 'Bearer ' + getToken() };
      const url = config?.url;
      return { ...config, url, headers: authHeader};
    },
  ],

  // 响应拦截器
  responseInterceptors: [
    (response) => {
      // 拦截响应数据，进行个性化处理
      const { data } = response as unknown as ResponseStructure;
      // 打印响应数据用于调试
      console.log('data', data?.data);
      // todo: 发现问题：因为每次请求都会触发 current 请求，current 请求成功后，会覆盖掉其他有问题的请求
      // 当响应的状态码不为 200，打印错误
      if (data?.code !== 200) {
        message.error('请求失败: '+data?.msg)
        // throw new Error('请求失败: '+msg);
      }
      if (data.token) {
        setToken(data.token)
      }
      return response;
    },
  ],
};
