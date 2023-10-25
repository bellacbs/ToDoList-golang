package main

import (
	"database/sql"
)

func CreateUserTable(db *sql.DB) error {
	create, err := db.Query("CREATE TABLE IF NOT EXISTS Users (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), password varchar(255));")
	if err != nil {
		return err
	}
	defer create.Close()
	return nil
}

func CreateTaskTable(db *sql.DB) error {
	create, err := db.Query("CREATE TABLE IF NOT EXISTS Tasks (id varchar(255) PRIMARY KEY, title varchar(255), description varchar(255), startDate DATETIME, endDate DATETIME, userId varchar(255), FOREIGN KEY (userId) REFERENCES Users(id));")
	if err != nil {
		return err
	}
	defer create.Close()
	return nil
}
