package errors

import (
	"net/http"

	"github.com/MarcoVitoC/memori/internal/util"
)

func InternalServerError(w http.ResponseWriter, errors any) {
	util.WriteResponse(w, http.StatusInternalServerError, nil, errors)
}

func BadRequest(w http.ResponseWriter, errors any) {
	util.WriteResponse(w, http.StatusBadRequest, nil, errors)
}

func Conflict(w http.ResponseWriter, errors any) {
	util.WriteResponse(w, http.StatusConflict, nil, errors)
}

func NotFound(w http.ResponseWriter, errors any) {
	util.WriteResponse(w, http.StatusNotFound, nil, errors)
}

func Unauthorized(w http.ResponseWriter, errors any) {
	util.WriteResponse(w, http.StatusUnauthorized, nil, errors)
}