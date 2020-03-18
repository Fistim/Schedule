package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	// "strings"
)

var Lessons []Lesson

func BellScheduleRoute(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateLesson(w, r)
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateLesson(w, r)
		return
	}

	path := replacePath(r.URL.Path, "/api/bellschedule/")

	if path == "" {
		GetLessons(w, r)
		return
	}

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&Lessons)
	if num < 0 || num > len(Lessons)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		//db.Where("name = ?", "jinzhu").First(&user)
		GetLessonById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateLesson(w http.ResponseWriter, r *http.Request) {
	var lesson Lesson
	var LessonToUpdate Lesson
	err := json.NewDecoder(r.Body).Decode(&lesson)

	

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Lesson"})
		return
	}
	settings.DB.First(&LessonToUpdate, lesson.ID)
	settings.DB.Model(&LessonToUpdate).Updates(lesson)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated Lesson"})
	return
}

func CreateLesson(w http.ResponseWriter, r *http.Request) {
	var newLesson Lesson
	err := json.NewDecoder(r.Body).Decode(&newLesson)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Lesson"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newLesson)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new Lesson"})
	return
}

func GetLessonById(w http.ResponseWriter, r *http.Request, num int) {
	var Lesson Lesson
	settings.DB.Where("id = ?", num).First(&Lesson)
	json.NewEncoder(w).Encode(Lesson)
	return
}

func GetLessons(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var Lessons1 []Lesson
		settings.DB.Find(&Lessons1)
		json.NewEncoder(w).Encode(Lessons1)
	}
	return
}
