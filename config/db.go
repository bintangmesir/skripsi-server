package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DatabaseConnection(DB_USER string, DB_PASSWORD string, DB_URI string, DB_NAME string) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_URI, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Tidak bisa menghubungkan pada database...")
	}

	// if err := db.AutoMigrate(&models.LaporanDanaSantunanAnakAsuh{}, &models.DanaSantunanAnakAsuh{}); err != nil {
	// 	panic(err)
	// }

	DB = db
}
