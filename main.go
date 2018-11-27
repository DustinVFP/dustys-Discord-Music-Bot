package main

// bot link https://discordapp.com/api/oauth2/authorize?client_id=516941386369597441&scope=bot&permissions=518208

import (
	"fmt"
	"flag"
	//"os"
	//"os/signal"
	//"syscall"
	"io/ioutil"
	"encoding/json"
	"strings"
	"time"
	"strconv"
	//"reflect"
	
	"github.com/andersfylling/disgord"
)

type Lowconf struct {
	Token		string
	DBGLvl		string
	Prefix		[]string
	Name		string
}

var (
	conf_Name	string
	conf_Token	string
	conf_Debug	int
	conf_Prefix	[]string
)

type error interface {
    Error() string
}

const version = "0.0.0.0:alpha"

var UseTUI bool
var chk1 int = 0
var messagechk1 string = "~~~~~~"
var messagechk2 string = "~~~~~~~"

func init() {
	flag.BoolVar(&UseTUI, "tui", false, "Use Tui, true/false")
	flag.Parse()
	LogReporter(0,5,"tui flag set to", strconv.FormatBool(UseTUI))
	
	setupConf()
	
	LogReporter(0,99,"Starting up", conf_Name)
	LogReporter(1,99,"Version", version)
	setupConf();
	LogReporter(0,5,"Prefix is", conf_Prefix[0])
}

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
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		conf_Debug = 0
		_error = true
		LogReporter(50, 999999, "Failed to read core config", "")
		LogReporter(51, 999999, err.Error(), "")
		//fmt.Println(" > Critical Error -> Failed to read core config \n > Critical Error -> Error: ", err)
	}
	var data Lowconf
	
	err = json.Unmarshal([]byte(file),&data)
	if err != nil {
		conf_Debug = 0
		_error = true
		LogReporter(50, 999999, "Failed to read core config", "")
		LogReporter(51, 999999, err.Error(), "")
	}
	
	var dbg int
	switch data.DBGLvl {
		case "ALL":
			dbg = 0
		case "HIGH": 
			dbg = 10
		case "REGULAR":
			dbg = 20
		case "MINIMAL":
			dbg = 30
	}
	
	conf_Token = data.Token
	conf_Debug = dbg
	conf_Prefix = data.Prefix
	conf_Name = data.Name
	
	if !_error {
		LogReporter(0, 25, "Loaded Config", "")
	} else {
		conf_Debug = 0
		LogReporter(50, 999999, "Core Config Failed to load", "")
		panic(err)
	}
}

func func_init() {

	//fmt.Println(Token)
}

func prefixCheck(data string) bool {
	prearraylen := len(conf_Prefix)
	LogReporter(0,5,"Prefix Amount", strconv.Itoa(prearraylen))
	for i := 0; i<prearraylen; i++ {
		pfx := conf_Prefix[i]
		LogReporter(0,5,"Prefix", pfx)
		if strings.HasPrefix(data, pfx) {
			return true
			break
		}
	}
	return false
}

func messageDo(message string, session disgord.Session, data *disgord.MessageCreate) /*(string, string, error)*/ {
	var responce/*, meta*/ string
	//var err error
	
	msg := data.Message
	
	messagechk1 = msg.Content
	
	ckprfx := prefixCheck(message)
	
	if ckprfx {
		responce = "hello"
		msg.RespondString(session, responce)
	}
	
	messagechk1 = "~~~~~~"
	//return responce, meta, err
}

func main() {

	session, err := disgord.NewSession(&disgord.Config{
		Token: conf_Token,
	})
	if err != nil {
		LogReporter(50, 999999, "Failed to open discord session", "")
		LogReporter(51, 999999, err.Error(), "")
		panic(err)
	}
	
	myself, err := session.GetCurrentUser()
	if err != nil {
		LogReporter(50, 999999, "Discord Session error", "")
		LogReporter(51, 999999, err.Error(), "")
		panic(err)
	}
	
	session.On(disgord.EventMessageCreate, func(session disgord.Session, data *disgord.MessageCreate) {
		msg := data.Message
		LogReporter(5, 15, "Message recived", msg.Content)
		
		messagechk2 = msg.Content
		
		if (messagechk1 != messagechk2) {
			go messageDo(msg.Content, session, data)
			//print(meta)
			//if err != nil {
			//	LogReporter(30, 45, "message error", err.Error())
			//	msg.RespondString(session, "Error, Something went wrong")
			//} else {
			//	msg.RespondString(session, responce)
			//}
			}
	})
	
	err = session.Connect()
	if err != nil {
		LogReporter(50, 999999, "Discord Session error", "")
		LogReporter(51, 999999, err.Error(), "")
		panic(err)
	}
	
	//fmt.Printf("Hello, %s!\n", myself.String())
	LogReporter(0,15, "Running under user", myself.String())
	
	session.DisconnectOnInterrupt()	
}
