package http

import (
	"encoding/json"
	"net/http"

	"github.com/PhilLar/ToISD/models"
)

const (
	headerContentType   = "Content-Type"
	mimeApplicationJSON = "application/json"
)

type UserStore interface {
	CreateUser(user models.User) (err error)
	AuthUser(creds models.Creds) (tokenString string, err error)
}

type Env struct {
	Store UserStore
}

func respondJSON(w http.ResponseWriter, code int, data interface{}) {

	if data == nil {
		w.WriteHeader(code)
		return
	}

	w.Header().Set(headerContentType, mimeApplicationJSON)
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		respondError(w, err)
		return
	}

	w.Header().Set(headerContentType, mimeApplicationJSON)
}

func respondError(w http.ResponseWriter, err error) {
	respondJSON(w, http.StatusInternalServerError, err)
}

func respondOK(w http.ResponseWriter, data interface{}) {
	respondJSON(w, http.StatusOK, data)
}

func processRequest(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}

	// metas, err := srv.validate.Struct(data)
	// if err != nil {
	// 	return err
	// }

	// if metas != nil {
	// 	return model.ValidationFailureError().WithMetas(metas...)
	// }

	return nil
}
