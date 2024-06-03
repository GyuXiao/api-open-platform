import { request } from '@umijs/max';

/** listTopInvokeInterfaceInfo GET /api/analysis/top/interface/invoke */
export async function listTopInvokeInterfaceInfoUsingGET(
  params: { limit?: number },
  options?: { [key: string]: any }
) {
  return request<API.BaseResponseTopNInvokeInterfaceInfo>('/gyu_api/v1/analysis/top/interfaceInfo', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
