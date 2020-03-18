package schedule

type TypeRow uint8
const (
	IsCycleTR = 0
	IsStartTR = 1
	IsPmTR    = 2
	IsSubTR   = 3
)

type PlanXLSX struct {
	ProgramCode   string     `json:"program_code"`
	NameCode      string     `json:"name_code"`
	PeriodOfStudy string     `json:"period_of_study"`
	GroupNumber   string     `json:"group_number"`
	Lines         []LineXLSX `json:"lines"`
}

type LineXLSX struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	CertForm  string        `json:"cert_form"`
	StudyLoad StudyLoadXLSX `json:"study_load"`
	Course    [4][2]int     `json:"course"`
	TypeRow   TypeRow       `json:"typerow"`
}

type StudyLoadXLSX struct {
	Max       int `json:"max"`
	SelfStudy int `json:"self_study"`
	AllStudy  int `json:"all_study"`
	Lectures  int `json:"lectures"`
	Labs      int `json:"labs"`
	Projects  int `json:"projects"`
}

type DayType uint8
const (
	SUNDAY 		DayType = 0
	MONDAY 		DayType = 1
	TUESDAY 	DayType = 2
	WEDNESDAY 	DayType = 3
	THURSDAY 	DayType = 4
	FRIDAY 		DayType = 5
	SATURDAY 	DayType = 6
)

type SubgroupType uint8
const (
	A 	SubgroupType = 0
	B 	SubgroupType = 1
	ALL SubgroupType = 2
)

type SubjectType uint8
const (
	THEORETICAL	SubjectType = 0
	PRACTICAL 	SubjectType = 1
)

type Generator struct {
	Day      DayType
	Debug    bool
	Groups   map[string]*Group
	Teachers map[string]*Teacher
	Cabinets map[string]*Cabinet
	// Blocked map[string]bool
	Reserved Reserved
}

type Reserved struct {
	Teachers map[string][]bool
	Cabinets map[string][]bool
}

type Schedule struct {
	Day   DayType
	Group string
	Table []Row
}

type Row struct {
	Subject [ALL]string
	Teacher [ALL]string
	Cabinet [ALL]string
}

type Teacher struct {
	Name     string    `json:"name"`
	Cabinets []Cabinet `json:"cabinets"`
}

type Cabinet struct {
	Name       string `json:"name"`
	IsComputer bool   `json:"is_computer"`
}

type Group struct {
	Name     string
	Quantity uint // students count
	Subjects map[string]*Subject
}

type Subject struct {
	Name       string
	IsComputer bool
	Teacher    string
	Teacher2   string
	SaveWeek   uint
	Lessons    Lessons
}

type Lessons struct {
	Theory   uint
	Practice Subgroup
	Week     Subgroup
}

type Subgroup struct {
	A uint
	B uint
}

type GroupJSON struct {
	Name     string `json:"name"`
	Quantity uint `json:"quantity"`
	Subjects []SubjectJSON `json:"subjects"`
}

type SubjectJSON struct {
	Name       string `json:"name"`
	Teacher    string `json:"teacher"`
	IsComputer bool `json:"is_computer"`
	Lessons    LessonsJSON `json:"lessons"`
}

type LessonsJSON struct {
	Theory   uint `json:"theory"`
	Practice uint `json:"practice"`
	Week     uint `json:"week"`
}
