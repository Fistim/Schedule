package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	var newSubject Subject
	err := json.NewDecoder(r.Body).Decode(&newSubject)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Subject"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newSubject)
	// if err != nil {
	// 	json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Subject"})
	// 	fmt.Println(err)
	// 	return
	// }
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new Subject"})
	return
}

func ModuleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	showAPIRequest(r)
	if r.Method == "POST" {
		var newModule Professionalmodule
		err := json.NewDecoder(r.Body).Decode(&newModule)
		if err != nil {
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new cycle"})
			fmt.Println(err)
			return
		}
		settings.DB.Create(&newModule)
		json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new cycle"})
		return
	}
	path := replacePath(r.URL.Path, "/api/module/")
	if r.Method == "GET" && path != "" {
		num, err := strconv.Atoi(path)

		if err != nil {
			showError(r, err)
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
			return
		}

		var output Professionalmodule
		settings.DB.First(&output, num)
		if output.ID == 0 {
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error getting teacher"})
			return
		}
		json.NewEncoder(w).Encode(output)
		return
	}

	if r.Method == "GET" && path == "" {
		var output []Professionalmodule
		settings.DB.Find(&output)
		json.NewEncoder(w).Encode(output)
	}

	if r.Method == "PATCH" || r.Method == "UPDATE" {
		var module Professionalmodule
		var moduleToUpdate Professionalmodule
		err := json.NewDecoder(r.Body).Decode(&module)

		if err != nil {
			showError(r, err)
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update module"})
			return
		}
		settings.DB.First(&moduleToUpdate, module.ID)
		settings.DB.Model(&moduleToUpdate).Updates(module)
		json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated module"})
		return
	}
}

func CycleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	showAPIRequest(r)
	if r.Method == "POST" {
		var newCycle Cycle
		err := json.NewDecoder(r.Body).Decode(&newCycle)
		if err != nil {
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new cycle"})
			fmt.Println(err)
			return
		}
		settings.DB.Create(&newCycle)
		json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new cycle"})
		return
	}
	path := replacePath(r.URL.Path, "/api/cycle/")
	if r.Method == "GET" && path != "" {

		num, err := strconv.Atoi(path)

		if err != nil {
			showError(r, err)
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
			return
		}

		var output Cycle
		settings.DB.First(&output, num)
		if output.ID == 0 {
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error getting teacher"})
			return
		}
		json.NewEncoder(w).Encode(output)
		return
	}

	if r.Method == "GET" && path == "" {
		var output []Cycle
		settings.DB.Find(&output)
		json.NewEncoder(w).Encode(output)
	}

	if r.Method == "PATCH" || r.Method == "UPDATE" {
		var cycle Cycle
		var cycleToUpdate Cycle
		err := json.NewDecoder(r.Body).Decode(&cycle)

		if err != nil {
			showError(r, err)
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update cycle"})
			return
		}
		settings.DB.First(&cycleToUpdate, cycle.ID)
		settings.DB.Model(&cycleToUpdate).Updates(cycle)
		json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated cycle"})
		return
	}
}

func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	var subject Subject
	var subjectToUpdate Subject
	err := json.NewDecoder(r.Body).Decode(&subject)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update subject"})
		return
	}
	settings.DB.First(&subjectToUpdate, subject.ID)
	settings.DB.Model(&subjectToUpdate).Updates(subject)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated subject"})
	return
}

func SubjectRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	showAPIRequest(r)

	if r.Method == "POST" {
		CreateSubject(w, r)
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateSubject(w, r)
		return
	}

	path := r.URL.Path

	path = strings.Replace(path, "/api/subject/", "", 1)

	if path == "" {
		if r.Method == "GET" {
			var output []Subject
			settings.DB.Find(&output)
			json.NewEncoder(w).Encode(output)
		}
		return
	}

	if path == "" && r.Method == "GET" {
		GetSubjects(w, r)
	}

	if r.Method == "GET" {
		GetSubject(w, r)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "It is not GET method"})
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	path := replacePath(r.URL.Path, "/api/subject/")
	num, err := strconv.Atoi(path)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		return
	}

	var output Subject
	settings.DB.First(&output, num)
	if output.ID == 0 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error getting subject"})
		return
	}
	json.NewEncoder(w).Encode(output)
	return
}

func GetSubjects(w http.ResponseWriter, r *http.Request) {
	var output []Subject
	settings.DB.Find(&output)
	json.NewEncoder(w).Encode(output)
}



func GetSubjectByModule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	showAPIRequest(r)
	if r.Method == "GET" {
		path := r.URL.Path
		path = strings.Replace(path, "/api/subject/module/", "", 1)
		num, err := strconv.Atoi(path)

		if err != nil {
			showError(r, err)
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
			return
		}
		var output []Subject
		settings.DB.Where("id_professionalmodule = ?", num).Find(&output)
		json.NewEncoder(w).Encode(output)
	}
}


