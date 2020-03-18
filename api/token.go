package api

import(
	"encoding/json"
	"encoding/base64"
	"crypto/rand"
	"io/ioutil"
	"strings"
	"net/http"
	"fmt"
)

type token struct{
	Name string
	Token string
}

var tokens []token

func generateRandomBytes(max int) []byte {
    var slice []byte = make([]byte, max)
    _, err := rand.Read(slice)
    if err != nil {
        return nil
    }
    return slice
}

func base64Encode(data []byte) string {
    return base64.StdEncoding.EncodeToString(data)
}

func GenerateNewToken(name string) string{
	generated:=base64Encode(generateRandomBytes(10))
	var newToken token = token{name, generated}
	tokens = append(tokens, newToken)

	writeTokensToFile()

	return generated
}

func writeTokensToFile(){

	tokens=getTokens()
	jsonTokens, err:=json.Marshal(tokens)

	if err!=nil{
		printError(err)
	}

	err=ioutil.WriteFile("tokens.json", jsonTokens, 0644)

	if err!=nil{
		printError(err)
	}
	return
}

func getTokens() []token{
	data, err := ioutil.ReadFile("tokens.json")
	fmt.Println(string(data))
	if err != nil{
		printError(err)
		return nil
	}
	var tokens []token
	err = json.Unmarshal(data, &tokens)
	if err!=nil{
		printError(err)
		return nil
	}
	return tokens
}

func CheckToken(r *http.Request) bool{
	newToken := r.Header.Get("Authorization")
	newToken = strings.Replace(newToken, "Bearer ", "", 1)

	tokens=getTokens()
	fmt.Println(tokens)
	for _, check := range tokens{
		if newToken==check.Token{
			return true
		}
	}

	return false
}

