// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost struct {
	PostId    uint64      `json:"postId"    ` // 岗位ID
	PostCode  string      `json:"postCode"  ` // 岗位编码
	PostName  string      `json:"postName"  ` // 岗位名称
	PostSort  int         `json:"postSort"  ` // 显示顺序
	Status    uint        `json:"status"    ` // 状态（0正常 1停用）
	Remark    string      `json:"remark"    ` // 备注
	CreatedBy uint64      `json:"createdBy" ` // 创建人
	UpdatedBy uint64      `json:"updatedBy" ` // 修改人
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
}
