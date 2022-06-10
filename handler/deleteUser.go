package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/models"
)

func DeleteRow(writer http.ResponseWriter, request *http.Request) {
	var req models.UpdateUser
	decoder := json.NewDecoder(request.Body)
	addErr := decoder.Decode(&req)
	log.Printf(req.ID)
	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, err := helper.DeleteUser(req.ID)
	log.Printf(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, userErr := helper.GetUser(userID)
	if userErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
}
