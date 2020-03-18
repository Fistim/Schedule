package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var subjecttype []Subjecttype

func SubjecttypeRoute(w http.ResponseWriter, r *http.Request) {
	
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	
	if r.Method == "POST" {
		CreateSubjecttype(w, r)
		return
	}

	
	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateSubjecttype(w, r)
		return
	}

	
	path := replacePath(r.URL.Path, "/api/subjecttype/")

	if path == "" {
		GetSubjecttype(w, r)
		return
	}
	

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&ScheduleOfGroups)
	if num < 0 || num > len(ScheduleOfGroups)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetSubjecttypeById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateSubjecttype(w http.ResponseWriter, r *http.Request) {
	var subjecttype Subjecttype
	var SubjecttypeToUpdate Subjecttype
	err := json.NewDecoder(r.Body).Decode(&subjecttype)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update subjecttype"})
		return
	}
	settings.DB.First(&SubjecttypeToUpdate, subjecttype.ID)
	settings.DB.Model(&SubjecttypeToUpdate).Updates(subjecttype)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated subjecttype"})
	return
}

func CreateSubjecttype(w http.ResponseWriter, r *http.Request) {
	var newSubjecttype Subjecttype
	err := json.NewDecoder(r.Body).Decode(&newSubjecttype)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Subjecttype"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newSubjecttype)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new Subjecttype"})
	return
}

func GetSubjecttypeById(w http.ResponseWriter, r *http.Request, num int) {
	var subjecttype Subjecttype
	settings.DB.Where("id = ?", num).First(&subjecttype)
	json.NewEncoder(w).Encode(subjecttype)
	return
}

func GetSubjecttype(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var subjecttype []Subjecttype
		settings.DB.Find(&subjecttype)
		json.NewEncoder(w).Encode(subjecttype)
	}
	return
}
