import request from '/@/utils/request';

// 查询列表
export function listSysUserOnline(query: object) {
  return request({
    url: '/api/v1/system/online/list',
    method: 'get',
    params: query
  })
}

// 强退用户
export function forceLogout(ids: number[]) {
  return request({
    url: '/api/v1/system/online/forceLogout',
    method: 'delete',
    data: {ids}
  })
}

// 强退所有用户
export function forceLogoutAll() {
  return request({
    url: '/api/v1/system/online/forceLogoutAll',
    method: 'delete'
  })
}
