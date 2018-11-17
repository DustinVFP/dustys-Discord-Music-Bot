package main

// bot link https://discordapp.com/api/oauth2/authorize?client_id=511179733384429578&scope=bot&permissions=108514368

import (
	"fmt"
	//"flag"
	//"os"
	//"os/signal"
	//"syscall"
	"io/ioutil"
	"encoding/json"
	"strings"
	"time"
	//"reflect"
	
	"github.com/andersfylling/disgord"
)

type Lowconf struct {
	Token		string
	DBGLvl		string
	Prefix		string
	Prefix2		string
	Prefix3		string
	Prefix4		string
}

var (
	CFG_Token	string
	CFG_Debug	int
	CFG_Prefix1	string
	CFG_Prefix2	string
	CFG_Prefix3	string
	CFG_Prefix4	string
)

type error interface {
    Error() string
}

const version = "0.0.0.0:alpha"

func LogReporter(importance, dbglevel int, info, info2 string) {
	// Importance levels are, 0: info, 5: message, 10: log, 20: warning, 30: Error, 40: Alert, 50, Critical
	
	// Debug levels are	 5: Only report when set to all (unimportant), 15: report on high (regular debug), 
	// 					25: report on regular or higher (normal log messages), 35: report on minimal or higher (Important log messages)
	
	// Check debug level of report compared to debug level set in the settings to decide if a message is worth showing
	if dbglevel > CFG_Debug {
		// check report type
		var report string
		switch importance {
			case 0:
				report = "> Info :"
			case 1:
				report = "| Info :"
			case 5: 
				report = "> message :"
			case 10:
				report = "> Log :"
			case 11:
				report = "| Log :"
			case 20:
				report = "> Warning :"
			case 21:
				report = "| Warning :"
			case 30:
				report = "> Error :"
			case 31:
				report = "| Error :"
			case 40:
				report = "> Alert :"
			case 41:
				report = "| Alert :"
			case 50:
				report = "> Critical :"
			case 51:
				report = "| Critical :"
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
		CFG_Debug = 0
		_error = true
		LogReporter(50, 999999, "Failed to read core config", "")
		LogReporter(51, 999999, err.Error(), "")
		//fmt.Println(" > Critical Error -> Failed to read core config \n > Critical Error -> Error: ", err)
	}
	var data Lowconf
	
	err = json.Unmarshal([]byte(file),&data)
	if err != nil {
		CFG_Debug = 0
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
	
	CFG_Token = data.Token
	CFG_Debug = dbg
	CFG_Prefix1 = data.Prefix
	CFG_Prefix2 = data.Prefix2
	CFG_Prefix3 = data.Prefix3
	CFG_Prefix4 = data.Prefix4
	
	if !_error {
		LogReporter(0, 25, "Loaded Config", "")
	} else {
		CFG_Debug = 0
		LogReporter(50, 999999, "Core Config Failed to load", "")
		panic(err)
	}
}

func func_init() {
	LogReporter(0,99,"Starting up DustysDMB version", version)
	setupConf();
	LogReporter(0,5,"Prefix is", CFG_Prefix1)
	//fmt.Println(Token)
}

func messageDo(msg string) (string, string, error) {
	var responce, meta string
	var err error
	fmt.Println(CFG_Prefix1, "", strings.HasPrefix(msg, CFG_Prefix1))
	if (strings.HasPrefix(msg, CFG_Prefix1)) == true {
		responce = "hello"
	}
	
	return responce, meta, err
}

func main() {
	func_init()
	
	session, err := disgord.NewSession(&disgord.Config{
		Token: CFG_Token,
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
		LogReporter(5, 15, "Message recived with content", msg.Content)
		
		responce, meta, err := messageDo(msg.Content)
		print(meta)
		if err != nil {
			LogReporter(30, 45, "message error", err.Error())
			msg.RespondString(session, "Error, Something went wrong")
		} else {
			msg.RespondString(session, responce)
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



/* func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error opening session: ", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot active: Press CTRL-C to exit.")
	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
*/
