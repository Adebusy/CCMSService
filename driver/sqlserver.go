package driver

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
)

var DbGorm *gorm.DB
var ErrGorm error

//ConnectGorm connection string for gorm
func ConnectGorm() error {
	godotenv.Load()
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", GoDotEnvVariable("user"), GoDotEnvVariable("Password"), GoDotEnvVariable("Server"), GoDotEnvVariable("Database"))
	DbGorm, ErrGorm = gorm.Open(GoDotEnvVariable("client"), connectionString)
	if ErrGorm != nil {
		return ErrGorm
	}
	return nil
}

func GetDB() *gorm.DB {
	return DbGorm
}

//GoDotEnvVariable load env file
func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {

		//log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
