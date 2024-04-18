// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost struct {
	PostId    uint64      `json:"postId"    orm:"post_id"    ` // 岗位ID
	PostCode  string      `json:"postCode"  orm:"post_code"  ` // 岗位编码
	PostName  string      `json:"postName"  orm:"post_name"  ` // 岗位名称
	PostSort  int         `json:"postSort"  orm:"post_sort"  ` // 显示顺序
	Status    uint        `json:"status"    orm:"status"     ` // 状态（0正常 1停用）
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // 创建人
	UpdatedBy uint64      `json:"updatedBy" orm:"updated_by" ` // 修改人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // 删除时间
}
