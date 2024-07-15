package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// InitMysql 初始化mysql会话
func InitMysql(dbfile string) *gorm.DB {
	// 日志配置
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 彩色打印
		},
	)

	// 初始化会话
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单一表名, eg. `User` => `user`
		},
		DisableForeignKeyConstraintWhenMigrating: true,   // 禁用自动创建外键约束
		Logger:                                   logger, // 自定义Logger
	})
	if err != nil {
		log.Fatal("initMysql gorm.Open err:", err)
	}
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("initMysql db.DB err:", err)
	}
	// 数据库空闲连接池最大值
	sqlDB.SetMaxIdleConns(10)
	// 数据库连接池最大值
	sqlDB.SetMaxOpenConns(100)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(2) * time.Hour)
	//db.Debug().AutoMigrate(&user.UserModel{}, &invit.InvitModel{}, &key.KeyModel{})
	return db
}
