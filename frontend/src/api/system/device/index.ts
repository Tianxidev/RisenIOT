import request from '/@/utils/request';

/**
 * 设备列表
 * @method params 设备列表参数
 */
export function DeviceList(params: object): Promise<any> {
  return request({
    url: '/api/v1/device/info/list',
    method: 'get',
    data: params,
  });
}


/**
 * 设备模型列表
 * @method params 设备模型列表参数
 */
export function DeviceKindList(params: object): Promise<any> {
  return request({
    url: '/api/v1/device/kind/list',
    method: 'get',
    data: params,
  });
}