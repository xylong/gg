package db

import (
	"gg/src/model/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Orm *gorm.DB

func InitDB() {
	Orm = gormDB()
	migrate()
}

func migrate() {
	Orm.AutoMigrate(user.NewUser())
}

func gormDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	mysql.SetMaxIdleConns(5)
	mysql.SetMaxOpenConns(10)
	mysql.SetConnMaxLifetime(time.Second * 30)
	return db
}
