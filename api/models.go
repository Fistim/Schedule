package api

// import "time"

type BellSchedule struct {
	ID        uint8  `json:"id"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"TimeEnd"`
}

type Classroom struct {
	ID            string `json: "id"`
	PlaceQuantity uint8  `json:"placeQuantity"`
	IsComputer    bool   `json:"isComputer"`
}

type Group struct {
	ID               uint16 `json:"id"`
	SpecialtyId      uint8  `json:"specialty"`
	StudentsQuantity uint8  `json:"studentsQuantity"`
}

type Subject struct {
	ID        uint8  `json:"id"`
	Name      string `json:"name"`
	ID_Module uint8  `json:"module"`
}

type Teacher struct {
	ID        uint8     `json:"id"`
	Name      string    `json:"name"`
	Classroom Classroom `json:"classroom"`
}

type Specialty struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
}

type ProfessionalModule struct {
	ID        uint8  `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type TrainingPractices struct {
	ID   uint8 `json:"id"`
	Name uint8 `json:"id"`
}

type PracticeOfGroup struct {
	Group            Group             `json:"group"`
	TrainingPractice TrainingPractices `json:"trainingPractice"`
	HoursQuantity    uint8             `json:"hoursQuantity"`
}

type SubjectOfGroup struct {
	Group         Group   `json:"group"`
	Subject       Subject `json:"subject"`
	HoursQuantity uint16  `json:"hoursQuantity"`
	Teacher       Teacher `json:"teacher"`
}

type Schedule struct {
	ID        uint8          `json:"id"`
	Date      string         `json:"date"`
	Subject   SubjectOfGroup `json:"subject"`
	Classroom Classroom      `json:"classroom"`
}
