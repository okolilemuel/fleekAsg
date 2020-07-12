package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func respondWithJSON(data interface{}, statusCode int, success bool, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	httpResponse := response{
		Success: success,
		Data:    data,
		Message: message,
	}

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(httpResponse)

	if err != nil {
		log.Println("[ERROR] respondWithJSON() : encountered and error while converting data to JSON")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}

}

func respondWithFile(w http.ResponseWriter, file []byte, filename string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	w.Write(file)
}
