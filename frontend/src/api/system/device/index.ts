import request from '/@/utils/request';
import {appendQuery} from "/@/utils/url";

/**
 * 设备列表
 * @method params 设备列表参数
 * @return Promise
 */
export function DeviceInfoList(params: any): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/info/list', params),
        method: 'get',
    });
}

/**
 * 添加设备信息
 * @method params 添加设备参数
 * @return Promise
 */
export function DeviceInfoAdd(params: any): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/info/add', {}),
        method: 'post',
        data: params,
    })
}

/**
 * 编辑设备信息
 * @method params 编辑设备参数
 * @return Promise
 */
export function DeviceInfoEdit(params: any): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/info/edit', {}),
        method: 'put',
        data: params,
    })
}

/**
 * 删除设备信息
 * @param ids 设备ID
 * @return Promise
 */
export function DeviceInfoDel(ids: number[]): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/info/delete', {
            ids: ids,
        }),
        method: 'delete',
    })
}

/**
 * 设备产品类型列表
 * @method params 产品类型列表查询参数
 * @return Promise
 */
export function DeviceKindList(params?: any): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/kind/list', params),
        method: 'get',
    });
}

/**
 * 设备产品类型添加
 * @param params 产品类型信息
 * @constructor
 */
export function DeviceKindAdd(params?: any): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/kind/add', {}),
        method: 'post',
        data: params,
    });
}

/**
 * 设备产品类型删除
 * @param ids 产品类型ID列表
 * @constructor
 */
export function DeviceKindDel(ids: number[]): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/kind/delete', {
            ids: ids,
        }),
        method: 'delete',
    })
}


/**
 * 获取设备分组列表
 * @method params 分组列表分页参数
 */
export function DeviceGroupList(params?: any): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/group/list', params),
        method: 'get',
    })
}

/**
 * 删除设备分组
 * @param ids 设备分组 ID 列表
 * @return Promise
 */
export function DeviceGroupDel(ids: number[]): Promise<any> {
    return request({
        url: appendQuery('/api/v1/device/group/delete', {
            ids: ids,
        }),
        method: 'delete',
    })
}

/**
 * 添加设备分组
 * @param params 设备分组参数
 * @return Promise
 */
export function DeviceGroupAdd(params?:any): Promise<any>{
    return request({
        url: appendQuery('/api/v1/device/group/add', {}),
        method: 'post',
        data: params,
    })
}

/**
 * 编辑设备分组
 * @param params 编辑设备分组参数
 * @return Promise
 */
export function DeviceGroupEdit(params?:any): Promise<any>{
    return request({
        url: appendQuery('/api/v1/device/group/edit', {}),
        method: 'put',
        data: params,
    })
}

/**
 * 设备数据类型列表
 * @param params 数据类型列表查询参数
 * @return Promise
 */
export function DeviceCategoryList(params?:any): Promise<any>{
    return request({
        url: appendQuery('/api/v1/device/category/list', params),
        method: 'get',
    })
}

