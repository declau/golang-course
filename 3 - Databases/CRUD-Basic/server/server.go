package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Create a user in DB
func CreatUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Failed to read request body"))
		return
	}

	var user user

	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		w.Write([]byte("Erro t convert user to struct"))
		return
	}
}
