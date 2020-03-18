package api

import (
	"../settings"
	"encoding/json"
	"fmt"
    "net/http"
    "../schedule"
	// "strconv"
	// "errors"
	// "math/rand"
)

var tmpteachers []schedule.Teacher
var tmpcabinets []schedule.Cabinet

var tmpgroups []schedule.Group
var tmpsubjects []schedule.Subject
// var tmplessons []schedule.LessonJSON

func GetDataTeachers(w http.ResponseWriter, r *http.Request) {
	tmpteachers = nil
	tmpcabinets = nil
	var classrooms []Classroom
	var teachers []Teacher
	settings.DB.Find(&teachers)
	settings.DB.Find(&classrooms)
	var cabinet schedule.Cabinet
	for _, room := range classrooms{
		cabinet.Name = room.Name
		cabinet.IsComputer = room.Iscomputer
		tmpcabinets = append(tmpcabinets, cabinet)
	}

	for _, teacher := range teachers{
		var scheduleTeacher schedule.Teacher
		scheduleTeacher.Name = teacher.Name + " " + string(teacher.Surname[:2]) + "." + string(teacher.Patronymic[:2]) + "."
		for _, room := range classrooms{
			if teacher.IDClassroom == room.ID {
				var tmpcabinet schedule.Cabinet
				tmpcabinet.Name = room.Name
				tmpcabinet.IsComputer = room.Iscomputer
				scheduleTeacher.Cabinets = append(scheduleTeacher.Cabinets, tmpcabinet)
				tmpcabinet = schedule.Cabinet{}
				break 
			}
		}
		tmpteachers = append(tmpteachers, scheduleTeacher) 
	}
	 jsondata, err := json.MarshalIndent(tmpteachers, "", "\t")
	 if err!= nil {
 		println(err)
	 }
	settings.WriteFile("input/teachers.json", string(jsondata))
	json.NewEncoder(w).Encode(tmpteachers)

}

func GetDataGroups(w http.ResponseWriter, r *http.Request){
	
	var groups []Group
	var subjects []Subject
	var tmpsubofgroup []Subjectofgroup
	
	settings.DB.Find(&groups)
	settings.DB.Find(&subjects)


	for _, group := range groups{
		settings.DB.Where("id_group = ?", group.ID).Find(&tmpsubofgroup)

		var tmpgroup schedule.Group
		tmpgroup.Name = group.Groupnumber
		tmpgroup.Quantity = group.Studentsquantity

		for _, subject := range tmpsubofgroup{
			var tmpsubject schedule.Subject
			var tsubject Subject
			var teacher Teacher
			var semester Semester
			settings.DB.Where("id = ?", subject.ID).First(&tsubject)
			settings.DB.Where("id = ?", subject.IDTeacher).First(&teacher)
			settings.DB.Where("id = ?", subject.IDSemester).First(&semester)
			tmpsubject.Name = tsubject.Shortname
			tmpsubject.Teacher = teacher.Name + " " + string(teacher.Surname[:2]) + "." + string(teacher.Patronymic[:2]) + "."
			// tmpsubject.Teacher2 = tmpsubject.Teacher // TODO
			if tsubject.IDStype == 1{
				tmpsubject.IsComputer = false
				// tmpsubject.Theory = 0
				// tmpsubject.Practice.A = subject.Hoursquantity/2
				// tmpsubject.Practice.B = subject.Hoursquantity/2
			} else if tsubject.IDStype == 0{
				tmpsubject.IsComputer = true
				// tmpsubject.Theory = subject.Hoursquantity/2
				// tmpsubject.Practice.A = 0
				// tmpsubject.Practice.B = 0
			}
			fmt.Println(semester)
			// tmpsubject.WeekLessons.A = (subject.Hoursquantity/semester.Weeksquantity)/2
			// tmpsubject.WeekLessons.B = (subject.Hoursquantity/semester.Weeksquantity)/2
			tmpsubjects = append(tmpsubjects, tmpsubject)
		}

		
	}

	
	// jsondata, err := json.MarshalIndent(tmpteachers, "", "\t")
	//  if err!= nil {
 // 		println(err)
	//  }
	// settings.WriteFile("input/groups.json", string(jsondata))

	json.NewEncoder(w).Encode(tmpgroups)
	
}
