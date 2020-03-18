package schedule

import (
    // "fmt"
    "strings"
    "strconv"
    "github.com/tealeg/xlsx"
)

func ImportXLSX(filename string) []*PlanXLSX {
    var (
        listOfPlans []*PlanXLSX
    )
    xlFile, err := xlsx.OpenFile(filename)
    if err != nil {
        return nil
    }
    for _, sheet := range xlFile.Sheets {
        var plan = new(PlanXLSX)
        plan.ProgramCode   = sheet.Rows[5].Cells[1].String()
        plan.NameCode      = strings.Split(sheet.Rows[5].Cells[6].String(), " ")[0]
        plan.PeriodOfStudy = sheet.Rows[7].Cells[6].String()
        splited := strings.Split(sheet.Rows[9].Cells[1].String(), " ")
        plan.GroupNumber   = strings.Split(splited[len(splited)-1], "-")[0]

        for i, row := range sheet.Rows {
            if i < 20 {
                continue
            }
            if len(row.Cells) < 20 {
                continue
            }
            if row.Cells[2].String() == "" || row.Cells[2].String() == "ИТОГО" {
                continue
            }

            if strings.HasPrefix(row.Cells[1].String(), "Консультации") || strings.HasPrefix(row.Cells[1].String(), "Государственная") {
                continue
            }

            splited := strings.Split(row.Cells[1].String(), ".")
            if len(splited) != 2 {
                continue
            }

            typerow := TypeRow(IsSubTR)
            switch {
            case strings.Contains(row.Cells[2].String(), "цикл"):
                typerow = IsCycleTR
            case splited[1] == "00":
                typerow = IsStartTR
            case splited[0] == "ПМ":
                typerow = IsPmTR
            }

            max, _       := row.Cells[4].Int()
            selfstudy, _ := row.Cells[5].Int()
            allstudy, _  := row.Cells[6].Int()
            lectures, _  := row.Cells[7].Int()
            labs, _      := row.Cells[8].Int()
            projects, _  := row.Cells[9].Int()
            c1s1, _      := row.Cells[12].Int()
            c1s2, _      := row.Cells[13].Int()
            c2s1, _      := row.Cells[14].Int()
            c2s2, _      := row.Cells[15].Int()
            c3s1, _      := row.Cells[16].Int()
            c3s2, _      := row.Cells[17].Int()
            c4s1, _      := row.Cells[18].Int()
            c4s2, _      := row.Cells[19].Int()

            certform := ""
            splited = strings.Split(row.Cells[3].String(), ",")
            for _, v := range splited {
                if v != "-" {
                    certform = v
                    break
                }
            }

            plan.Lines = append(plan.Lines, LineXLSX{
                ID:        row.Cells[1].String(),
                Name:      row.Cells[2].String(),
                CertForm:  certform,
                StudyLoad: StudyLoadXLSX{
                    Max:       max,
                    SelfStudy: selfstudy,
                    AllStudy:  allstudy,
                    Lectures:  lectures,
                    Labs:      labs,
                    Projects:  projects,
                },
                Course: [4][2]int{
                    [2]int{
                        c1s1,
                        c1s2,
                    },
                    [2]int{
                        c2s1,
                        c2s2,
                    },
                    [2]int{
                        c3s1,
                        c3s2,
                    },
                    [2]int{
                        c4s1,
                        c4s2,
                    },
                },
                TypeRow: typerow,
            })
        }
        listOfPlans = append(listOfPlans, plan)
    }
    return listOfPlans
}

func CreateXLSX(filename string) (*xlsx.File, string) {
    file := xlsx.NewFile()
    _, err := file.AddSheet("Init")
    if err != nil {
        return nil, ""
    }
    err = file.Save(filename)
    if err != nil {
        return nil, ""
    }
    return file, filename
}

func (gen *Generator) WriteXLSX(file *xlsx.File, filename string, schedule []*Schedule, iter int) error {
    const (
        colWidth = 30
        rowHeight = 30
        MAXCOL = 3
    )

    rowsNext := uint(len(schedule)) / MAXCOL
    if rowsNext == 0 || uint(len(schedule)) % MAXCOL != 0 {
        rowsNext += 1
    }

    var (
        colNum = uint(NUM_TABLES + 2)
        row = make([]*xlsx.Row, colNum * rowsNext) //  * (rowsNext + 1)
        cell *xlsx.Cell
        dayN = gen.Day
        day = ""
    )

    if dayN == SUNDAY {
        dayN = SATURDAY
    } else {
        dayN -= 1
    }

    switch dayN {
    case SUNDAY: day = "Sunday"
    case MONDAY: day = "Monday"
    case TUESDAY: day = "Tuesday"
    case WEDNESDAY: day = "Wednesday"
    case THURSDAY: day = "Thursday"
    case FRIDAY: day = "Friday"
    case SATURDAY: day = "Saturday"
    }

    sheet, err := file.AddSheet(day + "-" + strconv.Itoa(iter))
    if err != nil {
        return err
    }

    sheet.SetColWidth(2, int(MAXCOL)*3+1, COL_W)

    for r := uint(0); r < rowsNext; r++ {
        for i := uint(0); i < colNum; i++ {
            row[(r*colNum)+i] = sheet.AddRow() // (r*rowsNext)+
            row[(r*colNum)+i].SetHeight(ROW_H)
            cell = row[(r*colNum)+i].AddCell()
            if i == 0 {
                cell.Value = "Пара"
                continue
            }
            cell.Value = strconv.Itoa(int(i-1))
        }
    }

    index := uint(0)
    exit: for r := uint(0); r < rowsNext; r++ {
        for i := uint(0); i < MAXCOL; i++ {
            if uint(len(schedule)) <= index {
                break exit
            }

            savedCell := row[(r*colNum)+0].AddCell()
            savedCell.Value = "Группа " + schedule[index].Group

            cell = row[(r*colNum)+0].AddCell()
            cell = row[(r*colNum)+0].AddCell()

            savedCell.Merge(2, 0)

            cell = row[(r*colNum)+1].AddCell()
            cell.Value = "Предмет"

            cell = row[(r*colNum)+1].AddCell()
            cell.Value = "Преподаватель"

            cell = row[(r*colNum)+1].AddCell()
            cell.Value = "Кабинет"

            for j, trow := range schedule[index].Table {
                cell = row[(r*colNum)+uint(j)+2].AddCell()
                if trow.Subject[A] == trow.Subject[B] {
                    cell.Value = trow.Subject[A]
                } else {
                    if trow.Subject[A] != "" {
                        cell.Value = trow.Subject[A] + " (A)"
                    }
                    if trow.Subject[B] != "" {
                        cell.Value += "\n" + trow.Subject[B] + " (B)"
                    }
                }

                cell = row[(r*colNum)+uint(j)+2].AddCell()
                if trow.Teacher[A] == trow.Teacher[B] {
                    cell.Value = trow.Teacher[A]
                } else {
                    if trow.Teacher[A] != "" {
                        cell.Value = trow.Teacher[A]
                    }
                    if trow.Teacher[B] != "" {
                        cell.Value += "\n" + trow.Teacher[B]
                    }
                }

                sheet.SetColWidth(colWidthForCabinets(int(j)))
                cell = row[(r*colNum)+uint(j)+2].AddCell()
                if trow.Cabinet[A] == trow.Cabinet[B] {
                    cell.Value = trow.Cabinet[A]
                } else {
                    if trow.Cabinet[A] != "" {
                        cell.Value = trow.Cabinet[A]
                    }
                    if trow.Cabinet[B] != "" {
                        cell.Value += "\n" + trow.Cabinet[B]
                    }
                }
            }

            index++
        }
    }

    err = file.Save(filename)
    if err != nil {
        return err
    }

    return nil
}
