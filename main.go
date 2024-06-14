package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
	Pin     json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Roll    string
	Address Address
}

func main() {
	fmt.Println("hello")
	dir := "./"
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error occurred", err)
	}

	users := []User{
		{"Myguy", "20", "95477107XX", "1401", Address{"Varanasi", "Varanasi", "UP", "800421"}},
		{"Myguy", "20", "95477107XX", "1401", Address{"Varanasi", "Varanasi", "UP", "800421"}},
		{"Myguy", "20", "95477107XX", "1401", Address{"Varanasi", "Varanasi", "UP", "800421"}},
	}
	fmt.Println(users)
	// fmt.Println(db)

	for _, val := range users {
		db.Write("users", val.Name, User{
			Name:    val.Name,
			Age:     val.Age,
			Contact: val.Contact,
			Roll:    val.Roll,
			Address: val.Address,
		})
	}

	data, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error occurred ", err)
	}
	fmt.Println(data)

	usersList := []User{}
	for _, f := range data {
		userFound := User{}
		if err := json.Unmarshal([]byte(f), &userFound); err != nil {

			fmt.Println("errorr occurred")
		}
		usersList = append(usersList, userFound)
	}
}
