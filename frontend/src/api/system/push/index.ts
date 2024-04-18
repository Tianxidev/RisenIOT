import request from "/@/utils/request";


// TextMsgToPushDeer 推送文本消息到 PushDeer
export function TextMsgToPushDeer(key: string, text: string): Promise<any> {
  return request({
    url: '/api/v1/system/push/push_deer/text?key=' + key + '&text=' + text,
    method: 'get',
  })
}

// QueryPushDeerConfig 查询 PushDeer 推送配置
export function QueryPushDeerConfig(): Promise<any> {
  return request({
    url: '/api/v1/system/push/push_deer/config',
    method: 'get',
  })
}

// 保存 PushDeer 推送配置
export function SavePushDeerConfig(key: string): Promise<any> {
  return request({
    url: '/api/v1/system/push/push_deer/config',
    method: 'post',
    data: {
      key: key
    }
  })
}