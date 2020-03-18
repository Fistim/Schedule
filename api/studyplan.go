package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var studyplan []Studyplan

func StudyplanRoute(w http.ResponseWriter, r *http.Request) {
	
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateStudyPlan(w, r)
		return
	}

	
	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateStudyPlan(w, r)
		return
	}

	
	path := replacePath(r.URL.Path, "/api/studyplan/")

	if path == "" {
		GetStudyPlan(w, r)
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
		GetStudyPlanById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateStudyPlan(w http.ResponseWriter, r *http.Request) {
	var StudyPlan Studyplan
	var StudyplanToUpdate Studyplan
	err := json.NewDecoder(r.Body).Decode(&StudyPlan)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update StudyPlan"})
		return
	}
	settings.DB.First(&StudyplanToUpdate, StudyPlan.ID)
	settings.DB.Model(&StudyplanToUpdate).Updates(StudyPlan)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated StudyPlan"})
	return
}

func CreateStudyPlan(w http.ResponseWriter, r *http.Request) {
	var newStudyPlan Studyplan
	err := json.NewDecoder(r.Body).Decode(&newStudyPlan)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new StudyPlan"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newStudyPlan)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new StudyPlan"})
	return
}

func GetStudyPlanById(w http.ResponseWriter, r *http.Request, num int) {
	var StudyPlan Studyplan
	settings.DB.Where("id = ?", num).First(&StudyPlan)
	json.NewEncoder(w).Encode(StudyPlan)
	return
}

func GetStudyPlan(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var StudyPlan []Studyplan
		settings.DB.Find(&StudyPlan)
		json.NewEncoder(w).Encode(StudyPlan)
	}
	return
}

type Hours struct{
	IDSubject uint
	Variative uint
}

// func HourseComparison(w http.ResponseWriter, r *http.Request, num int){
// 	var exampleprogram Exampleprogram
// 	var subplan []Subjectofplan
// 	var group Group
// 	var opop []Subjectofexample
// 	var hoursdifferense HourseComparison
// 	var hoursDifference Hours
	
// 	settings.DB.Where("groupnumber = ?", num).First(&group)
// 	groupid := group.ID
// 	var studyplan Studyplan
// 	settings.DB.Where("id_group = ?", groupid).First(&studyplan)
// 	studyplanid := studyplan.ID
// 	// Находит нужный subjectplan по номеру группы
// 	settings.DB.Where("id_plan = ?", studyplanid).Find(&subplan)
// 	//Находит нужный subjectofexample по номеру группы
// 	exampleprogramid := settings.DB.Where("id_specialty = ?", studyplan.IDSpecialty).First(&exampleprogram)
// 	settings.DB.Where("id_exampleprogram = ?", exampleprogramid).Find(&opop)

// 	for _, subjectPlan := range subplan{
// 		for _, subjectOpop := range opop{
// 			var hd Hours
// 				if subjectPlan.ID == subjectOpop.ID{
// 				hd.IDSubject = subjectPlan.IDSubject
// 				hd.Variative = subjectPlan.Hoursquantitytotal - subjectOpop.Totalhours
// 				hoursDifference = append(hoursDifference, hd)
// 			}
// 		}
// 	}

// 	json.NewEncoder(w).Encode(hoursDifference)

// 	// hl := subplan.Hoursquantitytotal - opop.Totalhours

// 	// if hl >= 0 {
// 	// 	hoursdifferense.Hoursleft = hl
// 	// 	json.NewEncoder(w).Encode(Hoursleft)
// 	// } else {
// 	// 	json.NewEncoder(w).Encode(struct{ Result string }{Result: fmt.Sprintf("hours differense = %d",hl) })

// 	// }
// 	return

// }
