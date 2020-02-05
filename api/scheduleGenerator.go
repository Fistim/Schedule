package api

import(
	"../schedule"
	"encoding/json"
	"net/http"
	"fmt"
)

const (
	INDATA = "input/"
)

func Generate(w http.ResponseWriter, r *http.Request){
	if r.Method=="GET"{
		var result [][]*schedule.Schedule
		var generator = schedule.NewGenerator(&schedule.Generator{
			Day: schedule.FRIDAY,
			NumTables: 11,
			Groups: schedule.ReadGroups(INDATA + "groups.json"),
			Teachers: schedule.ReadTeachers(INDATA + "teachers.json"),
		})
		for iter := 1; iter <= 5; iter++ {
			result = append(result, generator.Generate())
		}
		json.NewEncoder(w).Encode(result)
	}
}

func printJSON(data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(jsonData))
}
