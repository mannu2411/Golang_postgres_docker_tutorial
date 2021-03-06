package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tejashwikalptaru/tutorial/database/helper"
)

func AllUsers(writer http.ResponseWriter, request *http.Request) {
	user, userErr := helper.GetAllUser()
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
