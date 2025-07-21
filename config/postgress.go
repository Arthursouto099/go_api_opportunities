package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dataBaseLog Logger
)

func InitializeDatabaseConfig() *gorm.DB {

	dataBaseLog = *GetLogger("dbConnection")
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		envs["DB_HOST"],
		envs["DB_USER"],
		envs["DB_PASSWORD"],
		envs["DB_NAME"],
		envs["DB_PORT"],
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dataBaseLog.Error(err)
	}

	return db
}

// func InitializeDatabaseConfig() {

// 	dataBaseLog = *GetLogger("dbConnection")

// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")
// 	port := os.Getenv("DB_PORT")

// 	//  get connection to database
// 	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s", user, password, dbName, port)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		dataBaseLog.ErrorF("Error connecting to database: %v", err)
// 		return
// 	}

// 	DB = db
// 	dataBaseLog.InfoF("Successful database connection!")
// }
