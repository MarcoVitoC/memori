package handler

import (
	"net/http"
	"time"

	"github.com/MarcoVitoC/memori/internal/auth"
	"github.com/MarcoVitoC/memori/internal/repository"
	"github.com/MarcoVitoC/memori/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Server struct {
	Addr string
	DB *pgxpool.Pool
	Authenticator *auth.JWTAuthenticator
}

func (s *Server) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	repo := repository.NewRepository(s.DB)
	svc := service.NewService(repo, s.Authenticator)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", svc.Auth.Register)
		r.Post("/login", svc.Auth.Login)
		r.Post("/logout", svc.Auth.Logout)
	})
	
	r.Route("/diaries", func(r chi.Router) {
		r.Use(auth.AuthMiddleware(s.Authenticator))
		
		r.Get("/", svc.Diary.GetAll)
		r.Get("/{id}", svc.Diary.GetById)
		r.Post("/", svc.Diary.Create)
		r.Put("/{id}", svc.Diary.Update)
		r.Delete("/{id}", svc.Diary.Delete)
	})

	return r
}

func (s *Server) Run(logger *zap.SugaredLogger, mux http.Handler) error {
	server := http.Server{
		Addr: s.Addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	logger.Infof("Server has started at %s", s.Addr)
	return server.ListenAndServe()
}