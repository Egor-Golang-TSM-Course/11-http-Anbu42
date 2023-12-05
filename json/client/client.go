package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	newUser1 := User{Name: "Bob", Age: 25}
	newUser2 := User{Name: "Sam", Age: 30}
	usersList := []User{newUser1, newUser2}
	for _, v := range usersList {
		err := createUser(v)
		if err != nil {
			fmt.Println("Error creating user:", err)
			return
		}
	}

	users, err := getUsers()
	if err != nil {
		fmt.Println("Error getting users:", err)
		return
	}

	fmt.Println("List of users:")
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
	}
}

func createUser(user User) error {
	url := "http://localhost:8080/user"

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("POST Response:", resp.Status)
	fmt.Println(string(body))

	return nil
}

func getUsers() ([]User, error) {
	url := "http://localhost:8080/user"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("GET Response:", resp.Status)
	fmt.Println(string(body))

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
