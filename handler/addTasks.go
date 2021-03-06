package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/models"
	"github.com/tejashwikalptaru/tutorial/utilities"
)

func AddTask(writer http.ResponseWriter, request *http.Request) {
	sessionId, err, flag := utilities.MiddlewareAuth(writer, request)
	//log.Printf(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if flag {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var addTask models.AddTask
	err = json.NewDecoder(request.Body).Decode(&addTask)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	taskID, err := helper.CreateTask(sessionId, addTask.Task)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write([]byte(fmt.Sprintf("Task: %s has been created", taskID)))
}
