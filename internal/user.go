package internal

import (
	"fmt"
	"net/http"
)

func UserHandler(writer http.ResponseWriter, req *http.Request) {
	data := []byte("users endpoint works")

	_, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error")
	}

}

// data := []byte(writer)

// _, err := writer.Write(data)
// if err != nil {
// 	fmt.Println("Oshibka")
// }
