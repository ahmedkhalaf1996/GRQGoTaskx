package database

import (
	"github.com/ahmedkhalaf1996/GRQGoTaskx/models"
	// "fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CallDB() *gorm.DB {
	// ----- db end ------ //
	// dsn := "postgresql://root:secret@localhost:5432/meetmeup_dev?sslmode=disable"
	// dsn := "postgresql://root:secret@postgres:5432/meetmeup_dev?sslmode=disable"
	dsn := viper.GetString("db.dsn")
	// fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
	// ---- db end --------//

	return db
}
