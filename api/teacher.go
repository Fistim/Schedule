package api

import(
	"encoding/json"
	"strings"
	"strconv"
	"net/http"
)

var Teachers = make(map[uint8]Teacher)

func CreateTeacher(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err!=nil{
		//TODO
		return
	}

	teacher.ID = uint8(len(Teachers))
	Teachers[teacher.ID] = teacher
	json.NewEncoder(w).Encode(teacher)
}

func TeacherById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if(r.Method=="POST"){
		var newTeacher Teacher
		err:=json.NewDecoder(r.Body).Decode(&newTeacher)
		if err!=nil{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error add new teacher"})
			return
		}
		Teachers[uint8(len(Teachers))] = newTeacher
		json.NewEncoder(w).Encode(struct{Result string}{Result: "added new teacher"})
		return
	}

	if(r.Method=="UPDATE" || r.Method=="PATCH"){
		var newTeacher Teacher
		err:=json.NewDecoder(r.Body).Decode(&newTeacher)
		if err!=nil{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update teacher"})
			return
		}
		if _, ok:=Teachers[newTeacher.ID];!ok{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update teacher"})
			return
		}
		Teachers[newTeacher.ID]=newTeacher
		json.NewEncoder(w).Encode(struct{Error string}{Error:"Successfully updated teacher"})
		return
	}

	path:=r.URL.Path

	path=strings.Replace(path, "/api/teacher/", "", 1)

	num, err := strconv.Atoi(path)

	if err!=nil{
		json.NewEncoder(w).Encode(struct{Error string}{Error:"strconv error"})
		return
	}

	if num<0 || num>len(Teachers){
		json.NewEncoder(w).Encode(struct{Error string}{Error:"strconv error"})
		return
	}

	if r.Method=="GET"{
		json.NewEncoder(w).Encode(Teachers[uint8(num)])
		return
	}

	
	json.NewEncoder(w).Encode(struct{Error string}{Error:"It is not GET method"})
}

func GetTeacher(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method=="GET"{
		json.NewEncoder(w).Encode(Teachers)
	}
		
}