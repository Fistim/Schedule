package main

import (
	"./api"
	"./settings"
	"fmt"
	"net/http"
	"os"
	"time"
	// "./schedule"
	 "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	HTMLPATH    = "html/"
	PATH_STATIC = "static/"
	NAVPATH     = "html/nav.html"
	HEADPATH    = "html/head.html"
)

func init() {
	var err error
	settings.DB, err = gorm.Open("mysql", settings.CONSTR)
	if err != nil {
		fmt.Println(err)
		return
	}
	settings.DB.SingularTable(true)
}

func GetTime() time.Time{
	return time.Now()
}

func main() {
	http.Handle("/static/", http.StripPrefix(
		"/static/",
		handleFileServer(http.Dir(PATH_STATIC))),
	)
	_, err := os.Create("log/Requests_API.log")
	if err != nil {
		fmt.Println("api log file not created")
		return
	}
	t := time.Now()
	fmt.Println("[", t.Format(settings.TimeLayout), "]", "API log file can be found at log/Requests_API.log")

	_, err = os.Create("log/Requests_URL.log")
	if err != nil {
		fmt.Println("url log file not created")
		return
	}
	fmt.Println("[", t.Format(settings.TimeLayout), "]", "URL log file can be found at log/Requests_URL.log")

	_, err = os.Create("log/error.log")
	if err != nil {
		fmt.Println("Error log file not created")
		return
	}
	fmt.Println("[", t.Format(settings.TimeLayout), "]", "Error log file can be found at log/error.log")

	// API functions handling start
	http.HandleFunc("/api/teacher/", api.TeacherRoute)
	http.HandleFunc("/api/classroom/", api.ClassroomRoute)
	http.HandleFunc("/api/classroom/computer", api.GetClassroomByComputer)
	http.HandleFunc("/api/classroom/lecture", api.GetClassroomByComputer)
	http.HandleFunc("/api/classroom/name/", api.ClassroomByNumber)
	http.HandleFunc("/api/group/", api.GroupRoute)
	http.HandleFunc("/api/groupsubject/", api.SubjectOfGroupRoute)
	http.HandleFunc("/api/schedule/generate", api.Generate)
	http.HandleFunc("/api/subject/module/", api.GetSubjectByModule)
	http.HandleFunc("/api/subject/", api.SubjectRoute)
	http.HandleFunc("/api/module/", api.ModuleApi)
	http.HandleFunc("/api/cycle/", api.CycleApi)
	http.HandleFunc("/api/specialty/", api.SpecialtyRoute)
	http.HandleFunc("/api/building/", api.GetBuilding)
	http.HandleFunc("/api/groupschedule/", api.ScheduleOfGroupRoute)
	http.HandleFunc("/api/bellschedule/", api.BellScheduleRoute)
	http.HandleFunc("/api/attestation/", api.AttestationRoute)
	http.HandleFunc("/api/semester/", api.SemesterRoute)
	http.HandleFunc("/api/studyplan/", api.StudyplanRoute)
	http.HandleFunc("/api/subjectofplan/", api.SubjectofplanRoute)
	http.HandleFunc("/api/subjecttype/", api.SubjecttypeRoute)
	http.HandleFunc("/api/logs", testPage)
	http.HandleFunc("/api/getteachers", api.GetDataTeachers)
	http.HandleFunc("/api/getgroups/", api.GetDataGroups)

	http.HandleFunc("/api/opop/", api.OpopRoute)
	http.HandleFunc("/api/subjectofexample/", api.SubjectofexampleRoute)
	http.HandleFunc("/api/importplan/", api.ImportPlanRoute)

	http.HandleFunc("/api/auth/", api.LoginRoute)
	// API functions handling end

	// HTTP pages handling start
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/plan/", planPage)
	http.HandleFunc("/teacher/", teacherPage)
	http.HandleFunc("/classroom/", classroomPage)
	http.HandleFunc("/groupcard/", groupCardPage)
	http.HandleFunc("/schedule/", schedulePage)
	http.HandleFunc("/teachercard/", teacherCardPage)
	http.HandleFunc("/group/", groupPage)
	http.HandleFunc("/cycle/", cyclePage)
	http.HandleFunc("/subject/", subjectPage)
	http.HandleFunc("/specialty/", specialtyPage)
	http.HandleFunc("/bellschedule/", bellPage)
	http.HandleFunc("/api/", apiPage)
	http.HandleFunc("/auth/", loginPage)
	http.HandleFunc("/api/modelsDownload", downloadModels)
	http.HandleFunc("/wekan", wekanRedirect)
	http.HandleFunc("/opop/", opopPage)
	http.HandleFunc("/test/", testPage) // testing token
	// HTTP pages handling end

	api.PrintConsole("Server is listening")

	http.ListenAndServe(":80", nil)
}

func wekanRedirect(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/wekan" || r.URL.Path == "/wekan/"{
		http.Redirect(w, r, "http://192.168.14.173:8000", 301)
		return
	}
}

func testPage(w http.ResponseWriter, r *http.Request) {
	if api.CheckToken(r) {
		fmt.Println("Access granted")
	}
	fmt.Println("Access denied")
}

func downloadModels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "models.txt"))
	http.ServeFile(w, r, "models.txt")
}

func RenameLogFiles() {
	now := time.Now()
	t := now.Format("2 Jan 2006 15:04:05")

	newpath := "log/Requests_API " + t + ".log"
	err := os.Rename("log/Requests_API.log", newpath)

	if err != nil {
		fmt.Println("Error during saving API log file")
	}

	newpath = "log/Requests_URL " + t + ".log"
	err = os.Rename("log/Requests_URL.log", newpath)

	if err != nil {
		fmt.Println("Error during saving URL log file")
	}
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	showRequest(r)
	now := time.Now()
	t := now.Format(time.Stamp)
	fmt.Println("Shutting down...")
	newpath := "log/Requests_API " + t + ".log"
	err := os.Rename("log/Requests_API.log", newpath)

	if err != nil {
		fmt.Println("Error during saving API log file")
	}

	newpath = "log/Requests_URL " + t + ".log"
	err = os.Rename("log/Requests_URL.log", newpath)

	if err != nil {
		fmt.Println("Error during saving URL log file")
	}

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
