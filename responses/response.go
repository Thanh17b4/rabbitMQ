package responses

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, message string) error {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		return err
	}
	return nil
}

func Success(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error().Err(err).Interface("data", data).Msg("Could not get response successfully ")
		return
	}
}
