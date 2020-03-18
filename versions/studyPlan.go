package versions

import(
	"../settings"
	"net/http"
	"../api"
	"encoding/json"
	"io/ioutil"
	"os"
)

type plan struct{
	ID uint
	IDPlan		  uint
	IDSubject	  uint
	Hoursquantitytotal uint
	independentwork	  uint
	consulthours  uint
	IDSpecialty uint
}

func getStudyPlan(id uint){
	var subjects []api.Subjectofplan
	var group api.Group
	settings.DB.Where("id_plan = ?", id).Find(&subjects)

	jsonOut, err := json.Marshal(subjects)

	settings.DB.Where("id_plan = ", id).First(&group)
	t:=GetTime()

	os.Mkdir("studyPlan", 0777)

	filename := group.Groupnumber + " " + t.Format(settings.TimeLayout) + " " + t.Format(settings.DateLayout)
	_ = ioutil.WriteFile()
}