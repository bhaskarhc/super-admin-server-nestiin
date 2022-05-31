package db

import (
	"fmt"

	"github.com/bhaskarhc/admin-nestiin/modules"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func InitDB(cfg *viper.Viper) error {

	username := cfg.GetString("postgresConfig.DB_USER")
	password := cfg.GetString("postgresConfig.DB_PASS")
	dbName := cfg.GetString("postgresConfig.DB_NAME")
	dbHost := cfg.GetString("postgresConfig.DB_HOST")
	dbPort := cfg.GetString("postgresConfig.DB_PORT")

	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, username, password, dbName, dbPort)
	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
		return err
	}

	db = conn
	db.Debug().AutoMigrate(&modules.UserData{}, &modules.SellarInformationReq{}, &modules.QualityDetails{}, &modules.TechnicalRequirement{})
	return nil
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
