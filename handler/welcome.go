package handler

import (
	"encoding/json"
	"github.com/tejashwikalptaru/tutorial/database/helper"
	"net/http"
)

func isErr(err error, typeErr string) bool {
	return err.Error() == "pq: duplicate key value violates unique constraint "+typeErr
}

func Greet(writer http.ResponseWriter, request *http.Request) {
	userID, err := helper.CreateUser("test", "test@test.com")
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
