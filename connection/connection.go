package connection

import (
	"log"
	"test-backend-dev/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/user"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed Connect Database")
	}

	db.AutoMigrate(&entity.Nation{}, &entity.Customer{}, &entity.Family_list{})

	return db
}
