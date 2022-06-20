package utilities

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tejashwikalptaru/tutorial/database/helper"
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

func MiddlewareAuth(writer http.ResponseWriter, request *http.Request) (string, error, bool) {
	session := request.Header.Values("session_token")
	sessionId := session[0]

	exist, err := helper.SessionExist(sessionId)
	if err != nil {
		return "", err, true
	}
	if !exist {
		return "", err, true
	}
	isExpired, err := helper.IsExpired(sessionId)

	if isExpired {
		return "", err, true
	}
	log.Printf(sessionId)
	return sessionId, nil, false
}
