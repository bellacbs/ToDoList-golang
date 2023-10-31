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

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	UserId      uuid.UUID `json:"userId"`
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

func (task *Task) CheckDates() []string {
	erros := []string{}
	currentDate := time.Now()
	if currentDate.After(task.StartDate) {
		erros = append(erros, "Current Date is over then startDate")
	}
	if task.EndDate.Before(task.StartDate) {
		erros = append(erros, "StartDate is over then EndDate")
	}
	if len(erros) == 0 {
		return nil
	}
	return erros
}
