package api

import(
	"encoding/json"
	"strings"
	"strconv"
	"net/http"
)

var Groups = make(map[uint16]Group)

func CreateGroup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var Group Group
	err := json.NewDecoder(r.Body).Decode(&Group)
	if err!=nil{
		//TODO
		return
	}

	Groups[Group.ID] = Group
	json.NewEncoder(w).Encode(Group)
}

func GroupById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method=="POST"{
		var newGroup Group
		err:=json.NewDecoder(r.Body).Decode(&newGroup)
		if err!=nil{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error add new Group"})
			return
		}
		Groups[newGroup.ID] = newGroup
		json.NewEncoder(w).Encode(struct{Result string}{Result: "added new Group"})
		return
	}

	if(r.Method=="UPDATE" || r.Method=="PATCH"){
		var newGroup Group
		err:=json.NewDecoder(r.Body).Decode(&newGroup)
		if err!=nil{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update Group"})
			return
		}
		if _, ok:=Groups[newGroup.ID];!ok{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update Group"})
			return
		}
		Groups[newGroup.ID]=newGroup
		json.NewEncoder(w).Encode(struct{Error string}{Error:"Successfully updated Group"})
		return
	}

	path:=r.URL.Path

	path=strings.Replace(path, "/api/group/", "", 1)

	num, err:=strconv.Atoi(path)

		if err!=nil || num<0{
			json.NewEncoder(w).Encode(struct{Error string}{Error: "strconv error"})
		}

	if _, ok:=Groups[uint16(num)];!ok{
			json.NewEncoder(w).Encode(struct{Error string}{Error:"error update Group"})
			return
		}

	if r.Method=="GET"{
		json.NewEncoder(w).Encode(Groups[uint16(num)])
		return
	}

	
	json.NewEncoder(w).Encode(struct{Error string}{Error:"It is not GET method"})
}

func GetGroup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method=="GET"{
		json.NewEncoder(w).Encode(Groups)
	}
}

func GetGroupBySpecialty(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method=="GET"{
		path:=r.URL.Path
		path=strings.Replace(path, "/api/group/specialty/", "", 1)
		var output []Group
		num, err:=strconv.Atoi(path)

		if err!=nil || num<0{
			json.NewEncoder(w).Encode(struct{Error string}{Error: "strconv error"})
		}

		for _, group := range Groups{
			if group.SpecialtyId==uint8(num){
				output = append(output, group)
			}
		}
		json.NewEncoder(w).Encode(output)
	}
}