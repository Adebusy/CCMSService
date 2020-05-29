package main

import (
	"/Users/ramonalao/Projects/github.com/adebusy/CCMSService/docs"

	"github.com/gin-gonic/gin"
)

func main() {

	docs.SwaggerInfo.Title = "Consumer complaint Management System API"
	docs.SwaggerInfo.Description = "This is a consumers complaints management system API.."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.Host = "localhost:8093"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.Default()

}
