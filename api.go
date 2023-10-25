package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIServer struct {
	Router     *gin.Engine
	repository Repository
}

func (api *APIServer) SetupRoutes() {

	usersGroup := api.Router.Group("/users")
	usersGroup.POST("/signup", api.createUser)
	usersGroup.POST("/login", api.loginUser)
}

func (api *APIServer) createUser(c *gin.Context) {
	var user User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal Error Try again later")
	}
	userId := uuid.New()
	user.ID = userId
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal Error Try again later")
	}
	user.Password = hashPassword

	err = api.repository.RegisterUser(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, "Internal Error Try again later")
	}

	c.JSON(http.StatusOK, user)
}

func (api *APIServer) loginUser(c *gin.Context) {
	var userLoginDTO UserLoginDTO
	err := json.NewDecoder(c.Request.Body).Decode(&userLoginDTO)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal Error Try again later")
	}
	c.JSON(http.StatusOK, "login")
}
