package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGenerateFileHandler(t *testing.T) {
	tests := []struct {
		size       int
		statusCode int
	}{
		{1, http.StatusOK},
		{10, http.StatusOK},
		{100, http.StatusOK},
		{1000, http.StatusOK},
		{0, http.StatusBadRequest},
		{1001, http.StatusBadRequest},
		{-1, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run("Size "+strconv.Itoa(tt.size), func(t *testing.T) {
			req, err := http.NewRequest("GET", "/file/"+strconv.Itoa(tt.size), nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(generateFileHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.statusCode)
			}

			if tt.statusCode == http.StatusOK {
				expectedSize := tt.size * 1024 * 1024
				if rr.Header().Get("Content-Length") != strconv.Itoa(expectedSize) {
					t.Errorf("handler returned wrong file size: got %v want %v",
						rr.Header().Get("Content-Length"), expectedSize)
				}
			}
		})
	}
}
