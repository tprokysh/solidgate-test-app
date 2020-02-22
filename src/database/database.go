package config

import (
	"encoding/json"
	"fmt"
	"os"

	models "../models"

	"github.com/jinzhu/gorm"
)

type DataBase struct {
	Username string
	Password string
	Host     string
	Database string
	Port     string
}

const configFile = "./etc/db.json"

func getConfig() DataBase {
	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	DBconfiguration := DataBase{}
	err := decoder.Decode(&DBconfiguration)

	if err != nil {
		fmt.Println("error:", err)
	}

	return DBconfiguration
}

func GetConnection() *gorm.DB {
	cfg := getConfig()
	fmt.Println(cfg)
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", cfg.Host, cfg.Username, cfg.Database, cfg.Password, cfg.Port)
	dbConnection, err := gorm.Open("postgres", dbUri)

	if err != nil {
		panic("failed to connect database")
	}

	dbConnection.AutoMigrate(&models.Customer{})

	return dbConnection
}
