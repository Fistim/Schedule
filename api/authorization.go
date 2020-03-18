package api

import(
	"gopkg.in/korylprince/go-ad-auth.v2"
	"net/http"
	"encoding/json"
)

type User struct{
	Login string `json:"login"`
	Password string `json:"password"`
}

func LoginRoute(w http.ResponseWriter, r *http.Request){
	showAPIRequest(r)
	var userGroups []string
	Groups := []string{"Студенты", "Персонал", "Бухгалтерия", "Преподаватели", "Админы"}
	if r.Method == "POST"{
		config := &auth.Config{
		Server:   "192.168.10.121",
        Port:     389,
        BaseDN:   "DC=ttit,DC=local",
        Security: auth.SecurityNone,
		}

		var user User
		error := json.NewDecoder(r.Body).Decode(&user)
		if error != nil{
			json.NewEncoder(w).Encode(struct{ Error string }{Error: "an error has occured during decoding"})
			showError(r, error)
			return
		}
		status, _, groups, err := auth.AuthenticateExtended(config, user.Login, user.Password, []string{"cn"}, Groups)
		if err != nil {
	   	 	json.NewEncoder(w).Encode(struct{ Error string }{Error: "an error has occured"})
	   	 	showError(r, err)
	   	 	return
		} else if !status {
    		json.NewEncoder(w).Encode(struct{ Error string }{Error: "no user found"})
    		return
		} 
		for _, group := range groups{
			for _, constGroup := range Groups{
				if group == constGroup{
					userGroups = append(userGroups, group)
				}
			}
		}
		json.NewEncoder(w).Encode(userGroups)
		return
	}
}