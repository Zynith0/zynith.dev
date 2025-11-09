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
	// mux.HandleFunc("/Pomodoro", handler.HandlePomodoro)
	mux.HandleFunc("/yt-mp3", handler.HandleYt)
	mux.HandleFunc("POST /startDownload", handler.HandleDownload)
	mux.HandleFunc("/test", handler.HandleTest)

	fmt.Printf("Server listening on :8080\n")
	http.ListenAndServe(":8080", mux)
}
