package model

import "backend/internal/model/entity"

type SysDeptTreeRes struct {
	*entity.SysDept
	Children []*SysDeptTreeRes `json:"children"`
}
