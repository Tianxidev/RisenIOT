package dataCodec

import (
	"backend/api/v1/device"
	"backend/internal/model"
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// HttpDecode http解码
func (s *sDataCodec) HttpDecode(ctx context.Context, dataContent interface{}) (dmesg *model.DeviceDecodeMsg, err error) {
	var msgCon *device.DataAddReq
	if nil == dataContent {
		return nil, gerror.Newf("device parse dataContent is nil, dataContent:%v", dataContent)
	}

	dmesg = &model.DeviceDecodeMsg{}

	// g.Log().Print(ctx, "dataContent:", dataContent)
	msgCon, ok := dataContent.(*device.DataAddReq)

	if !ok {
		g.Log().Printf(ctx, "interface to type err")
		return nil, gerror.Newf("param errror")
	}

	// g.Log().Print(ctx, "msgCon:", msgCon)

	dmesg.DeviceInfo, err = service.DeviceInfo().GetAllInfo(ctx, msgCon.DeviceId, msgCon.DeviceSn)
	if err != nil {
		g.Log().Errorf(ctx, "get deviceinfo err:%v, info:%v", err, dmesg.DeviceInfo)
		return
	}

	if dmesg.DeviceInfo == nil {
		g.Log().Debug(ctx, "get deviceInfo failed")
		return nil, gerror.New("get deviceInfo failed")
	}

	if msgCon.Property != nil && len(dmesg.DeviceInfo.CategoryList) > 0 {
		jsonContent := gjson.New(msgCon.Property)

		// g.Log().Print(ctx, "parse param :", jsonContent)

		for _, category := range dmesg.DeviceInfo.CategoryList {
			// g.Log().Print(ctx, "get param time:", jsonContent.Get("Time"), jsonContent.Get("time"))
			var dtime *gtime.Time

			if nil == dtime {
				// dtime = gtime.Now()
				dtime = gtime.NewFromStr(msgCon.Time)
			}

			if nil == dtime {
				dtime = gtime.Now()
			}

			// dmesg.dataList[index].Time = dtime
			// index = index + 1
			pdata := &model.DeviceData{
				CategoryId: category.Id,
				Name:       category.Mark,
				Type:       category.DataType,
				Ratio:      category.Ratio,
				Data:       jsonContent.Get(category.Mark),
				Time:       dtime,
			}
			dmesg.DataList = append(dmesg.DataList, pdata)
		}

	}

	if msgCon.Event != nil {
		jsonContent := gconv.MapStrStr(msgCon.Event)
		// dmesg.eventList = make([]*DeviceEvent, len(jsonContent))
		for key, value := range jsonContent {
			pevent := &model.DeviceEvent{
				Name: key,
				Data: value,
			}
			dmesg.EventList = append(dmesg.EventList, pevent)
		}
	}

	g.Log().Print(ctx, "encode device indo:", dmesg.DataList, dmesg.EventList)
	return dmesg, nil

}
