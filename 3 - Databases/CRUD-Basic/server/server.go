package server

import (
	"encoding/json"
	"fmt"
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

	db, erro := db.connect()
	if erro != nil {
		w.Write([]byte("DB connection error"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into users (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Statement creator error"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Nome, user.Email)
	if erro != nil {
		w.Write([]byte("Statement execution error"))
		return
	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("ID insert error"))
		return
	}

	w.Write([]byte(fmt.Sprintf("insert user ok! ID: %d", idInsert)))
}
