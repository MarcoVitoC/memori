package service

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/auth"
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
	Auth interface {
		Register(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
		Logout(w http.ResponseWriter, r *http.Request)
	}
}

func NewService(repo repository.Repository, authenticator *auth.JWTAuthenticator) Service {
	return Service{
		Diary: &DiaryService{repo},
		Auth: &AuthService{repo, authenticator},
	}
}