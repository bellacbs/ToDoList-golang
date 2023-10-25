package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := os.Getenv("PORT")
	api := APIServer{}
	api.Router = gin.Default()
	repository, err := NewMySqlStore()
	if err != nil {
		log.Fatal(err)
	}
	api.repository = repository
	// defer db.Close()
	err = repository.Migrations()
	if err != nil {
		log.Fatal(err)
	}
	api.SetupRoutes()
	api.Router.Run(fmt.Sprintf(":%s", port))
}
