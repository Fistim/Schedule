package main

import (
	"./color"
	"./settings"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/auth/", 302)
}

func planPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/plan/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"index.html",
		NAVPATH,
		HEADPATH,
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/auth/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"login.html",
		HEADPATH,
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func specialtyPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/specialty/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"specialty.html",
		NAVPATH,
		HEADPATH,
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func opopPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/opop/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"opop.html",
		NAVPATH,
		HEADPATH,
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func bellPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/bellschedule/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"bellschedule.html",
		NAVPATH,
		HEADPATH,
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func schedulePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/schedule/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"schedule.html",
		NAVPATH,
		HEADPATH,
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func apiPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/api/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH + "api.html",
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func subjectPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/subject/" {
		http.Redirect(w, r, "/404Err", 404)
		return
	}

	t, err := template.ParseFiles(
		HTMLPATH+"subject.html",
		NAVPATH,
		HEADPATH,
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
		NAVPATH,
		HEADPATH,
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
		NAVPATH,
		HEADPATH,
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
		NAVPATH,
		HEADPATH,
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
		NAVPATH,
		HEADPATH,
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
		NAVPATH,
		HEADPATH,
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
		NAVPATH,
		HEADPATH,
	)

	if err != nil {
		panic("can't load hmtl files")
	}
	showRequest(r)
	t.Execute(w, nil)
}

func showRequest(r *http.Request) {
	t := time.Now()
	fmt.Println("[", t.Format(settings.TimeLayout), "]", color.Blue, "Request", color.Reset, "from", r.RemoteAddr, " to ", r.URL.Path)

	urllog, err := os.OpenFile("log/Requests_URL.log", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("[", t.Format(settings.TimeLayout), "]", "Error writing to URL log file")
	}
	s := "[" + t.Format(settings.TimeLayout) + "] " + "Request from " + r.RemoteAddr + " to " + r.URL.Path + "\n"
	_, err = urllog.WriteString(s)
	if err != nil {
		fmt.Println("[", t.Format(settings.TimeLayout), "]", "Error writing to log file")
	}
}
