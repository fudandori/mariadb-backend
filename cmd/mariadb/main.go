package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"fudandori/pkg/db"
	"fudandori/pkg/utils"
)

func main() {
	props, err := utils.ReadPropertiesFile("../../config/db.properties")
	if err != nil {
		log.Printf("Error while reading properties file")
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		props["user"],
		props["pass"],
		props["host"],
		props["port"],
		props["schema"])

	datasource, err := sql.Open("mysql", connection)

	if err != nil {
		log.Print(err.Error())
	}

	router := gin.Default()

	router.GET("/places", func(context *gin.Context) {
		response := db.GetPlaces(datasource)
		context.JSON(http.StatusOK, response)
	})

	router.GET("/earn", func(context *gin.Context) {
		response := db.GetEarnings(datasource)
		context.JSON(http.StatusOK, response)
	})

	router.GET("/clients", func(context *gin.Context) {
		response := db.GetClients(datasource)
		context.JSON(http.StatusOK, response)
	})

	router.GET("/items", func(context *gin.Context) {
		response := db.GetItems(datasource)
		context.JSON(http.StatusOK, response)
	})

	router.Run(":8080")
}
