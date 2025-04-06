package service

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/errors"
	"github.com/MarcoVitoC/memori/internal/repository"
	"github.com/MarcoVitoC/memori/internal/util"
	"github.com/go-chi/chi/v5"
)

type CreateOrUpdateDiaryPayload struct {
	Content string `json:"content" validate:"required"`
}

type DiaryService struct {
	repo repository.Repository
}

func (s *DiaryService) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	diaries, err := s.repo.Diary.GetAll(ctx)
	if err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, diaries, nil)
}

func (s *DiaryService) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	diary, err := s.repo.Diary.GetById(ctx, id)
	if err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, diary, nil)
}

func (s *DiaryService) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload CreateOrUpdateDiaryPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	if errs := util.Validate(payload); len(errs) > 0 {
		errors.BadRequest(w, errs)
		return
	}
	
	newDiary := &repository.Diary{
		Content: payload.Content,
	}

	if err := s.repo.Diary.Create(ctx, newDiary); err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusCreated, true, nil)
}

func (s *DiaryService) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload CreateOrUpdateDiaryPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	if errs := util.Validate(payload); len(errs) > 0 {
		errors.BadRequest(w, errs)
		return
	}

	id := chi.URLParam(r, "id")
	updatedDiary := repository.Diary{
		Content: payload.Content,
	}

	if err := s.repo.Diary.Update(ctx, id, &updatedDiary); err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, true, nil)
}

func (s *DiaryService) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if err := s.repo.Diary.Delete(ctx, id); err != nil {
		errors.InternalServerError(w, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, true, nil)
}
