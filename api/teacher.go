package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"../settings"
)

var Teachers []Teacher

func TeacherRoute(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateTeacher(w, r)
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateTeacher(w, r)
		return
	}

	path := r.URL.Path

	path = strings.Replace(path, "/api/teacher/", "", 1)

	if path==""{
		GetTeachers(w, r)
		return
	}

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&Teachers)
	if num < 0 || num > len(Teachers)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		//db.Where("name = ?", "jinzhu").First(&user)
		GetTeacherById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request){
		var teacher Teacher
		var teacherToUpdate Teacher
		err:= json.NewDecoder(r.Body).Decode(&teacher)

		if err!=nil{
			showError(r, err)
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update teacher"})
			return
		}
		settings.DB.First(&teacherToUpdate, teacher.ID)
		settings.DB.Model(&teacherToUpdate).Updates(teacher)
		json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated teacher"})
		return
}

func CreateTeacher(w http.ResponseWriter, r *http.Request){
	var newTeacher Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeacher)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new teacher"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newTeacher)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new teacher"})
	return
}

func GetTeacherById(w http.ResponseWriter, r *http.Request, num int){
	var teacher Teacher
	settings.DB.Where("id = ?", num).First(&teacher)
	json.NewEncoder(w).Encode(teacher)
	return
}

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var Teachers []Teacher
		settings.DB.Find(&Teachers)
		json.NewEncoder(w).Encode(Teachers)
	}
	return
}
