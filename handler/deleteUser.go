package handler

import (
	"encoding/json"
	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/models"
	"log"
	"net/http"
)

func DeleteRow(writer http.ResponseWriter, request *http.Request) {
	var req models.DeleteUser
	decoder := json.NewDecoder(request.Body)
	addErr := decoder.Decode(&req)

	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := helper.DeleteUser(req.ID)
	if err == nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf(req.ID)
}
