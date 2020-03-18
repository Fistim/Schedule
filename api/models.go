package api





type Professionalmodule struct {
	ID        uint
	Name      string
	Shortname string
	IDCycle   uint
}

type Cycle struct {
	ID        uint
	Name      string
	Shortname string
}

type Lesson struct {
	ID        uint
	Timestart string
	Timeend   string
}

type Attestation struct {
	ID        uint
	Name      string
	Shortname string
}

type Subjectofgroup struct {
	ID            uint
	IDSubject     uint
	IDGroup       uint
	IDTeacher     uint
	IDSemester    uint
	IDAttestation uint
	Hoursquantity uint
}

type Subject struct {
	ID                    uint
	Name                  string
	Shortname             string
	ID_professionalmodule uint
	IDStype				  uint
}

type Group struct {
	ID          uint
	IDSpecialty uint
	Year        uint
	IDTeacher   uint
	Groupnumber string 
	Studentsquantity uint
	Isbudget bool
	IDPlan uint
}

type Semester struct {
	ID            uint
	Weeksquantity uint
}


type Exampleprogram struct {
	ID            uint
	IDSpecialty   uint
}

type Teacher struct {
	ID          uint
	Name        string
	IDClassroom uint
	Surname     string
	Patronymic  string
}

type Scheduleofgroup struct {
	ID             uint
	IDSubject      uint
	IDGroup        uint
	IDSchedule     uint
	IDLessonnumber uint
	IDClassroom    uint
}

type Schedule struct {
	ID         uint
	IsShort    bool
	IDGroup    uint
	IsEvenWeek bool
	Weekday    string
}

type Classroom struct {
	ID            uint
	Placequantity uint
	Iscomputer    bool
	IDBuilding    uint
	Name          string
}

type Building struct {
	ID      uint
	Address string
}

type Specialty struct {
	ID         uint
	Code       string
	Name       string
	IDDuration uint
}

type DurationOfStudy struct {
	ID            uint
	Yearsquantity uint
}


type Subjecttype struct {
	ID            uint
	Name 		  string
}


type Subjectofplan struct {
	ID            uint
	IDPlan		  uint
	IDSubject	  uint
	Hoursquantitytotal uint
	independentwork	  uint
	consulthours  uint
}


type Studyplan struct {
	ID            uint
	IDSpecialty   uint
	IDGroup       uint
	Isarchive 	  bool
}


type Subjectofexample struct {
	ID            uint
	IDSubject     uint
	IDExampleprogram     uint
	Totalhours uint
	Prefferedcourse string
	Labhours uint
	Practicehours uint

}

type HourseComparison struct{
	Hoursleft int

}