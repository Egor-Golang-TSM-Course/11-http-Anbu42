package main

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User
