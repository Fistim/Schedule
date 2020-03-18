package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var group []Group

func GroupRoute(w http.ResponseWriter, r *http.Request) {
	
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateGroup(w, r)
		return
	}

	
	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateGroup(w, r)
		return
	}

	
	path := replacePath(r.URL.Path, "/api/group/")

	if path == ""  && r.Method=="GET"{
		GetGroup(w, r)
		return
	}

	if strings.Contains(path, "specialty") && r.Method=="GET"{
		GetGroupBySpecialty(w, r)
		return
	}
	

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&group)
	if num < 0 || num > len(group)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}



	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var group Group
	var groupToUpdate Group
	err := json.NewDecoder(r.Body).Decode(&group)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Group"})
		return
	}
	settings.DB.First(&groupToUpdate, group.ID)
	settings.DB.Model(&groupToUpdate).Updates(group)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated Group"})
	return
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var newGrop Group
	err := json.NewDecoder(r.Body).Decode(&newGrop)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Group"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newGrop)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new newGrop"})
	return
}

func GetGroupBySpecialty(w http.ResponseWriter, r *http.Request,) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	path := replacePath(r.URL.Path, "/api/group/specialty/")
	var output []Group
	num, err := strconv.Atoi(path)
	if err != nil || num < 0 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		showError(r, err)
		return
	}
	settings.DB.Where("id_specialty = ?", num).Find(&output)
	json.NewEncoder(w).Encode(output)
	return
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var group []Group
		settings.DB.Find(&group)
		json.NewEncoder(w).Encode(group)
	}
	return
}

// func InArchive(w http.ResponseWriter, r *http.Request, num int){
// 	var group Group
// 	var specialty Specialty
// 	var durationofsudy Durationofsudy
// 	var studyplan Studypaln

// 	settings.DB.Where("groupnumber = ?", num).First(&group)
// 	settings.DB.Where("id = ?", group.IDSpecialty).First(&specialty)
// 	settings.DB.Where("id_duration = ?", specialty.IDDuration).First(&durationofsudy)
// 	settings.DB.Where("id_group = ?", group.ID).First(&studyplan)
	
	
// 	if group.Year > durationofsudy.Yearsquantity {
// 		studyplan.IsArchive = true
// 	}

// 	return

// }


