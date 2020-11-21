package json_utils

import (
	"Honest-Game-Reviews/src/utils/errors"
	"encoding/json"
	"net/http"
)


func JsonResponse(w http.ResponseWriter, statusCode int, body interface{})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func JsonErrorResponse(w http.ResponseWriter, err *errors.RestErrors) {
	JsonResponse(w, err.Status, err)
}

func ClientErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
