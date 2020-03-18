package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var ScheduleOfGroups []Scheduleofgroup

func ScheduleOfGroupRoute(w http.ResponseWriter, r *http.Request) {

	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateScheduleOfGroup(w, r)
		return
	}

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateScheduleOfGroup(w, r)
		return
	}

	path := replacePath(r.URL.Path, "/api/groupschedule/")

	if path == "" {
		GetScheduleOfGroups(w, r)
		return
	}

	if strings.Contains(path, "day/") {

		day := strings.Replace(path, "day/", "", 1)
		switch day {
		case "monday", "tuesday", "wednesday", "thursday", "friday", "saturday":
		default:
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "undefined day error"})
			return
		}

		GetScheduleByDay(w, r, day)
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
		GetScheduleOfGroupById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func GetScheduleByDay(w http.ResponseWriter, r *http.Request, day string) {
	var ScheduleOfday []Schedule
	settings.DB.Where("weekday = ?", day).Find(&ScheduleOfday)
	json.NewEncoder(w).Encode(ScheduleOfday)
	return
}

func UpdateScheduleOfGroup(w http.ResponseWriter, r *http.Request) {
	var ScheduleOfGroup Scheduleofgroup
	var ScheduleOfGroupToUpdate Scheduleofgroup
	err := json.NewDecoder(r.Body).Decode(&ScheduleOfGroup)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update ScheduleOfGroup"})
		return
	}
	settings.DB.First(&ScheduleOfGroupToUpdate, ScheduleOfGroup.ID)
	settings.DB.Model(&ScheduleOfGroupToUpdate).Updates(ScheduleOfGroup)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated ScheduleOfGroup"})
	return
}

func CreateScheduleOfGroup(w http.ResponseWriter, r *http.Request) {
	var newScheduleOfGroup Scheduleofgroup
	err := json.NewDecoder(r.Body).Decode(&newScheduleOfGroup)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new ScheduleOfGroup"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newScheduleOfGroup)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new ScheduleOfGroup"})
	return
}

func GetScheduleOfGroupById(w http.ResponseWriter, r *http.Request, num int) {
	var ScheduleOfGroup Scheduleofgroup
	settings.DB.Where("id = ?", num).First(&ScheduleOfGroup)
	json.NewEncoder(w).Encode(ScheduleOfGroup)
	return
}

func GetScheduleOfGroups(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var ScheduleOfGroups []Scheduleofgroup
		settings.DB.Find(&ScheduleOfGroups)
		json.NewEncoder(w).Encode(ScheduleOfGroups)
	}
	return
}
