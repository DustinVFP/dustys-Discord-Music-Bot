package main

import(
	//"fmt"
	//"time"
	"io/ioutil"
	"encoding/json"
	"./dlogger"
)

type dbConf struct {
	DBType		string
	DBAddress	string
	DBPass		string
	DBUser		string
	DBName		string
}

type lowconf struct {
	Token	string
	DBGLvl	string
	Prefix	[]string
	Name	string
	Database dbConf
}

// shenanigans for loading in the configuration
func loadConf() {
	_error := false
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		confDebug = 5
		_error = true
		dlogger.LogOld(50, 999999, "Failed to read core config", "")
		dlogger.LogOld(51, 999999, err.Error(), "")
		//fmt.Println(" > Critical Error -> Failed to read core config \n > Critical Error -> Error: ", err)
	}
	var data lowconf

	err = json.Unmarshal([]byte(file),&data)
	if err != nil {
		confDebug = 5
		_error = true
		dlogger.LogOld(50, 999999, "Failed to read core config", "")
		dlogger.LogOld(51, 999999, err.Error(), "")
	}

	var dbg int
	switch data.DBGLvl {
		case "ALL":
			dbg = 0
		case "HIGH":
			dbg = 10
		case "NORMAL":
			dbg = 20
		case "REGULAR":
			dbg = 20
		case "DEFAULT":
			dbg = 20
		case "MINIMAL":
			dbg = 30
	}

	confToken = data.Token
	confDebug = dbg
	confPrefix = data.Prefix
	confName = data.Name

	confdbType = data.Database.DBType
	confdbAddress = data.Database.DBAddress
	confdbUser = data.Database.DBUser
	confdbPass = data.Database.DBPass
	confdbName = data.Database.DBName

	if !_error {
		dlogger.LogOld(0, 35, "Loaded Config", "")
	} else {
		confDebug = 0
		dlogger.LogOld(50, 999999, "Core Config Failed to load", "")
		panic(err)
	}
}
