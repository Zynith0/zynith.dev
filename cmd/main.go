package main

import (
	"fmt"
	"net/http"

	"github.com/Zynith0/personal_website/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleRoot)
	mux.HandleFunc("/AboutMe", handler.HandleDesc)
	mux.HandleFunc("/Keyboards", handler.HandleKeyboards)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", mux)
}
