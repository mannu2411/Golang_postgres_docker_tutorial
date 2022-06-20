package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/utilities"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	sessionId, err, flag := utilities.MiddlewareAuth(writer, request)
	log.Printf(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if flag {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	task, err := helper.GetAllTasks(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(task)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
}
