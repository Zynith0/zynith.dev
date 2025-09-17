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
	mux.HandleFunc("/yt-mp3", handler.HandleYt)
	mux.HandleFunc("POST /startDownload", handler.HandleDownload)
	mux.HandleFunc("/download", handler.HandleTest)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", mux)
}
