package punchclock

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Start() {

	db_host := ""

	// Connect to the database
	db, err := gorm.Open("mysql", fmt.Sprintf("user:password@%s/dbname?charset=utf8&parseTime=True&loc=Local", db_host)
	if err != nil {
		panic("Cannot connect to database." + err.Error())
	}

	_=db
}