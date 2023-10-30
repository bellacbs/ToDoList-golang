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

	tasksGroup := api.Router.Group("/tasks")
	tasksGroup.Use(api.requireTokenAuthorization())
	tasksGroup.POST("/create", api.createTask)
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
		return
	}
	missingFields := userLoginDTO.CheckEmptyKeyAndValue()
	if missingFields != nil {
		c.JSON(http.StatusConflict, HandleError{Code: http.StatusConflict, Message: missingFields})
		return
	}
	user, err := api.repository.GetUserByEmail(userLoginDTO.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{"Internal Error Try again later"}})
		return
	}
	if user == nil {
		c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{"User doesn't exist"}})
		return
	}
	err = CheckPasswordHash(user.Password, userLoginDTO.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusUnauthorized, Message: []string{"Incorrect Password"}})
		return
	}
	userResponse := UserResponseDTO{ID: user.ID, Email: user.Email, Name: user.Name}
	token, err := GenerateToken(&userResponse)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{"Error To generate  token try again later"}})
		return
	}
	userResponse.Token = token
	c.JSON(http.StatusOK, userResponse)
}

func (api *APIServer) createTask(c *gin.Context) {
	var createTaskDTO CreateTaskDTO
	err := json.NewDecoder(c.Request.Body).Decode(&createTaskDTO)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal Error Try again later")
		return
	}
	missingFields := createTaskDTO.CheckEmptyKeyAndValue()
	if missingFields != nil {
		c.JSON(http.StatusConflict, HandleError{Code: http.StatusConflict, Message: missingFields})
		return
	}

	fmt.Println(createTaskDTO)
}
func (api *APIServer) requireTokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("authorization")
		if tokenString == "" {
			c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusUnauthorized, Message: []string{"Invalid Token"}})
			c.Abort()
			return
		}
		user, err := CheckToken(tokenString)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{err.Error()}})
		}
		userDB, err := api.repository.GetUserByEmail(user.Email)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{"Error To generate  token try again later"}})
			c.Abort()
			return
		}
		if userDB == nil {
			c.JSON(http.StatusInternalServerError, HandleError{Code: http.StatusInternalServerError, Message: []string{"User doesn't exist"}})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
