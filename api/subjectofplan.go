package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var subjectofplan []Subjectofplan

func SubjectofplanRoute(w http.ResponseWriter, r *http.Request) {
	
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	
	if r.Method == "POST" {
		CreateSubjectofplan(w, r)
		return
	}

	
	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateSubjectofplan(w, r)
		return
	}

	
	path := replacePath(r.URL.Path, "/api/subjectofplan/")

	if path == "" {
		GetSubjectofplan(w, r)
		return
	}
	

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&subjectofplan)
	if num < 0 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetSubjectofplanById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateSubjectofplan(w http.ResponseWriter, r *http.Request) {
	var subjectofplan Subjectofplan
	var SubjectofplanToUpdate Subjectofplan
	err := json.NewDecoder(r.Body).Decode(&subjectofplan)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Subjectofplan"})
		return
	}
	settings.DB.First(&SubjectofplanToUpdate, subjectofplan.ID)
	settings.DB.Model(&SubjectofplanToUpdate).Updates(subjectofplan)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated Subjectofplan"})
	return
}

func CreateSubjectofplan(w http.ResponseWriter, r *http.Request) {
	var newSubjectofplan Subjectofplan
	err := json.NewDecoder(r.Body).Decode(&newSubjectofplan)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Subjectofplan"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newSubjectofplan)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new Subjectofplan"})
	return
}

func GetSubjectofplanById(w http.ResponseWriter, r *http.Request, num int) {
	var subjectofplan Subjectofplan
	settings.DB.Where("id = ?", num).First(&subjectofplan)
	json.NewEncoder(w).Encode(subjectofplan)
	return
}

func GetSubjectofplan(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var subjectofplan []Subjectofplan
		settings.DB.Find(&subjectofplan)
		json.NewEncoder(w).Encode(subjectofplan)
	}
	return
}
