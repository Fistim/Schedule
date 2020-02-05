package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Subjects = make(map[uint8]Subject)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Subject Subject
	err := json.NewDecoder(r.Body).Decode(&Subject)
	if err != nil {
		//TODO
		return
	}

	Subjects[Subject.ID] = Subject
	json.NewEncoder(w).Encode(Subject)
}

func SubjectById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var newSubject Subject
		err := json.NewDecoder(r.Body).Decode(&newSubject)
		if err != nil {
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Subject"})
			return
		}
		Subjects[newSubject.ID] = newSubject
		json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new Subject"})
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		var newSubject Subject
		err := json.NewDecoder(r.Body).Decode(&newSubject)
		if err != nil {
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Subject"})
			return
		}
		if _, ok := Subjects[newSubject.ID]; !ok {
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Subject"})
			return
		}
		Subjects[newSubject.ID] = newSubject
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "Successfully updated Subject"})
		return
	}

	path := r.URL.Path

	path = strings.Replace(path, "/api/Subject/", "", 1)

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
	}

	if _, ok := Subjects[uint8(num)]; !ok {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		return
	}

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(Subjects[uint8(num)])
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "It is not GET method"})
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(Subjects)
	}
}

// func GetSubjectByComputer(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	if r.Method == "GET" {
// 		path := r.URL.Path
// 		path = strings.Replace(path, "/api/Subject/", "", 1)
// 		var output []Subject
// 		for _, room := range Subjects {
// 			if room.IsComputer && path == "computer" {
// 				output = append(output, room)
// 			}
// 			if !room.IsComputer && path == "lecture" {
// 				output = append(output, room)
// 			}
// 		}
// 		json.NewEncoder(w).Encode(output)
// 	}
// }
