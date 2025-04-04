package service

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/repository"
	"github.com/MarcoVitoC/memori/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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
		util.WriteError(w, http.StatusInternalServerError, err.Error())
	}

	util.WriteResponse(w, http.StatusOK, diaries)
}

func (s *DiaryService) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	diary, err := s.repo.Diary.GetById(ctx, id)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, diary)
}

func (s *DiaryService) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload CreateOrUpdateDiaryPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		util.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		util.WriteError(w, http.StatusBadRequest, "Content is required!")
		return
	}
	
	newDiary := &repository.Diary{
		Content: payload.Content,
	}

	if err := s.repo.Diary.Create(ctx, newDiary); err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusCreated, nil)
}

func (s *DiaryService) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload CreateOrUpdateDiaryPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id := chi.URLParam(r, "id")
	updatedDiary := repository.Diary{
		Content: payload.Content,
	}

	if err := s.repo.Diary.Update(ctx, id, &updatedDiary); err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, nil)
}

func (s *DiaryService) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if err := s.repo.Diary.Delete(ctx, id); err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, nil)
}
