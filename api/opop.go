package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var opop []Exampleprogram

func OpopRoute(w http.ResponseWriter, r *http.Request) {
	
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		CreateOpop(w, r)
		return
	}

	
	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateOpop(w, r)
		return
	}

	
	path := replacePath(r.URL.Path, "/api/opop/")

	if path == "" {
		GetOpop(w, r)
		return
	}
	

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&opop)
	if num < 0 || num > len(opop)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetOpopById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateOpop(w http.ResponseWriter, r *http.Request) {
	var Opop Exampleprogram
	var OpopToUpdate Exampleprogram
	err := json.NewDecoder(r.Body).Decode(&Opop)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Opop"})
		return
	}
	settings.DB.First(&OpopToUpdate, Opop.ID)
	settings.DB.Model(&OpopToUpdate).Updates(Opop)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated Opop"})
	return
}

func CreateOpop(w http.ResponseWriter, r *http.Request) {
	var newOpop Exampleprogram
	err := json.NewDecoder(r.Body).Decode(&newOpop)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error add new Opop"})
		fmt.Println(err)
		return
	}
	settings.DB.Create(&newOpop)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "added new Opop"})
	return
}

func GetOpopById(w http.ResponseWriter, r *http.Request, num int) {
	var Opop Exampleprogram
	settings.DB.Where("id = ?", num).First(&Opop)
	json.NewEncoder(w).Encode(Opop)
	return
}

func GetOpop(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var Opop []Exampleprogram
		settings.DB.Find(&Opop)
		json.NewEncoder(w).Encode(Opop)
	}
	return
}
