package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {
	var hotDog hotdog
	http.ListenAndServe(":8080", hotDog)
}
