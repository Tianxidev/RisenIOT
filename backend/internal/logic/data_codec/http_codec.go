package dataCodec

import (
	"backend/api/v1/device"
	"backend/internal/consts"
	"backend/internal/model"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// HttpDecode http解码
func (s *sDataCodec) HttpDecode(ctx context.Context, dataContent interface{}) (decodeMsg *model.DeviceDecodeMsg, err error) {
	var msgCon *device.DataAddReq
	decodeMsg = &model.DeviceDecodeMsg{}

	liberr.ValueIsNil(ctx, dataContent, "解析设备数据为空")
	liberr.ErrIsNil(ctx, g.Try(ctx, func(ctx context.Context) {
		msgCon, _ = dataContent.(*device.DataAddReq)
	}), "interface to type err")

	g.Log().Print(ctx, "msgCon:", msgCon)

	// 获取设备信息
	decodeMsg.DeviceInfo, err = service.DeviceInfo().Info(ctx, msgCon.DeviceId, msgCon.DeviceSn)
	liberr.ErrIsNil(ctx, err, "获取设备信息异常")
	liberr.ValueIsNil(ctx, decodeMsg.DeviceInfo, "不存在该设备")

	// 判断上报数据类型
	if msgCon.Property != nil && len(decodeMsg.DeviceInfo.CategoryList) > 0 {
		jsonContent := gjson.New(msgCon.Property)
		g.Log().Print(ctx, "反序列化设备数据:", jsonContent)
		for _, category := range decodeMsg.DeviceInfo.CategoryList {
			var dtime *gtime.Time

			if nil == dtime {
				dtime = gtime.NewFromStr(msgCon.Time)
			}

			if nil == dtime {
				dtime = gtime.Now()
			}

			// 构建设备数据模型
			decodeMsg.DataList = append(decodeMsg.DataList, &model.DeviceData{
				CategoryId: category.Id,
				Name:       category.Mark,
				Type:       category.DataType,
				Ratio:      category.Ratio,
				Data:       jsonContent.Get(category.Mark),
				Time:       dtime,
			})
		}

	}

	// 判断上报事件不为空
	if msgCon.Event != nil {
		jsonContent := gconv.MapStrStr(msgCon.Event)
		g.Log().Print(ctx, "反序列化设备事件:", jsonContent)
		for key, value := range jsonContent {
			prevent := &model.DeviceEvent{
				Name: key,
				Data: value,
			}
			decodeMsg.EventList = append(decodeMsg.EventList, prevent)
		}
	}
	return decodeMsg, nil
}

// Save 保存数据
func (s *sDataCodec) Save(ctx context.Context, decodeMsg *model.DeviceDecodeMsg) (err error) {
	if nil == decodeMsg {
		return gerror.New("param is null")
	}
	req := &device.CategoryDataAddReq{}
	for _, data := range decodeMsg.DataList {
		req.CategoryId = data.CategoryId
		req.DeviceId = decodeMsg.DeviceInfo.Info.Id
		req.CreatedAt = data.Time
		switch data.Type {
		case consts.CategoryDataTypeBit:
			fallthrough
		case consts.CategoryDataTypeByte:
			fallthrough
		case consts.CategoryDataTypeShort:
			fallthrough
		case consts.CategoryDataTypeUnShort:
			fallthrough
		case consts.CategoryDataTypeInt:
			fallthrough
		case consts.CategoryDataTypeUnInt:
			req.DataInt = gconv.Uint(data.Data)
			if data.Ratio != "" && len(data.Ratio) > 0 {
				req.DataInt = gconv.Uint(gconv.Float64(req.DataInt) * gconv.Float64(data.Ratio))
			}
			g.Log().Print(ctx, "save int", req.DataDouble, data.Data)
		case consts.CategoryDataTypeFloat:
			fallthrough
		case consts.CategoryDataTypeDouble:
			req.DataDouble = gconv.Float64(data.Data)

			if data.Ratio != "" && len(data.Ratio) > 0 {
				req.DataDouble = gconv.Float64(req.DataDouble) * gconv.Float64(data.Ratio)
			}
			g.Log().Print(ctx, "save float", req.DataDouble, data.Data)
		default:
			req.DataStr = gconv.String(data.Data)
			g.Log().Print(ctx, "save string", req.DataDouble, data.Data)
		}

		// 写入数据库
		err = service.DeviceCategoryData().Add(ctx, req)
		liberr.ErrIsNil(ctx, err, "save device data err")
	}

	// 更新设备状态
	err = service.DeviceStatus().ChangeStatus(ctx, decodeMsg.DeviceInfo.Info.Id, consts.DeviceStatusOnLine)
	liberr.ErrIsNil(ctx, err, "更新设备状态失败")

	// 更新设备数据最后上报时间
	err = service.DeviceInfo().UpdateDataLastTime(ctx, decodeMsg.DeviceInfo.Info.Id)
	liberr.ErrIsNil(ctx, err, "更新设备最后上报时间失败")
	return
}
