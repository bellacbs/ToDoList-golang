package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Router *gin.Engine
}

func (api *APIServer) SetupRoutes() {

	usersGroup := api.Router.Group("/users")
	usersGroup.POST("/signup", api.createUser)
}

func (api *APIServer) createUser(c *gin.Context) {
	c.String(http.StatusOK, "Create User")
}
