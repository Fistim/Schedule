package api

import(
	"net/http"
	"encoding/json"
	"os"
	"mime/multipart"
	"crypto/rand"
	xlsx "../schedule"
)

func ImportPlanRoute(w http.ResponseWriter, r *http.Request){
	input, _, err := r.FormFile("uploadfile")
	if err != nil {
		json.NewEncoder(w).Encode(struct{Error string}{Error: "Error read upload file"})
		return
	}
	tempname := readFile(input)
	result := xlsx.ImportXLSX(tempname)
	json.NewEncoder(w).Encode(result)
}

func readFile(input multipart.File) string {
	var (
		size   = uint64(0)
		buffer = make([]byte, 2 << 20) // 1MiB
	)

	tempname := GenerateRandomString(16) + ".xlsx"

	output, err := os.OpenFile(
		tempname,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if err != nil {
		printError(err)
		return "error"
	}

	for {
		length, err := input.Read(buffer)
		if err != nil {
			printError(err)
			break
		}
		size += uint64(length)
		output.Write(buffer[:length])
	}

	input.Close()
	output.Close()
	return tempname
}

func GenerateRandomString(max int) string {
    var slice []byte = make([]byte, max)
    _, err := rand.Read(slice)
    if err != nil {
    	printError(err)
        return ""
    }

    return string(slice)
}