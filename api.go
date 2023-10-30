package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	var createUserDTO CreateUserDTO
	err := json.NewDecoder(c.Request.Body).Decode(&createUserDTO)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal Error Try again later")
		return
	}
	missingFields := createUserDTO.CheckEmptyKeyAndValue()
	if missingFields != nil {
		c.JSON(http.StatusConflict, HandleError{Code: http.StatusConflict, Message: missingFields})
		return
	}
	userId := uuid.New()
	userResponse := UserResponseDTO{ID: userId, Email: createUserDTO.Email, Name: createUserDTO.Name}
	token, err := GenerateToken(&userResponse)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{"Error To generate  token try again later"}})
		return
	}
	userResponse.Token = token
	hashPassword, err := HashPassword(createUserDTO.Password)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal Error Try again later")
		return
	}
	user := User{ID: userId, Email: createUserDTO.Email, Name: createUserDTO.Name, Password: hashPassword}
	user.Password = hashPassword
	err = api.repository.RegisterUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, HandleError{Code: http.StatusConflict, Message: []string{"Email have already exist"}})
			return
		}
		c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{"Internal Error Try again later"}})
		return
	}
	c.JSON(http.StatusOK, userResponse)
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
