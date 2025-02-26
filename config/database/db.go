package database

import (
	"crud_fiber.com/m/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB
var err error

func connectDatabase() {
	appConfig := config.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DbHostPsql, appConfig.DbPortPsql, appConfig.DbUserNamePsql, appConfig.DbPasswordPsql, appConfig.DbPsql)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // This instructs GORM to not pluralize table names
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panicln("error occurred while trying to connect database arguments: ", err.Error())
		return
	}

}

func InitializeDatabase() {
	connectDatabase()
}

func GetInstanceDatabase() *gorm.DB {
	return db
}
