package database

import (
	"RisenIOT/backend/common/config"
	"RisenIOT/backend/common/global"
	"RisenIOT/backend/pkg/env"
	"RisenIOT/backend/pkg/logger"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Mysql struct {
}

func (e *Mysql) Setup() {
	db, err := sql.Open("mysql", e.GetConnect())
	if err != nil {
		logger.GlobalLogger.ERROR(e.GetDriver()+" connect error :", err)
	}
	global.Cfg.SetDb(&config.DBConfig{
		Driver: "mysql",
		DB:     db,
	})
	global.Eloquent, err = e.Open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.GlobalLogger.ERROR(e.GetDriver()+" connect error :", err)
	} else {
		logger.GlobalLogger.INFO(e.GetDriver() + " connect success!")
	}

	if global.Eloquent.Error != nil {
		logger.GlobalLogger.ERROR(e.GetDriver()+" connect error :", global.Eloquent.Error)
	}

	//if toolsConfig.LoggerConfig.EnabledDB {
	//	global.Eloquent.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//		SlowThreshold: time.Second,
	//		Colorful:      true,
	//		LogLevel:      logger.Info,
	//	})
	//}

}

// Open 打开数据库连接
func (e *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// GetConnect 获取数据库连接
func (e *Mysql) GetConnect() string {
	host, _ := env.GetEnv("DB_HOST")
	port, _ := env.GetEnv("DB_PORT")
	db, _ := env.GetEnv("DB_DATABASE")
	user, _ := env.GetEnv("DB_USERNAME")
	password, _ := env.GetEnv("DB_PASSWORD")

	return user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
}

// GetDriver 获取数据库驱动
func (e *Mysql) GetDriver() string {
	return "mysql"
}
