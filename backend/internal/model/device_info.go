package model

import "backend/internal/model/entity"

type DeviceAllInfo struct {
	Info         *entity.SysDeviceInfo       `json:"info"`
	Kind         *entity.SysDeviceKind       `json:"kind"`
	CategoryList []*entity.SysDeviceCategory `json:"categoryList"`
}
