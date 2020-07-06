package driver

import (
	"fmt"
	"github.com/Adebusy/CCMSService/utilities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
)

var DbGorm *gorm.DB
var ErrGorm error

func init() {
	//fmt.Println("connect init")
	godotenv.Load()
	//fmt.Println("connect Load")
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", utilities.GoDotEnvVariable("user"), utilities.GoDotEnvVariable("Password"), utilities.GoDotEnvVariable("Server"), utilities.GoDotEnvVariable("Database"))
	DbGorm, ErrGorm = gorm.Open(utilities.GoDotEnvVariable("client"), connectionString)
	if ErrGorm != nil {
		fmt.Println("connect error")
		fmt.Println(ErrGorm.Error())
	}
	//fmt.Println("mo connect....")
	//return nil
}

func GetDB() *gorm.DB {
	//fmt.Println("connect GetDB")
	return DbGorm
}
