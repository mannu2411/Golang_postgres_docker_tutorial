package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/models"
)

func SignInUser(writer http.ResponseWriter, request *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		log.Printf("JSON ERROR")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	expectPass, err := helper.GetPass(creds.Email)

	if err != nil || expectPass != creds.Pass {
		writer.WriteHeader(http.StatusUnauthorized)
	}

	expireAt := time.Now().Add(120 * time.Second)
	sessionId, err := helper.GetSession(creds.Email)

	if sessionId != "" {
		err = helper.UpdateSession(expireAt, sessionId)

	} else {
		sessionId, err = helper.CreateSession(creds.Email, expireAt)
	}
	if err != nil {
		log.Printf("SESSION ERROR")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionId,
		Expires: expireAt,
	})

}
