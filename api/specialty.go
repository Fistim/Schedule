package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CreateSpecialty(w http.ResponseWriter, r *http.Request) {
	var newSpecialty Specialty
	err := json.NewDecoder(r.Body).Decode(&newSpecialty)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Specialty"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newSpecialty)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new Specialty"})
	return
}

func GetSpecialty(w http.ResponseWriter, r *http.Request) {
	var Specialty Specialty
	path := replacePath(r.URL.Path, "/api/specialty/")
	num, err := strconv.Atoi(path)
	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		return
	}
	settings.DB.Where("id = ?", num).First(&Specialty)
	json.NewEncoder(w).Encode(Specialty)
	return
}

func GetSpecialties(w http.ResponseWriter, r *http.Request) {
	var Specialties []Specialty
	settings.DB.Find(&Specialties)
	json.NewEncoder(w).Encode(Specialties)
}

func UpdateSpecialty(w http.ResponseWriter, r *http.Request) {
	var specialty Specialty
	var specialtyToUpdate Specialty
	err := json.NewDecoder(r.Body).Decode(&specialty)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update specialty"})
		return
	}
	settings.DB.First(&specialtyToUpdate, specialty.ID)
	settings.DB.Model(&specialtyToUpdate).Updates(specialty)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated specialty"})
	return
}

func SpecialtyRoute(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateSpecialty(w, r)
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateSpecialty(w, r)
		return
	}

	path := replacePath(r.URL.Path, "/api/specialty/")

	if path == "" && r.Method == "GET" {
		GetSpecialties(w, r)
		return
	}

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if num < 0 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetSpecialty(w, r)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "It is not GET method"})
}
