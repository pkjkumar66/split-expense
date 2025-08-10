package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, Response{Error: message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(Response{Data: payload})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
