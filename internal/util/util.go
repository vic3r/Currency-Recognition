package util

import (
	"encoding/json"
	"net/http"

	e "github.com/vic3r/Currency-Recognition/errors"
)

// ResponseOk adjust the response in the  header and returns a new object
func ResponseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

// ResponseError returns an error Reponse
func ResponseError(w http.ResponseWriter, err *e.ErrResponse) {
	if err == nil {
		return
	}
	w.WriteHeader(err.Code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": err.Name,
	}
	json.NewEncoder(w).Encode(body)
}
