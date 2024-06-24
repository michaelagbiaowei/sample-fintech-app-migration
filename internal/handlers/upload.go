package handlers

import (
	"html/template"
	"net/http"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/upload"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/upload.html"))
		tmpl.Execute(w, struct{ Title string }{Title: "Upload"})
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileURL, err := upload.UploadToDigitalOcean(file, header)
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully. URL: " + fileURL))
}