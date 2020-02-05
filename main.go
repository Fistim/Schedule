package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	// "encoding/json"
	// "strings"
	// "strconv"
	"time"

	"./api"
)

const (
	HTMLPATH    = "html/"
	PATH_STATIC = "static/"
)

var classrooms []api.Classroom

func main() {
	// http.HandleFunc("/", indexPage)

	// db, err := sql.Open("mysql", "root:QWEasd123@/Schedule")

	// if err!=nil{
	// 	fmt.Println("error connecting to mysql")
	// 	return
	// }

	classrooms = append(classrooms, api.Classroom{ID: "402", PlaceQuantity: 16, IsComputer: true})
	classrooms = append(classrooms, api.Classroom{ID: "404", PlaceQuantity: 18, IsComputer: false})

	api.Groups[471] = api.Group{ID: 471, SpecialtyId: 2, StudentsQuantity: 26}
	api.Teachers[0] = api.Teacher{ID: 0, Name: "Test1"}
	api.Teachers[1] = api.Teacher{ID: 1, Name: "Test2"}
	api.Teachers[2] = api.Teacher{ID: 2, Name: "Test3"}

	http.Handle("/static/", http.StripPrefix(
		"/static/",
		handleFileServer(http.Dir(PATH_STATIC))),
	)

	// index := http.FileServer(http.Dir("./html"))

	http.HandleFunc("/api/teacher/", api.TeacherById)
	http.HandleFunc("/api/teacher", api.GetTeacher)
	http.HandleFunc("/api/classroom/", api.ClassroomById)
	http.HandleFunc("/api/classroom", api.GetClassroom)
	http.HandleFunc("/api/classroom/computer", api.GetClassroomByComputer)
	http.HandleFunc("/api/classroom/lecture", api.GetClassroomByComputer)
	http.HandleFunc("/api/group", api.GetGroup)
	http.HandleFunc("/api/group/", api.GroupById)
	http.HandleFunc("/api/group/specialty/", api.GetGroupBySpecialty)
	http.HandleFunc("/api/schedule/generate", api.Generate)
	http.HandleFunc("/api/subject", api.GetSubject)

	http.HandleFunc("/api/shutdown/", api.Shutdown)

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/plan/", planPage)
	http.HandleFunc("/teacher/", teacherPage)
	http.HandleFunc("/classroom/", classroomPage)
	http.HandleFunc("/groupcard/", groupCardPage)
	http.HandleFunc("/teachercard/", teacherCardPage)
	http.HandleFunc("/group/", groupPage)
	http.HandleFunc("/cycle/", cyclePage)
	http.HandleFunc("/shutdown/", shutdown)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8888", nil)
}

// type Teacher struct{
// 	ID int `json:"id"`
// 	Name string `json:"name"`
// }

func showRequest(r *http.Request) {
	t := time.Now()
	fmt.Println("[", t.Format(time.Kitchen), "]", "Request from ", r.RemoteAddr, " to ", r.URL.Path)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println("[", t.Format(time.Kitchen), "]", "Shutdown called by: ", r.RemoteAddr)
	os.Exit(0)
}

func handleFileServer(fs http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := fs.Open(r.URL.Path); os.IsNotExist(err) {
			http.Redirect(w, r, "/404Err", 404)
			return
		}
		http.FileServer(fs).ServeHTTP(w, r)
	})
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/plan/", 302)
}

func planPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/plan/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"index.html",
		HTMLPATH+"nav.html",
		HTMLPATH+"head.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func groupPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/group/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"group.html",
		HTMLPATH+"nav.html",
		HTMLPATH+"head.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func groupCardPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/groupcard/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"groupcard.html",
		HTMLPATH+"nav.html",
		HTMLPATH+"head.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func teacherCardPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/teachercard/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"teacherscard.html",
		HTMLPATH+"nav.html",
		HTMLPATH+"head.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func teacherPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/teacher/" {
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"teacher.html",
		HTMLPATH+"nav.html",
		HTMLPATH+"head.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func classroomPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/classroom/" {
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"classroom.html",
		HTMLPATH+"nav.html",
		HTMLPATH+"head.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func cyclePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cycle/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"cycle.html",
		HTMLPATH+"nav.html",
		HTMLPATH+"head.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}
