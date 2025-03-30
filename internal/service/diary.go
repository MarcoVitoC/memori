package service

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/util"
	"github.com/MarcoVitoC/memori/internal/repository"
)

type CreateDiaryPayload struct {
	Content string `json:"content"`
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

func (s *DiaryService) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload CreateDiaryPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		util.WriteError(w, http.StatusBadRequest, err.Error())
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