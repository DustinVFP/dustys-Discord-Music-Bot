package main

import(
	//"fmt"
	//"time"
	"io/ioutil"
	"encoding/json"
<<<<<<< HEAD
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
=======
	"gitea.pi.lan/dvf-productions/dlogger"
)

// depricated as it has been seperated off into its own package to be developed seperately
/*
func LogReporter(importance, dbglevel int, info, info2 string) {
	// Importance levels are, 0: info, 5: message, 10: log, 20: warning, 30: Error, 40: Alert, 50, Critical
	
	// Debug levels are	 5: Only report when set to all (unimportant), 15: report on high (regular debug), 
	// 					25: report on regular or higher (normal log messages), 35: report on minimal or higher (Important log messages)
	
	// Check debug level of report compared to debug level set in the settings to decide if a message is worth showing
	if dbglevel > conf_Debug {
		// check report type
		var report string
		switch importance {
			case 0:
				report = "> Info >"
			case 1:
				report = "| Info |"
			case 5: 
				report = "> message >"
			case 10:
				report = "> Log >"
			case 11:
				report = "| Log >"
			case 20:
				report = "> Warning >"
			case 21:
				report = "| Warning >"
			case 30:
				report = "> Error >"
			case 31:
				report = "| Error >"
			case 40:
				report = "> Alert >"
			case 41:
				report = "| Alert >"
			case 50:
				report = "> Critical >"
			case 51:
				report = "| Critical >"
		}
		// check if message has 2 values or 3 and then display as appropreate
		if info2 != "" {
			fmt.Println(time.Now().Format("2006-01-02 03:04.05PM"), "", report, info, ":", info2)
		} else {
			fmt.Println(time.Now().Format("2006-01-02 03:04.05PM"), "", report, info)
		}
	}
>>>>>>> ea46c40178f42a40a7332ab2ca2005d0bc1be1fa
}
*/

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
