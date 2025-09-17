package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var folderName string

type Url struct {
	Url string
	Format string
}

type Name struct {
	Name string
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	renderTemplate("view/index.html", w)
}

func HandleDesc(w http.ResponseWriter, r *http.Request) {
	renderTemplate("view/desc.html", w)
}

func HandleKeyboards(w http.ResponseWriter, r *http.Request) {
	renderTemplate("view/keyboards.html", w)
}

func HandleYt(w http.ResponseWriter, r *http.Request) {
	renderTemplate("view/ytmp3.html", w)
}

func renderTemplate(tmpl string, w http.ResponseWriter) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Println("skill issue", err)
	}
	t.Execute(w, "")
}


func HandleDownload(w http.ResponseWriter, r *http.Request) {
	requestData := Url{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return 
	}

	num1 := rand.Intn(9)
	num2 := rand.Intn(9)
	num3 := rand.Intn(9)
	num4 := rand.Intn(9)

	folderName = strconv.Itoa(num1) + strconv.Itoa(num2) + strconv.Itoa(num3) + strconv.Itoa(num4)
	
	os.Mkdir("/home/zynith/personal/go/personal_website/Videos/" + folderName, 0755)

	cmd := exec.Command("yt-dlp", "-t", requestData.Format, requestData.Url)
	if requestData.Format == "wav" {
		cmd = exec.Command("yt-dlp", "-x", "--audio-format", "wav", requestData.Url)
	}
	cmd.Dir = "/home/zynith/personal/go/personal_website/Videos/" + folderName
	output, err := cmd.Output()
	if err != nil {
		return 
	}
	fmt.Println(string(output))

	cmd2 := exec.Command("zip", "-r", folderName, folderName)
	cmd2.Dir = "/home/zynith/personal/go/personal_website/Videos"
	output2, err := cmd2.Output()
	if err != nil {
		fmt.Println("errer", err)
	}
	fmt.Println(string(output2))
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip") 
	w.Header().Set("Content-Disposition", "attachment; filename=\"Videos.zip\"")
	http.ServeFile(w, r, "/home/zynith/personal/go/personal_website/Videos/" + folderName + ".zip")
}
