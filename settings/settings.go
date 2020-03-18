package settings

import (
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB
const (
	DOMAIN 	   = "http://schedule.tomtit.tomsk.ru"
	TimeLayout = "15:04"
	dateLayout = ""
	login      = "dbuser"
	password   = "QWEasd123"
	ipaddr     = "192.168.10.14"
	port       = "3306"
	dbname     = "Schedule"
	protocol   = "tcp"
	args       = "parseTime=true&charset=utf8&loc=Local"
	CONSTR     = login + ":" + password + "@" + protocol + "(" + ipaddr + ":" + port + ")/" + dbname + "?" + args
)

func WriteFile(filename, data string) error {
 	file, err := os.Create(filename)
 	if err != nil {
 		return err 
 	}
 	file.WriteString(data)
 	file.Close()
 	return nil
}

