package service

import "net/http"

type Service struct {
	Diary interface {
		GetAll(w http.ResponseWriter, r *http.Request)
		Create(w http.ResponseWriter, r *http.Request)
	}
}

func NewService() *Service {
	return &Service{
		Diary: &DiaryService{},
	}
}