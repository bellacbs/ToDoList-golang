package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	api := APIServer{}
	api.Router = gin.Default()

	api.SetupRoutes()

	api.Router.Run(":3000")
}
