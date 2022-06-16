package utilities

import (
	"encoding/json"
	"net/http"

	"github.com/tejashwikalptaru/tutorial/models"
)

func JsonFetch(user *models.User, writer http.ResponseWriter) http.ResponseWriter {
	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return writer
	}
	writer.Write(jsonData)
	return writer
}
