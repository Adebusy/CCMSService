package main

import (
	"github.com/Adebusy/CCMSService/docs"
	"github.com/Adebusy/CCMSService/utilities"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	handleReq "github.com/Adebusy/CCMSService/handlers"
)

func init() {
	godotenv.Load()
	//dbGorm := driver.GetDB()
	//"github.com/Adebusy/CCMSService/features/complaints"
	//"github.com/Adebusy/CCMSService/features/users"

	// fmt.Println("mo connect oooo")
	// //fmt.Println("mo connect ")
	// //fmt.Println("mo connect ")
	// dbGorm.Debug().DropTableIfExists()
	// dbGorm.Debug().DropTableIfExists(&models.TblUsers{})
	// dbGorm.Debug().DropTableIfExists(&models.TblCases{})
	// dbGorm.Debug().DropTableIfExists(&models.TblComplaintChannels{})
	// dbGorm.Debug().DropTableIfExists(&models.TblComplaintSubCategories{})
	// dbGorm.Debug().DropTableIfExists(&models.TblRoles{})
	//dbGorm.Debug().DropTableIfExists(&models.TblComplaintCategories{})
	//dbGorm.Debug().DropTableIfExists(&models.TblComplaintSubCategories{})
	// // dbGorm.SingularTable(true)
	// dbGorm.Debug().AutoMigrate(&models.TblUsers{})
	// dbGorm.Debug().AutoMigrate(&models.TblRoles{})

	// dbGorm.Debug().AutoMigrate(&models.TblCases{})
	// dbGorm.Debug().AutoMigrate(&models.TblComplaintChannels{})
	// dbGorm.Debug().AutoMigrate(&models.TblComplaintSubCategories{})
	//dbGorm.Debug().AutoMigrate(&models.TblComplaintType{})
	//dbGorm.Debug().AutoMigrate(&models.TblComplaintCategories{})
	//dbGorm.Debug().AutoMigrate(&models.TblComplaintSubCategories{})
	// fmt.Println("Done with table creation")
	//}
}

// @title Consumers Complaint Management System API
// @version 1.0
// @description This service is meant to manage consumers complaint API.
// @termsOfService http://swagger.io/terms/
// @contact.name Alao ramon Adebisi
// @contact.email alao.adebusy@gmail.com
// @license.name Linear Logic Concept
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE
// @BasePath /api/v1
func main() {
	AppPort := utilities.GoDotEnvVariable("AppPort")
	docs.SwaggerInfo.Title = "Consumer complaint Management System API"
	docs.SwaggerInfo.Description = "This is a consumers complaints management system API.."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.Host = "localhost" + AppPort
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/Users/CreateUser/", handleReq.CreateUser)
	r.GET("/api/Users/GetUserFullInfoByEmail/:UserEmail", handleReq.GetUserFullInfoByEmail)
	r.POST("/api/Users/ChangeUserDetails/", handleReq.ChangeUserDetails)
	r.GET("/api/Users/FetchAvailableRoles", handleReq.FetchRoles)
	r.GET("/api/Users/ChangeUserStatus/:Username/:Status/:RoleID", handleReq.ChangeUserStatus)
	r.POST("/api/complaint/LogComplaintRequest/", handleReq.LogComplaintRequest)
	r.POST("/api/complaint/CreateComplaintCategories/", handleReq.CreateComplaintCategories)
	r.GET("/api/complaint/FetchComplaintCategories/", handleReq.FetchComplaintCategories)
	r.POST("/api/complaint/CreateComplaintSubCategories/", handleReq.CreateComplaintSubCategories) //done
	r.GET("/api/complaint/FetchComplaintSubCategories/", handleReq.FetchComplaintSubCategories)    //done
	r.GET("/api/complaint/GetComplaintByRefID/:ReferenceID", handleReq.GetComplaintByRefID)
	r.Run(AppPort)
}
