package main

import (
	"database/sql"
)

func CreateUserTable(db *sql.DB) error {
	create, err := db.Query("CREATE TABLE IF NOT EXISTS Users (id varchar(255) PRIMARY KEY, name varchar(255) NOT NULL, email varchar(255) UNIQUE, password varchar(255) NOT NULL);")
	if err != nil {
		return err
	}
	defer create.Close()
	return nil
}

func CreateTaskTable(db *sql.DB) error {
	create, err := db.Query("CREATE TABLE IF NOT EXISTS Tasks (id varchar(255) PRIMARY KEY, title varchar(255) NOT NULL, description varchar(255) NOT NULL, startDate DATETIME NOT NULL, endDate DATETIME NOT NULL, userId varchar(255) NOT NULL, FOREIGN KEY (userId) REFERENCES Users(id));")
	if err != nil {
		return err
	}
	defer create.Close()
	return nil
}
