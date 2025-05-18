package service

import (
	"net/http"
	"time"

	"github.com/MarcoVitoC/memori/internal/auth"
	"github.com/MarcoVitoC/memori/internal/errors"
	"github.com/MarcoVitoC/memori/internal/repository"
	"github.com/MarcoVitoC/memori/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repo repository.Repository
	authenticator *auth.JWTAuthenticator
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (s *AuthService) Register(w http.ResponseWriter, r *http.Request) {
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

	user, err := s.repo.User.GetByEmail(ctx, payload.Email)
	if err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	if user != nil {
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

func (s *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload LoginUserPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	if errs := util.Validate(payload); len(errs) > 0 {
		errors.BadRequest(w, errs)
		return
	}

	user, err := s.repo.User.GetByEmail(ctx, payload.Email)
	if err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	if user == nil || !auth.MatchPassword(user.Password, payload.Password) {
		errors.BadRequest(w, "Invalid email or password!")
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"iss": "memori",
		"aud": "memori",
	}

	jwt, jwtErr := s.authenticator.GenerateJWT(claims)
	if jwtErr != nil {
		errors.InternalServerError(w, jwtErr.Error())
		return
	}

	xsrf, xsrfErr := auth.GenerateToken(32)
	if xsrfErr != nil {
		errors.InternalServerError(w, xsrfErr.Error())
		return
	}

	setCookies(w, jwt, xsrf)
	util.WriteResponse(w, http.StatusOK, true, nil)
}

func (s *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	setCookies(w, "", "")
	util.WriteResponse(w, http.StatusOK, true, nil)
}

func setCookies(w http.ResponseWriter, jwt string, xsrf string) {
	expireTime := time.Unix(0, 0)
	
	if jwt != "" && xsrf != "" {
		expireTime = time.Now().Add(time.Hour * 24 * 3)
	}

	http.SetCookie(w, &http.Cookie{
		Name: "JWT-Token",
		Value: jwt,
		Path: "/",
		Expires: expireTime,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name: "X-CSRF-Token",
		Value: xsrf,
		Path: "/",
		Expires: expireTime,
		HttpOnly: false,
	})
}
