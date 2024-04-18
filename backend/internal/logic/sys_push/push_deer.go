package sysPush

import (
	"backend/api/v1/system"
	"backend/internal/dao"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"io/ioutil"
	"net/http"
)

/**
 * PushDeer 推送
 */

// TextMsgSendPushDeer 文本消息推送到 PushDeer
func (s *sSysPush) TextMsgSendPushDeer(ctx context.Context, key string, text string) bool {
	url := "https://api2.pushdeer.com/message/push?pushkey=" + key + "&text=" + text
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	liberr.ErrIsNil(ctx, err, "消息推送失败")
	res, err := client.Do(req)
	liberr.ErrIsNil(ctx, err, "消息推送失败")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		liberr.ErrIsNil(ctx, err, "消息推送失败")
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	g.Log().Infof(ctx, "PushDeer 推送结果: %s", string(body))
	// 反序列化
	var result map[string]interface{}
	err = gconv.Struct(body, &result)
	liberr.ErrIsNil(ctx, err, "消息推送失败")
	if result["code"] != 0 {
		return false
	}
	return true
}

// SavePushDeerConfig 保存 PushDeer 配置
func (s *sSysPush) SavePushDeerConfig(ctx context.Context, key string) bool {
	// 获取当前用户ID
	uid := int(service.UserCtx().GetUserId(ctx))

	// 获取当前用户表
	m := dao.SysUser.Ctx(ctx)
	mc := dao.SysUser.Columns()

	// 更新用户表
	_, err := m.Where(mc.Id, uid).Update(g.Map{
		mc.Pushdeer: key,
	})
	liberr.ErrIsNil(ctx, err, "保存推送配置失败")
	return true
}

// QueryPushDeerConfig 查询 PushDeer 配置
func (s *sSysPush) QueryPushDeerConfig(ctx context.Context, res *system.PushQueryConfigToPushDeerRes) {
	// 获取当前用户ID
	uid := int(service.UserCtx().GetUserId(ctx))

	// 获取当前用户表
	m := dao.SysUser.Ctx(ctx)
	mc := dao.SysUser.Columns()

	// 查询用户表中的推送配置
	record, err := m.Where(mc.Id, uid).Fields(mc.Pushdeer).One()
	liberr.ErrIsNil(ctx, err, "查询推送配置失败")
	if record.IsEmpty() {
		res.Key = ""
	}
	res.Key = gconv.String(record["pushdeer"].String())
	return

}
