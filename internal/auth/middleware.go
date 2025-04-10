package auth

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/errors"
)

func AuthMiddleware(authenticator *JWTAuthenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwt, err := r.Cookie("JWT_TOKEN")
			if err != nil || jwt.Value == "" {
				errors.Unauthorized(w, "Please login!")
				return
			}
	
			xsrf, err := r.Cookie("XSRF_TOKEN")
			if err != nil || xsrf.Value == "" {
				errors.Unauthorized(w, "Please login!")
				return
			}
	
			if r.Header.Get("X-CSRF-TOKEN") != xsrf.Value {
				errors.Unauthorized(w, "Invalid credentials!")
				return
			}
			
			if _, err := authenticator.VerifyJWT(jwt.Value); err != nil {
				errors.Unauthorized(w, "Invalid credentials!")
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}