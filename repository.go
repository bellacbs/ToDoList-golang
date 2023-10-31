package main

import (
	"database/sql"
	"os"
)

type Repository interface {
	RegisterUser(*User) error
	GetUserByEmail(string) (*User, error)
	CreateTask(task *Task) error
}

type MySqlStore struct {
	db *sql.DB
}

func NewMySqlStore() (*MySqlStore, error) {
	dataBaseUrl := os.Getenv("DATA_BASE_URL")
	driverName := os.Getenv("DRIVER_NAME")
	db, err := sql.Open(driverName, dataBaseUrl)
	if err != nil {
		return nil, err
	}
	return &MySqlStore{db: db}, nil
}

func (mysql *MySqlStore) RegisterUser(user *User) error {
	query := `insert into Users
	(id, name, email, password)
	values(?, ?, ?, ?)`

	_, err := mysql.db.Query(
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password)

	if err != nil {
		return err
	}
	return nil
}

func (mysql *MySqlStore) GetUserByEmail(email string) (*User, error) {
	query := `select * from Users where email=?`
	rows, err := mysql.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}

func (mysql *MySqlStore) CreateTask(task *Task) error {
	query := `insert into Tasks (id, title, description, startDate, endDate, userId) values (?,?,?,?,?,?)`
	_, err := mysql.db.Query(query, task.ID, task.Title, task.Description, task.StartDate, task.EndDate, task.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (mysql *MySqlStore) Migrations() error {
	err := CreateUserTable(mysql.db)
	if err != nil {
		return err
	}
	err = CreateTaskTable(mysql.db)
	if err != nil {
		return err
	}
	return nil
}
