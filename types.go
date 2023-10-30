package main

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDTO struct {
	Email    string
	Password string
}

type UserResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	Token string    `json:"token"`
}

type HandleError struct {
	Code    int      `json:"code"`
	Message []string `json:"message"`
}

type CreateTaskDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

func (createUserDTO *CreateUserDTO) CheckEmptyKeyAndValue() []string {
	valuesObject := reflect.ValueOf(createUserDTO)
	missingFields := CheckEmptyKeyAndValue(valuesObject)
	return missingFields
}

func (userLoginDTO *UserLoginDTO) CheckEmptyKeyAndValue() []string {
	valuesObject := reflect.ValueOf(userLoginDTO)
	missingFields := CheckEmptyKeyAndValue(valuesObject)
	return missingFields
}

func (createTaskDTO *CreateTaskDTO) CheckEmptyKeyAndValue() []string {
	valuesObject := reflect.ValueOf(createTaskDTO)
	missingFields := CheckEmptyKeyAndValue(valuesObject)
	return missingFields
}
