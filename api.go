package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIServer struct {
	Router *gin.Engine
}

func (api *APIServer) SetupRoutes() {

	usersGroup := api.Router.Group("/users")
	usersGroup.POST("/signup", api.createUser)
}

func (api *APIServer) createUser(c *gin.Context) {
	var user User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	userId := uuid.New()
	user.ID = userId
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		fmt.Println(err.Error())
	}
	user.Password = hashPassword

	c.JSON(http.StatusOK, user)
}
