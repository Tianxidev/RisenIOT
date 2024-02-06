// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"backend/api/v1/common"
	"backend/api/v1/device"
	"backend/api/v1/system"
)

type IV1Common interface {
	Captcha(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error)
}

type IV1Device interface {
	DataGet(ctx context.Context, req *device.DataGetReq) (res *device.DataGetRes, err error)
	DataAdd(ctx context.Context, req *device.DataAddReq) (res *device.DataAddRes, err error)
	InfoSearch(ctx context.Context, req *device.InfoSearchReq) (res *device.InfoSearchRes, err error)
	InfoGet(ctx context.Context, req *device.InfoGetReq) (res *device.InfoGetRes, err error)
	InfoAdd(ctx context.Context, req *device.InfoAddReq) (res *device.InfoAddRes, err error)
	InfoEdit(ctx context.Context, req *device.InfoEditReq) (res *device.InfoEditRes, err error)
	InfoDelete(ctx context.Context, req *device.InfoDeleteReq) (res *device.InfoDeleteRes, err error)
	KindSearch(ctx context.Context, req *device.KindSearchReq) (res *device.KindSearchRes, err error)
}

type IV1System interface {
	CacheRemove(ctx context.Context, req *system.CacheRemoveReq) (res *system.CacheRemoveRes, err error)
	PersonalInfo(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error)
	PersonalEdit(ctx context.Context, req *system.PersonalEditReq) (res *system.PersonalEditRes, err error)
	PersonalResetPwd(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error)
	RuleSearch(ctx context.Context, req *system.RuleSearchReq) (res *system.RuleSearchRes, err error)
	RuleAdd(ctx context.Context, req *system.RuleAddReq) (res *system.RuleAddRes, err error)
	RuleGetParams(ctx context.Context, req *system.RuleGetParamsReq) (res *system.RuleGetParamsRes, err error)
	RuleInfo(ctx context.Context, req *system.RuleInfoReq) (res *system.RuleInfoRes, err error)
	RuleUpdate(ctx context.Context, req *system.RuleUpdateReq) (res *system.RuleUpdateRes, err error)
	RuleDelete(ctx context.Context, req *system.RuleDeleteReq) (res *system.RuleDeleteRes, err error)
	ConfigSearch(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error)
	ConfigAdd(ctx context.Context, req *system.ConfigAddReq) (res *system.ConfigAddRes, err error)
	ConfigGet(ctx context.Context, req *system.ConfigGetReq) (res *system.ConfigGetRes, err error)
	ConfigEdit(ctx context.Context, req *system.ConfigEditReq) (res *system.ConfigEditRes, err error)
	ConfigDelete(ctx context.Context, req *system.ConfigDeleteReq) (res *system.ConfigDeleteRes, err error)
	DeptSearch(ctx context.Context, req *system.DeptSearchReq) (res *system.DeptSearchRes, err error)
	DeptAdd(ctx context.Context, req *system.DeptAddReq) (res *system.DeptAddRes, err error)
	DeptEdit(ctx context.Context, req *system.DeptEditReq) (res *system.DeptEditRes, err error)
	DeptDelete(ctx context.Context, req *system.DeptDeleteReq) (res *system.DeptDeleteRes, err error)
	DeptTreeSelect(ctx context.Context, req *system.DeptTreeSelectReq) (res *system.DeptTreeSelectRes, err error)
	GetDict(ctx context.Context, req *system.GetDictReq) (res *system.GetDictRes, err error)
	DictDataSearch(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error)
	DictDataAdd(ctx context.Context, req *system.DictDataAddReq) (res *system.DictDataAddRes, err error)
	DictDataGet(ctx context.Context, req *system.DictDataGetReq) (res *system.DictDataGetRes, err error)
	DictDataEdit(ctx context.Context, req *system.DictDataEditReq) (res *system.DictDataEditRes, err error)
	DictDataDelete(ctx context.Context, req *system.DictDataDeleteReq) (res *system.DictDataDeleteRes, err error)
	DictTypeSearch(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error)
	DictTypeAdd(ctx context.Context, req *system.DictTypeAddReq) (res *system.DictTypeAddRes, err error)
	DictTypeGet(ctx context.Context, req *system.DictTypeGetReq) (res *system.DictTypeGetRes, err error)
	DictTypeEdit(ctx context.Context, req *system.DictTypeEditReq) (res *system.DictTypeEditRes, err error)
	DictTypeDelete(ctx context.Context, req *system.DictTypeDeleteReq) (res *system.DictTypeDeleteRes, err error)
	DictTypeAll(ctx context.Context, req *system.DictTypeAllReq) (res *system.DictTypeAllRes, err error)
	DbInitIsInit(ctx context.Context, req *system.DbInitIsInitReq) (res *system.DbInitIsInitRes, err error)
	DbInitGetEnvInfo(ctx context.Context, req *system.DbInitGetEnvInfoReq) (res *system.DbInitGetEnvInfoRes, err error)
	DbInitCreateDb(ctx context.Context, req *system.DbInitCreateDbReq) (res *system.DbInitCreateDbRes, err error)
	UserLogin(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error)
	LoginLogSearch(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error)
	LoginLogDel(ctx context.Context, req *system.LoginLogDelReq) (res *system.LoginLogDelRes, err error)
	LoginLogClear(ctx context.Context, req *system.LoginLogClearReq) (res *system.LoginLogClearRes, err error)
	UserLoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error)
	MonitorServer(ctx context.Context, req *system.MonitorServerReq) (res *system.MonitorServerRes, err error)
	SysOperLogSearch(ctx context.Context, req *system.SysOperLogSearchReq) (res *system.SysOperLogSearchRes, err error)
	SysOperLogGet(ctx context.Context, req *system.SysOperLogGetReq) (res *system.SysOperLogGetRes, err error)
	SysOperLogDelete(ctx context.Context, req *system.SysOperLogDeleteReq) (res *system.SysOperLogDeleteRes, err error)
	SysOperLogClear(ctx context.Context, req *system.SysOperLogClearReq) (res *system.SysOperLogClearRes, err error)
	PostSearch(ctx context.Context, req *system.PostSearchReq) (res *system.PostSearchRes, err error)
	PostAdd(ctx context.Context, req *system.PostAddReq) (res *system.PostAddRes, err error)
	PostEdit(ctx context.Context, req *system.PostEditReq) (res *system.PostEditRes, err error)
	PostDelete(ctx context.Context, req *system.PostDeleteReq) (res *system.PostDeleteRes, err error)
	RoleList(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error)
	RoleGetParams(ctx context.Context, req *system.RoleGetParamsReq) (res *system.RoleGetParamsRes, err error)
	RoleAdd(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error)
	RoleGet(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error)
	RoleEdit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error)
	RoleDelete(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error)
	UserSearch(ctx context.Context, req *system.UserSearchReq) (res *system.UserSearchRes, err error)
	UserGetParams(ctx context.Context, req *system.UserGetParamsReq) (res *system.UserGetParamsRes, err error)
	UserAdd(ctx context.Context, req *system.UserAddReq) (res *system.UserAddRes, err error)
	UserEdit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error)
	UserGetEdit(ctx context.Context, req *system.UserGetEditReq) (res *system.UserGetEditRes, err error)
	UserResetPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error)
	UserStatus(ctx context.Context, req *system.UserStatusReq) (res *system.UserStatusRes, err error)
	UserDelete(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error)
	UserGetByIds(ctx context.Context, req *system.UserGetByIdsReq) (res *system.UserGetByIdsRes, err error)
	UserMenus(ctx context.Context, req *system.UserMenusReq) (res *system.UserMenusRes, err error)
	SysUserOnlineSearch(ctx context.Context, req *system.SysUserOnlineSearchReq) (res *system.SysUserOnlineSearchRes, err error)
	SysUserOnlineForceLogout(ctx context.Context, req *system.SysUserOnlineForceLogoutReq) (res *system.SysUserOnlineForceLogoutRes, err error)
}
