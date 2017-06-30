package main

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	CreatedOn	time.Time	`json:"createdOn"`
}

//store notes collection
var noteStore = make(map[string]Note)

//generate key for the collection
var id int = 0

//POST

func PostNoteHandler(w http.ResponseWriter, r *http.Request)
{
	var note Note
	//Decode the incoming JSON
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request)
{
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {

}
