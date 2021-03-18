package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gin-blog/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

func Setup() {
	var err error

	dsn := setting.DatabaseSetting.User + ":" + setting.DatabaseSetting.Password + "@tcp(" + setting.DatabaseSetting.Host + ")/" + setting.DatabaseSetting.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseSetting.TablePrefix,
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panicln("db.DB() err: ", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}
