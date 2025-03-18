package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/MarcoVitoC/memori/internal/service"
)

type Server struct {
	Addr string
}

func (s *Server) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	svc := service.NewService()
	
	r.Route("/diaries", func(r chi.Router) {
		r.Get("/", svc.Diary.GetAll)
		r.Post("/", svc.Diary.Create)
	})

	return r
}

func (s *Server) Run(mux http.Handler) error {
	server := http.Server{
		Addr: s.Addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("INFO: server has started at %s", s.Addr)
	return server.ListenAndServe()
}