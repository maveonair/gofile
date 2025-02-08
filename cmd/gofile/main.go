package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/file/", generateFileHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateFileHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/file/")
	size, err := strconv.Atoi(path)
	if err != nil || size < 1 || size > 1000 {
		http.Error(w, "Invalid size parameter. Must be between 1 and 1000.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%dMB_file.txt", size))
	w.Header().Set("Content-Type", "application/octet-stream")

	fileSize := size * 1024 * 1024  // Convert MB to bytes
	data := make([]byte, 1024*1024) // 1MB buffer

	for written := 0; written < fileSize; written += len(data) {
		if fileSize-written < len(data) {
			data = data[:fileSize-written] // Adjust the last chunk size
		}
		if _, err := w.Write(data); err != nil {
			http.Error(w, "Failed to write data", http.StatusInternalServerError)
			return
		}
	}
}
