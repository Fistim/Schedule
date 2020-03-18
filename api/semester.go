package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var Semesters []Semester

func SemesterRoute(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	path := replacePath(r.URL.Path, "/api/semester/")
	num, err := strconv.Atoi(path)

	if path == "" {
		GetSemesters(w, r)
		return
	}

	if strings.Contains(path, "weeks/") {
		GetSemestersByWeeks(w, r, num)
		return
	}

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&Semesters)
	if num < 0 || num > len(Semesters)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetSemesterById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func GetSemestersByWeeks(w http.ResponseWriter, r *http.Request, num int) {
	path := replacePath(r.URL.Path, "/api/semester/weeks/")

	num, err := strconv.Atoi(path)
	if err != nil || num < 0 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		return
	}

	var semesters []Semester
	settings.DB.Where("weeksquantity = ?", num).Find(&semesters)
	json.NewEncoder(w).Encode(semesters)
	return
}

func GetSemesterById(w http.ResponseWriter, r *http.Request, num int) {
	var Semester Semester
	settings.DB.Where("id = ?", num).First(&Semester)
	json.NewEncoder(w).Encode(Semester)
	return
}

func GetSemesters(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var Semesters []Semester
		settings.DB.Find(&Semesters)
		json.NewEncoder(w).Encode(Semesters)
	}
	return
}
