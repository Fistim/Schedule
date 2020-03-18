package api

import (
	"../settings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// "strings"
)

var attestations []Attestation

func AttestationRoute(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "UPDATE" || r.Method == "PATCH" {
		UpdateAttestation(w, r)
		return
	}

	path := replacePath(r.URL.Path, "/api/attestation/")

	if path == "" {
		GetAttestations(w, r)
		return
	}

	num, err := strconv.Atoi(path)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}
	settings.DB.Find(&attestations)
	if num < 0 || num > len(attestations)+1 {
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "strconv error"})
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		GetAttestationById(w, r, num)
		return
	}

	json.NewEncoder(w).Encode(struct{ Error string }{Error: "This method is not supported"})
}

func UpdateAttestation(w http.ResponseWriter, r *http.Request) {
	var attestation Attestation
	var AttestationToUpdate Attestation
	err := json.NewDecoder(r.Body).Decode(&attestation)

	if err != nil {
		showError(r, err)
		json.NewEncoder(w).Encode(struct{ Error string }{Error: "error update Attestation"})
		return
	}
	settings.DB.First(&AttestationToUpdate, attestation.ID)
	settings.DB.Model(&AttestationToUpdate).Updates(attestation)
	json.NewEncoder(w).Encode(struct{ Result string }{Result: "updated Attestation"})
	return
}

func GetAttestationById(w http.ResponseWriter, r *http.Request, num int) {
	var Attestation Attestation
	settings.DB.Where("id = ?", num).First(&Attestation)
	json.NewEncoder(w).Encode(Attestation)
	return
}

func GetAttestations(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var Attestations []Attestation
		settings.DB.Find(&Attestations)
		json.NewEncoder(w).Encode(Attestations)
	}
	return
}
