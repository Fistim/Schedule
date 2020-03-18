package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../schedule"
)

const (
	INDATA = "input/"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	showAPIRequest(r)
	if r.Method == "GET" {
		var result [][]*schedule.Schedule
		var generator = schedule.NewGenerator(&schedule.Generator{
			Day: schedule.MONDAY,
			// NumTables: 11,
			Groups:   schedule.ReadGroups(INDATA + "groups.json"),
			Teachers: schedule.ReadTeachers(INDATA + "teachers.json"),
		})
		for iter := 1; iter <= 7; iter++ {
			result = append(result, generator.Generate(nil))
		}
		json.NewEncoder(w).Encode(result)
	}
}

func printJSON(data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(jsonData))
}
