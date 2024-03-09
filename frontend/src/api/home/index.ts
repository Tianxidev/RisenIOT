import request from "/@/utils/request";


// 首页概览
export function OverviewApi(){
  return request({
    url: '/api/v1/home/overview',
    method: 'get',
  });
}
