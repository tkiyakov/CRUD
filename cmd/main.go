package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func userHandler(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Read Fail")
		return
	}
	_, err = writer.Write(body)

	if err != nil {
		fmt.Println("Write Fail")
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Email)

}

func main() {
	http.HandleFunc("/users", userHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}

}


body, err := io.ReadAll(req.Body)

if err != nil {
	fmt.Println(err)
}

_. err := writer.Write(body)

if err != nil {
	fmt.Println(err)
}

var user User

err = json.Unmarshal(body, &user)

if err != nil [
	fmt.Println(err)
	return
]

fmt.Println(user.Name)

fmt.Println(user.Age)

fmt.Println(user.Email)