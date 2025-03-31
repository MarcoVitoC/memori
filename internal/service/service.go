package service

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/repository"
)

type Service struct {
	Diary interface {
		GetAll(w http.ResponseWriter, r *http.Request)
		GetById(w http.ResponseWriter, r *http.Request)
		Create(w http.ResponseWriter, r *http.Request)
		Update(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
	}
}

func NewService(repo repository.Repository) Service {
	return Service{
		Diary: &DiaryService{repo},
	}
}