package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataBaseUrl := os.Getenv("DATA_BASE_URL")
	driverName := os.Getenv("DRIVER_NAME")
	port := os.Getenv("PORT")
	api := APIServer{}
	api.Router = gin.Default()
	db, err := sql.Open(driverName, dataBaseUrl)
	if err != nil {
		panic(err.Error())
	}
	api.Db = db
	defer db.Close()
	Migrations(api.Db)
	api.SetupRoutes()
	api.Router.Run(fmt.Sprintf(":%s", port))
}
