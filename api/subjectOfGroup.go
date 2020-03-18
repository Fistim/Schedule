package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//"strings"
)

var SubjectOfGroups []Subjectofgroup

func SubjectOfGroupRoute(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateSubjectOfGroup(w, r)
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateSubjectOfGroup(w, r)
		return
	}

	path := replacePath(r.URL.Path, "/api/groupsubject/")

	if path == "" {
		GetSubjectOfGroups(w, r)
		return
	}

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&SubjectOfGroups)
	if num < 0 || num > len(SubjectOfGroups)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetSubjectOfGroupById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateSubjectOfGroup(w http.ResponseWriter, r *http.Request) {
	var SubjectOfGroup Subjectofgroup
	var SubjectOfGroupToUpdate Subjectofgroup
	err := json.NewDecoder(r.Body).Decode(&SubjectOfGroup)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update subjectOfGroup"})
		return
	}
	settings.DB.First(&SubjectOfGroupToUpdate, SubjectOfGroup.ID)
	settings.DB.Model(&SubjectOfGroupToUpdate).Updates(SubjectOfGroup)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated subjectOfGroup"})
	return
}

func CreateSubjectOfGroup(w http.ResponseWriter, r *http.Request) {
	var newSubjectOfGroup Subjectofgroup
	err := json.NewDecoder(r.Body).Decode(&newSubjectOfGroup)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new subjectOfGroup"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newSubjectOfGroup)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new s ubjectOfGroup"})
	return
}

func GetSubjectOfGroupById(w http.ResponseWriter, r *http.Request, num int) {
	var SubjectOfGroup Subjectofgroup
	settings.DB.Where("id = ?", num).First(&SubjectOfGroup)
	json.NewEncoder(w).Encode(SubjectOfGroup)
	return
}

func GetSubjectOfGroups(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var SubjectOfGroups []Subjectofgroup
		settings.DB.Find(&SubjectOfGroups)
		json.NewEncoder(w).Encode(SubjectOfGroups)
	}
	return
}
