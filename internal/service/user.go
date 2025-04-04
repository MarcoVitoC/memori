package service

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/repository"
	"github.com/MarcoVitoC/memori/internal/util"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repo repository.Repository
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (s *UserService) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload RegisterUserPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		util.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	isExist, err := s.repo.User.FindByEmail(ctx, payload.Email)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if isExist {
		util.WriteError(w, http.StatusConflict, "User already exists!")
		return
	}

	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newUser := repository.User{
		Username: payload.Username,
		Email: payload.Email,
		Password: hashedPassword,
	}

	if err := s.repo.User.Register(ctx, &newUser); err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusCreated, nil)
}
