package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user-center/biz/dal/model"
)

var RDSHook *gorm.DB

func InitDB() {
	dsn := "root:Wenrh240004@tcp(127.0.0.1:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}

	RDSHook = db
}
