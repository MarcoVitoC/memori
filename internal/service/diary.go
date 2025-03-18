package service

import (
	"encoding/json"
	"net/http"
	"time"
)

type CreateDiaryPayload struct {
	ID 				string		`json:"id"`
	Content 	string		`json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DiaryService struct {}

func (s *DiaryService) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (s *DiaryService) Create(w http.ResponseWriter, r *http.Request) {
	var payload CreateDiaryPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Failed to read JSON request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}