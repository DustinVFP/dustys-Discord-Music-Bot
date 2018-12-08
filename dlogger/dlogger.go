package dlogger

import(
	"fmt"
	"time"
)

var loglevel int = 25
var txtorcmd [2]int

func SetLevels(lvl int) {
	loglevel = lvl
	//txtorcmd = toc
}

func Check() int {
	return loglevel
}

func LogOld(importance, dbglevel int, info, info2 string) {
	// Importance levels are, 0: info, 5: message, 10: log, 20: warning, 30: Error, 40: Alert, 50, Critical
	
	// Debug levels are	 5: Only report when set to all (unimportant), 15: report on high (regular debug), 
	// 					25: report on regular or higher (normal log messages), 35: report on minimal or higher (Important log messages)
	
	// Check debug level of report compared to debug level set in the settings to decide if a message is worth showing
	if dbglevel > loglevel {
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

func LogInfo(dbglevel int, info, info2 string) {
	// Debug levels are	 5: Only report when set to all (unimportant), 15: report on high (regular debug), 
	// 					25: report on regular or higher (normal log messages), 35: report on minimal or higher (Important log messages)
	
	// Check debug level of report compared to debug level set in the settings to decide if a message is worth showing
	if dbglevel > loglevel {
		// check if message has 2 values or 3 and then display as appropreate
//		if txtorcmd[0] == 1 {
			if info2 != "" {
				fmt.Println(time.Now().Format("2006-01-02 03:04.05PM"), "> Info >", info, ":", info2)
			} else {
				fmt.Println(time.Now().Format("2006-01-02 03:04.05PM"), "> Info >", info)
//			}
		}
	}
}

func LogExtraInfo(dbglevel int, info, info2 string) {
	// Debug levels are	 5: Only report when set to all (unimportant), 15: report on high (regular debug), 
	// 					25: report on regular or higher (normal log messages), 35: report on minimal or higher (Important log messages)
	
	// Check debug level of report compared to debug level set in the settings to decide if a message is worth showing
	if dbglevel > loglevel {
		// check if message has 2 values or 3 and then display as appropreate
		if info2 != "" {
			fmt.Println(time.Now().Format("2006-01-02 03:04.05PM"), "       |", info, ":", info2)
		} else {
			fmt.Println(time.Now().Format("2006-01-02 03:04.05PM"), "       |", info)
		}
	}
}
