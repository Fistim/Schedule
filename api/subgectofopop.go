package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var subjectofexample []Subjectofexample

func SubjectofexampleRoute(w http.ResponseWriter, r *http.Request) {
	
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateSubjectofexample(w, r)
		return
	}

	
	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateSubjectofexample(w, r)
		return
	}

	
	path := replacePath(r.URL.Path, "/api/subjectofexample/")

	if path == "" {
		GetSubjectofexample(w, r)
		return
	}
	

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&studyplan)
	if num < 0 || num > len(studyplan)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetSubjectofexampleById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateSubjectofexample(w http.ResponseWriter, r *http.Request) {
	var subjectofexample Subjectofexample
	var subjectofexampleToUpdate Subjectofexample
	err := json.NewDecoder(r.Body).Decode(&subjectofexample)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update subjectofexample"})
		return
	}
	settings.DB.First(&subjectofexampleToUpdate, subjectofexample.ID)
	settings.DB.Model(&subjectofexampleToUpdate).Updates(subjectofexample)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated subjectofexample"})
	return
}

func CreateSubjectofexample(w http.ResponseWriter, r *http.Request) {
	var newsubjectofexample Subjectofexample
	err := json.NewDecoder(r.Body).Decode(&newsubjectofexample)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new subjectofexample"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newsubjectofexample)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new subjectofexample"})
	return
}

func GetSubjectofexampleById(w http.ResponseWriter, r *http.Request, num int) {
	var subjectofexample Subjectofexample
	settings.DB.Where("id = ?", num).First(&subjectofexample)
	json.NewEncoder(w).Encode(subjectofexample)
	return
}

func GetSubjectofexample(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var subjectofexample []Subjectofexample
		settings.DB.Find(&subjectofexample)
		json.NewEncoder(w).Encode(subjectofexample)
	}
	return
}
