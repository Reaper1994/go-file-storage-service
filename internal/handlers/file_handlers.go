package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Reaper1994/go-file-storage-service/internal/parser"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (h *FileHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form to retrieve the file
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file locally
	tempFile, err := os.Create(filepath.Join(os.TempDir(), "uploaded_file.xml"))
	if err != nil {
		http.Error(w, "Unable to create temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// Parse the XML file and save to the database asynchronously
	go func() {
		err = parser.ParseAndSaveXML(tempFile.Name())
		if err != nil {
			log.Printf("Error parsing and saving XML: %v\n", err)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("File uploaded successfully"))
}
