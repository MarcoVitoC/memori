package service

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/auth"
	"github.com/MarcoVitoC/memori/internal/errors"
	"github.com/MarcoVitoC/memori/internal/repository"
	"github.com/MarcoVitoC/memori/internal/util"
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
		errors.InternalServerError(w, err.Error())
		return
	}

	if errs := util.Validate(payload); len(errs) > 0 {
		errors.BadRequest(w, errs)
		return
	}

	isExist, err := s.repo.User.GetByEmail(ctx, payload.Email)
	if err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	if isExist {
		errors.Conflict(w, "User already exists!")
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	newUser := repository.User{
		Username: payload.Username,
		Email: payload.Email,
		Password: hashedPassword,
	}

	if err := s.repo.User.Register(ctx, &newUser); err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusCreated, true, nil)
}
