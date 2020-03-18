package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"../settings"
)

// var Classrooms = []Classroom

func CreateClassroom(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	var newClassroom Classroom
	err := json.NewDecoder(r.Body).Decode(&newClassroom)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new classroom"})
		showError(r, err)
		return
	}
	settings.DB.Create(&newClassroom)
	fmt.Println(newClassroom)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new classroom"})
	return
}

func GetBuilding(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	path := r.URL.Path
	path = strings.Replace(path, "/api/building/", "", 1)

	if path == "" {
		var buildings []Building
		settings.DB.Find(&buildings)
		json.NewEncoder(w).Encode(buildings)
		return
	}
	var building Building
	settings.DB.Where("id = ?", path).First(&building)

	json.NewEncoder(w).Encode(building)
}

func ClassroomByNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	showAPIRequest(r)

	if r.Method == "GET" {
		var classroom Classroom
		path := r.URL.Path
		path = strings.Replace(path, "/api/classroom/name/", "", 1)

		settings.DB.Where("name = ?", path).First(&classroom)

		json.NewEncoder(w).Encode(classroom)
		return
	}
}

func UpdateClassroom(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	var newClassroom Classroom
	err := json.NewDecoder(r.Body).Decode(&newClassroom)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Classroom"})
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "Successfully updated Classroom"})
	return
}

func ClassroomRoute(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateClassroom(w, r)
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateClassroom(w, r)
		return
	}

	if replacePath(r.URL.Path, "/api/classroom/") == "" {
		GetClassrooms(w, r)
		return
	}

	if r.Method == "GET" {
		if replacePath(r.URL.Path, "/api/classroom/") == "" {
			GetClassrooms(w, r)
			return
		}
		GetClassroom(w, r)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func GetClassroom(w http.ResponseWriter, r *http.Request) {
	path := replacePath(r.URL.Path, "/api/classroom/")
	num, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		showError(r, err)
		return
	}
	var room Classroom
	settings.DB.Where("id = ?", num).First(&room)
	json.NewEncoder(w).Encode(room)
}

func GetClassrooms(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var classrooms []Classroom
		settings.DB.Find(&classrooms)
		json.NewEncoder(w).Encode(classrooms)
		return
	}
}

func GetClassroomByComputer(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		path := r.URL.Path
		path = strings.Replace(path, "/api/classroom/", "", 1)
		if path == "computer" {
			var Classrooms []Classroom
			settings.DB.Where("iscomputer = ?", true).Find(&Classrooms)
			json.NewEncoder(w).Encode(Classrooms)
			return
		}
		if path == "lecture" {
			var Classrooms []Classroom
			settings.DB.Where("iscomputer = ?", false).Find(&Classrooms)
			json.NewEncoder(w).Encode(Classrooms)
			return
		}
	}
}
