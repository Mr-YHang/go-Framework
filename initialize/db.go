package initialize

import (
	"fmt"
	"go-Framework/config"
	"go-Framework/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func InitializeDB() {
	// 根据驱动配置进行初始化
	global.App.DB = initMySqlGorm(global.App.Config.OrderDB)
}

func CloseDB() {
	if global.App.DB != nil {
		db, _ := global.App.DB.DB()
		db.Close()
	}
}

// 初始化 mysql gorm.DB
func initMySqlGorm(dbConfig config.DBConf) *gorm.DB {
	if dbConfig.Database == "" {
		panic(fmt.Errorf("initMySqlGorm -- 初始化数据库错误 \n"))
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN: dsn, // DSN data source name
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		//Logger:                                   getGormLogger(), // 使用自定义 Logger
	}); err != nil {
		panic(fmt.Errorf("initMySqlGorm -- 初始化mysql错误 : %s \n", err.Error()))
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		return db
	}
}
