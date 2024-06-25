package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/upload"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/upload.html"))
		if err := tmpl.Execute(w, struct{ Title string }{Title: "Upload"}); err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set a reasonable file size limit (e.g., 10 MB)
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from form: %v", err)
		http.Error(w, "Failed to get file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileURL, err := upload.UploadToDigitalOcean(file, header)
	if err != nil {
		log.Printf("Error uploading file to DigitalOcean: %v", err)
		http.Error(w, "Failed to upload file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully. URL: " + fileURL))
}