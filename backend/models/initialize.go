package models

import (
	"RisenIOT/backend/common/global"
	"RisenIOT/backend/pkg/logger"
)

func Setup() {
	var err error

	// 检查表 sys_user 是否存在
	if !global.Eloquent.Migrator().HasTable(&SysUser{}) {

		// 创建表
		err = global.Eloquent.AutoMigrate(&SysUser{})
		if err != nil {
			panic(err)
		}

		// 创建管理员用户数据
		_, err = SysUser{
			LoginM: LoginM{
				SysUsername: SysUsername{
					Username: "admin",
				},
				SysPassword: SysPassword{
					Password: "123456",
				},
			},
		}.Insert()
		if err != nil {
			logger.GlobalLogger.ERROR("插入管理员用户数据异常: %v", err)
		}
	}

	// 检查表 casbin_rule 是否存在
	if !global.Eloquent.Migrator().HasTable(&CasbinRule{}) {

		// 创建表
		err = global.Eloquent.AutoMigrate(&CasbinRule{})
		if err != nil {
			panic(err)
		}

		// 插入 root 角色的访问控制规则
		_, err = CasbinRule{
			PType: "p",
			V0:    "root",
			V1:    "*",
			V2:    "*",
		}.Insert()
		if err != nil {
			logger.GlobalLogger.ERROR("插入 root 角色的访问控制规则异常: %v", err)
		}

	}

	logger.GlobalLogger.INFO("初始化数据表完成")

}
