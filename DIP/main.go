package main

import (
	"fmt"
)

// Definition: High-level modules should not depend on low-level modules; both should depend on abstractions.

// Bad code
/*
type MySQLDatabase struct{}

func (db *MySQLDatabase) Connect() {
	fmt.Println("Connecting to MySQL database")
}

// Problem: UserService is tightly coupled with MySQLDatabase.
type UserService struct {
	db MySQLDatabase
}

func (s *UserService) GetUser() {
	s.db.Connect()
	fmt.Println("Getting user")
}
*/

// Low-level module
type Database interface {
	Connect()
}

type MySQLDatabase struct{}

func (m *MySQLDatabase) Connect() {
	fmt.Println("Connected Mysql")
}

type PostgesDatabase struct{}

func (m *PostgesDatabase) Connect() {
	fmt.Println("Connected Postgres")
}

// High-level module
type UserService struct {
	db Database
}

func (s *UserService) GetUser() {
	s.db.Connect()
	fmt.Println("Get user done")
}

func main() {
	us := &UserService{db: new(MySQLDatabase)}
	us.GetUser()

	pDB := new(PostgesDatabase)
	us = &UserService{}
	us.db = pDB
	us.GetUser()
}
