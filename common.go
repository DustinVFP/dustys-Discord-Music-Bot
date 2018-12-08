package main

import(
	"fmt"
	"time"
	"io/ioutil"
	"encoding/json"
	"go.pi.lan/dlogger"
)

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
}

func setupConf() {
	_error := false
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		conf_Debug = 5
		_error = true
		dlogger.LogOld(50, 999999, "Failed to read core config", "")
		dlogger.LogOld(51, 999999, err.Error(), "")
		//fmt.Println(" > Critical Error -> Failed to read core config \n > Critical Error -> Error: ", err)
	}
	var data Lowconf
	
	err = json.Unmarshal([]byte(file),&data)
	if err != nil {
		conf_Debug = 5
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
	
	conf_Token = data.Token
	conf_Debug = dbg
	conf_Prefix = data.Prefix
	conf_Name = data.Name
	
	if !_error {
		dlogger.LogOld(0, 35, "Loaded Config", "")
	} else {
		conf_Debug = 0
		dlogger.LogOld(50, 999999, "Core Config Failed to load", "")
		panic(err)
	}
}
