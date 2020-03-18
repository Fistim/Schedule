package schedule

import (
    "fmt"
    "encoding/json"
    "github.com/tealeg/xlsx"
)

func NewGenerator(data *Generator) *Generator {
    return &Generator{
        Day: data.Day,
        Debug: data.Debug,
        Groups: data.Groups,
        Teachers: data.Teachers,
        Cabinets: data.Cabinets,
        // Blocked: make(map[string]bool),
        Reserved: Reserved{
            Teachers: make(map[string][]bool),
            Cabinets: make(map[string][]bool),
        },
    }
}

func (gen *Generator) NewSchedule(group string) *Schedule {
    return &Schedule{
        Day: gen.Day,
        Group: group,
        Table: make([]Row, NUM_TABLES),
    }
}

func ReadGroups(filename string) map[string]*Group {
    var (
        groups = make(map[string]*Group)
        groupsList []GroupJSON
    )
    data := readFile(filename)
    err := json.Unmarshal([]byte(data), &groupsList)
    if err != nil {
        return nil
    }
    for _, gr := range groupsList {
        groups[gr.Name] = &Group{
            Name: gr.Name,
            Quantity: gr.Quantity,
        }
        groups[gr.Name].Subjects = make(map[string]*Subject)
        for _, sb := range gr.Subjects {
            if _, ok := groups[gr.Name].Subjects[sb.Name]; ok {
                groups[gr.Name].Subjects[sb.Name].Teacher2 = sb.Teacher
                continue
            }
            groups[gr.Name].Subjects[sb.Name] = &Subject{
                Name: sb.Name,
                Teacher: sb.Teacher,
                IsComputer: sb.IsComputer,
                SaveWeek: sb.Lessons.Week,
                Lessons: Lessons{
                    Theory: sb.Lessons.Theory,
                    Practice: Subgroup{
                        A: sb.Lessons.Practice,
                        B: sb.Lessons.Practice,
                    },
                    Week: Subgroup{
                        A: sb.Lessons.Week,
                        B: sb.Lessons.Week,
                    },
                },
            }
        }
    }
    return groups
}

func ReadTeachers(filename string) map[string]*Teacher {
    var (
        teachers = make(map[string]*Teacher)
        teachersList []Teacher
    )
    data := readFile(filename)
    err := json.Unmarshal([]byte(data), &teachersList)
    if err != nil {
        return nil
    }
    for _, tc := range teachersList {
        teachers[tc.Name] = &Teacher{
            Name: tc.Name,
            Cabinets: tc.Cabinets,
        }
    }
    return teachers
}

func ReadCabinets(filename string) map[string]*Cabinet {
    var (
        cabinets = make(map[string]*Cabinet)
        cabinetsList []Cabinet
    )
    data := readFile(filename)
    err := json.Unmarshal([]byte(data), &cabinetsList)
    if err != nil {
        return nil
    }
    for _, cb := range cabinetsList {
        cabinets[cb.Name] = &Cabinet{
            Name: cb.Name,
            IsComputer: cb.IsComputer,
        }
    }
    return cabinets
}

const (
    OUTDATA = "output/"
)
func (gen *Generator) Template() [][]*Schedule {
    var (
        weekLessons = make([][]*Schedule, 7)
        generator = new(Generator)
        file *xlsx.File
        name string
    )
    unpackJSON(packJSON(gen), generator)
    if gen.Debug {
        file, name = CreateXLSX(OUTDATA + "template.xlsx")
    }
    day := generator.Day
    for i := day; i < day+7; i++ {
        weekLessons[i % 7] = generator.Generate(nil)
        if gen.Debug {
            generator.WriteXLSX(
                file,
                name,
                weekLessons[i % 7],
                int(i % 7),
            )
        }
    }
    return weekLessons
}

func (gen *Generator) Generate(template [][]*Schedule) []*Schedule {
    var (
        list   []*Schedule
        templt []*Schedule
        groups = getGroups(gen.Groups)
    )
    if template == nil {
        templt = nil
    } else {
        templt = template[gen.Day]
    }
    for _, group := range groups {
        var (
            schedule = gen.NewSchedule(group.Name)
            subjects = getSubjects(group.Subjects)
            countLessons = new(Subgroup)
        )
        if gen.Day == SUNDAY {
            list = append(list, schedule)
            for _, subject := range subjects {
                saved := gen.Groups[group.Name].Subjects[subject.Name].SaveWeek
                gen.Groups[group.Name].Subjects[subject.Name].Lessons.Week.A = saved
                gen.Groups[group.Name].Subjects[subject.Name].Lessons.Week.B = saved
            }
            continue
        }
        for _, subject := range subjects {
            switch {
            case gen.haveTheoreticalLessons(subject):
                if gen.Debug {
                    fmt.Println(group.Name, subject.Name, ": not splited THEORETICAL;")
                }
                gen.tryGenerate(ALL, THEORETICAL, group, subject, schedule, countLessons, templt)
            // Практические пары начинаются только после завершения всех теоретических.
            default:
                // Если подгруппа неделимая, тогда провести практику в виде полной пары.
                // Иначе разделить практику на две подгруппы.
                if !gen.withSubgroups(group.Name) {
                    if gen.Debug {
                        fmt.Println(group.Name, subject.Name, ": not splited PRACTICAL;")
                    }
                    gen.tryGenerate(ALL, PRACTICAL, group, subject, schedule, countLessons, templt)
                } else {
                    switch RandSubgroup() {
                    case A:
                        if gen.Debug {
                            fmt.Println(group.Name, subject.Name, ": splited (A -> B);")
                        }
                        gen.tryGenerate(A, PRACTICAL, group, subject, schedule, countLessons, templt)
                        gen.tryGenerate(B, PRACTICAL, group, subject, schedule, countLessons, templt)
                    case B:
                        if gen.Debug {
                            fmt.Println(group.Name, subject.Name, ": splited (B -> A);")
                        }
                        gen.tryGenerate(B, PRACTICAL, group, subject, schedule, countLessons, templt)
                        gen.tryGenerate(A, PRACTICAL, group, subject, schedule, countLessons, templt)
                    }
                }
            }
        }
        list = append(list, schedule)
    }
    gen.Reserved.Teachers = make(map[string][]bool)
    gen.Reserved.Cabinets = make(map[string][]bool)
    gen.Day = (gen.Day + 1) % 7
    return sortSchedule(list)
}

func RandSubgroup() SubgroupType {
    return SubgroupType(random(0, 1))
}

func Load(filename string) *Generator {
    var generator = new(Generator)
    jsonData := readFile(filename)
    err := json.Unmarshal([]byte(jsonData), generator)
    if err != nil {
        return nil
    }
    return generator
}

func (gen *Generator) Dump(filename string) error {
    return writeJSON(filename, gen)
}
