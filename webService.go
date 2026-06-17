package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	const apiPath = "/menu"
	http.HandleFunc(apiPath, Handler)

	// start the web server
	//http.ListenAndServe("localhost:3000")
	http.ListenAndServe(":3000", nil) // localhost can be omitted
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// todo: error handling

	const filePath = "./menu.txt"
	file, _ := os.Open(filePath) // ignore the error for now

	fmt.Printf("Read the menu file from file \"%v\".  \n", filePath)

	io.Copy(w, file)
}
