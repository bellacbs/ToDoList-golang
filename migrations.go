package main

import (
	"database/sql"
)

func Migrations(db *sql.DB) {
	createUserTable(db)
	createTaskTable(db)
}

func createUserTable(db *sql.DB) {
	create, err := db.Query("CREATE TABLE IF NOT EXISTS Users (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), password varchar(255));")
	if err != nil {
		panic(err.Error())
	}
	defer create.Close()
}

func createTaskTable(db *sql.DB) {
	create, err := db.Query("CREATE TABLE IF NOT EXISTS Tasks (id varchar(255) PRIMARY KEY, title varchar(255), description varchar(255), startDate varchar(255), endDate varchar(255), userId varchar(255), FOREIGN KEY (userId) REFERENCES Users(id));")
	if err != nil {
		panic(err.Error())
	}
	defer create.Close()
}
