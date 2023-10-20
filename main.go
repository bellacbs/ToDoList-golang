package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	api := APIServer{}
	api.Router = gin.Default()
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/todolist")
	if err != nil {
		panic(err.Error())
	}
	api.Db = db
	defer db.Close()
	Migrations(api.Db)
	api.SetupRoutes()
	api.Router.Run(":3000")
}
