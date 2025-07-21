package config

import "gorm.io/gorm"

var (
	logger *Logger
	DB     *gorm.DB
)

// return error
func Init() error {
	// nulo -> não existe erro
	DB = InitializeDatabaseConfig()
	return nil
}

func GetLogger(p string) *Logger {
	// initialize Logger
	logger = NewLogger(p)
	return logger
}

// retorna o endereço do gorm DB
func GetPostgresDB() *gorm.DB {
	return DB
}
