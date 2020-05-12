package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	_const "starter-golang-new/const"
	"strconv"
	"users-microservice/logs"
)

/**
this function to connect to database
*/
func ConnectToDatabase() {
	Db, err = gorm.Open("mysql", os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+
		os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+"?charset=utf8&parseTime=True&loc=Local&character_set_server=utf8")
	if err != nil {
		logs.StoreLog("users", "database", "ConnectToDatabase", "fail to connect database", err.Error())
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	if debug {
		_const.Services.DB.LogMode(debug)
	}
}
