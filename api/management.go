package api

import (
	"../color"
	"../settings"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func showAPIRequest(r *http.Request) {
	t := time.Now()
	fmt.Println("[", t.Format(settings.TimeLayout), "] ", color.Green, r.Method, color.Reset, " request from ", r.RemoteAddr, " to ", r.URL.Path)

	urllog, err := os.OpenFile("./log/Requests_API.log", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("[", t.Format(settings.TimeLayout), "]", "Error writing to API log file")
	}
	s := "[" + t.Format(time.Kitchen) + "] " + r.Method + " request from " + r.RemoteAddr + " to " + r.URL.Path + "\n"
	_, err = urllog.WriteString(s)
	if err != nil {
		fmt.Println("[", t.Format(settings.TimeLayout), "]", "Error writing to log file")
	}
}

func showError(r *http.Request, Err error) {
	t := time.Now()
	fmt.Println("[", t.Format(settings.TimeLayout), "] ", color.Red, "Error:", color.Reset, Err.Error())

	errlog, err := os.OpenFile("./log/error.log", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("[", t.Format(settings.TimeLayout), "] ", color.Red, "Error:", color.Reset, Err.Error())
		return
	}

	s := "[" + t.Format(settings.TimeLayout) + "] " + "Error:" + Err.Error()

	_, err = errlog.WriteString(s)

	if err != nil {
		fmt.Println("[", t.Format(settings.TimeLayout), "] ", color.Red, "Error:", color.Reset, Err.Error())
	}
	return
}

func PrintConsole(message string) {
	t := time.Now()
	fmt.Println("[", t.Format(settings.TimeLayout), "] ", message)
}

func printError(Err error) {
	t := time.Now()
	s := "[" + t.Format(settings.TimeLayout) + "] " + "Error:" + Err.Error()
	fmt.Println(s)
	return
}

func replacePath(path string, replacement string) string {
	newPath := strings.Replace(path, replacement, "", 1)
	return newPath
}

func CheckMethod(r *http.Request, method string) bool {
	if r.Method == method {
		return true
	}
	return false
}

// func PrintToFile(){

// }
