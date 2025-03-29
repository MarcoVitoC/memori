package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MarcoVitoC/memori/internal/repository"
)

type CreateDiaryPayload struct {
	Content string `json:"content"`
}

type DiaryService struct {
	repo repository.Repository
}

func (s *DiaryService) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (s *DiaryService) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload CreateDiaryPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Failed to read JSON request", http.StatusBadRequest)
		return
	}

	newDiary := &repository.Diary{
		Content: payload.Content,
	}

	if err := s.repo.Diary.Create(ctx, newDiary); err != nil {
		log.Fatal("ERROR: failed to create new diary with error ", err)
	}

	w.WriteHeader(http.StatusCreated)
}