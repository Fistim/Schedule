package api

import(
	"encoding/json"
	"strings"
	// "strconv"
	"net/http"
)

var Classrooms = make(map[string]Classroom)

func CreateClassroom(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var Classroom Classroom
	err := json.NewDecoder(r.Body).Decode(&Classroom)
	if err!=nil{
		//TODO
		return
	}

	Classrooms[Classroom.ID] = Classroom
	json.NewEncoder(w).Encode(Classroom)
}

func ClassroomById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method=="POST"{
		var newClassroom Classroom
		err:=json.NewDecoder(r.Body).Decode(&newClassroom)
		if err!=nil{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error add new Classroom"})
			return
		}
		Classrooms[newClassroom.ID] = newClassroom
		json.NewEncoder(w).Encode(struct{Result string}{Result: "added new Classroom"})
		return
	}

	if(r.Method=="UPDATE" || r.Method=="PATCH"){
		var newClassroom Classroom
		err:=json.NewDecoder(r.Body).Decode(&newClassroom)
		if err!=nil{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update Classroom"})
			return
		}
		if _, ok:=Classrooms[newClassroom.ID];!ok{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update Classroom"})
			return
		}
		Classrooms[newClassroom.ID]=newClassroom
		json.NewEncoder(w).Encode(struct{Error string}{Error:"Successfully updated Classroom"})
		return
	}

	path:=r.URL.Path

	path=strings.Replace(path, "/api/Classroom/", "", 1)

	if _, ok:=Classrooms[path];!ok{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update Classroom"})
			return
		}

	if r.Method=="GET"{
		json.NewEncoder(w).Encode(Classrooms[path])
		return
	}

	
	json.NewEncoder(w).Encode(struct{Error string}{Error:"It is not GET method"})
}

func GetClassroom(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method=="GET"{
		json.NewEncoder(w).Encode(Classrooms)
	}
}
		
func GetClassroomByComputer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method=="GET"{
		path:=r.URL.Path
		path=strings.Replace(path, "/api/classroom/", "", 1)
		var output []Classroom
		for _, room := range Classrooms{
			if room.IsComputer && path=="computer"{
				output = append(output, room)
			}
			if !room.IsComputer && path=="lecture"{
				output = append(output, room)
			}
		}
		json.NewEncoder(w).Encode(output)
	}
}