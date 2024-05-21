package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"sandvik.dev/goApp/internal/types"
)

func HandleSecret(w http.ResponseWriter, r *http.Request) {
	tmplPath := "files/templates/main.gohtml"
	tmpl, err := template.New("main.gohtml").ParseFiles(tmplPath)
	if err != nil {
		fmt.Fprintf(w, "Could not find parse gohtml\nERROR: %s", err.Error())
		return
	}
	pass := os.Getenv("PASSWORD")
	if pass == "" {
		pass = "UNDEFINED"
	}
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "UNDEFINED"
	}
	passData := fmt.Sprintf("Your Password is: %s", pass)
	data := types.TmplData{
		Title:   "Super Secret",
		Data:    passData,
		PodName: podName,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Could not find parse gohtml\nERROR: %s", err.Error())
		return
	}
}

func HandleLorem(w http.ResponseWriter, r *http.Request) {
	tmplPath := "files/templates/main.gohtml"
	tmpl, err := template.New("main.gohtml").ParseFiles(tmplPath)
	if err != nil {
		fmt.Fprintf(w, "Could not find parse gohtml\nERROR: %s", err.Error())
		return
	}
	lorem_byte, err := os.ReadFile("./files/custom.config")
	lorem := string(lorem_byte)
	if err != nil {
		fmt.Printf("Could not open file ERR: %s", err.Error())
		lorem = "MAIN TEXT UNDEFINED"
	}
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "UNDEFINED"
	}
	data := types.TmplData{
		Title:   "All the text",
		Data:    lorem,
		PodName: podName,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Could not find parse gohtml\nERROR: %s", err.Error())
		return
	}
}

func ReportHealth(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("POD_NAME")
	if name == "" {
		name = "UNDEFINED"
	}
	health := os.Getenv("HEALTH")
	if health != "DEAD" {
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("ok PODNAME: %s", name)))
	} else {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Dead PODNAME: %s", name)))
	}
}

func SetHealth(w http.ResponseWriter, r *http.Request) {
	pass := os.Getenv("PASSWORD")

	authHeader := r.Header.Get("Authorization")
	if authHeader != pass {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print("Could not read Body")
	}
	os.Setenv("HEALTH", string(body))
}
