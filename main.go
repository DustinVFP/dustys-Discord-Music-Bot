package main

// bot link https://discordapp.com/api/oauth2/authorize?client_id=539992577517027342&scope=bot&permissions=518208

import (
	"flag"
	"fmt"
	//"os"
	//"os/signal"
	//"syscall"
	//"io/ioutil"
	//"encoding/json"
	"strings"
	"time"
	"strconv"
	//"reflect"
	"gitea.pi.lan/DVF-Productions/DustysDBMB/dlogger"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"

	"github.com/andersfylling/disgord"
)

var (
	confName   string
	confToken  string
	confDebug  int
	confPrefix []string

	confdbType  string
	confdbAddress string
	confdbUser	string
	confdbPass string
	confdbName	string
)

type error interface {
	Error() string
}

const version = "v0.0.1.1:alpha"
const appname = "Dustys Wip Discord Bot"

var starttime = time.Now()

var useTUI bool
var chk1 int
var messagechk string

type cmddata struct {
	cmdFunc     func(args []string, s disgord.Session, d *disgord.MessageCreate) error
	cmdName     string
	cmdCalls    []string
	cmdMinDesc  string
	cmdFullDesc string
	cmdFirstChr string
	cmdModule   string
	cmdShowHelp bool
}

var userSelf *disgord.User

var cmdarray = make([]cmddata, 0)


// do some checks and init some stuff
func init() {
	flag.BoolVar(&useTUI, "tui", false, "Use Tui, true/false")
	flag.Parse()
	confDebug = 5
	dlogger.LogOld(0, 15, "tui flag set to", strconv.FormatBool(useTUI))

	setupConf()

	dlogger.LogOld(0, 99, "Starting up", confName)
	dlogger.LogOld(1, 99, "Version", version)
	setupConf()
	dlogger.LogOld(0, 15, "Prefix is", confPrefix[0])
}

func prefixCheck(data string) (bool, string) {

	arraylen := len(confPrefix)
	dlogger.LogOld(0, 5, "Prefix Amount", strconv.Itoa(arraylen))
	dlogger.LogOld(0, 5, "Prefix", fmt.Sprint(confPrefix))

	for i := 0; i < arraylen; i++ {
		if strings.HasPrefix(data, confPrefix[i]) {
			return true, confPrefix[i]
			break
		}
	}
	return false, ""
}

func messageDo(session disgord.Session, data *disgord.MessageCreate) /*(string, string, error)*/ {
	//var responce/*, meta*/ string
	//var err error

	messagechk = data.Message.ID.String()

	ckprfx, prefix := prefixCheck(data.Message.Content)

	if ckprfx {
		msg := strings.Replace(data.Message.Content, prefix, "", -1)
		arg := strings.Fields(msg)
		cmd := strings.ToLower(arg[0])

		// some debug code
		if confDebug < 5 {
			dlogger.LogOld(0, 5, "cmd", cmd)
			dlogger.LogOld(0, 5, "args", fmt.Sprint(arg))

			arraylen3 := len(arg)
			dlogger.LogOld(0, 5, "Argamt", strconv.Itoa(arraylen3))

			for i := 0; i < arraylen3; i++ {
				dlogger.LogOld(0, 5, "Arg", arg[i])
			}
		}

		dlogger.LogOld(0, 5, "cmd data", data.Message.Content)

		arraylen := len(cmdarray)
		dlogger.LogOld(0, 5, "cmds count", strconv.Itoa(arraylen))

		for i := 0; i < arraylen; i++ {
			if strings.HasPrefix(cmd, cmdarray[i].cmdFirstChr) {

				arraylen2 := len(cmdarray[i].cmdCalls)

				dlogger.LogOld(0, 5, "call count", strconv.Itoa(arraylen2))

				for i2 := 0; i2 < arraylen2; i2++ {
					cmdc := cmdarray[i].cmdCalls[i2]
					dlogger.LogOld(0, 5, "cmdc", cmdc)
					if cmdc == "" {
						break
					}
					if cmd == cmdc {
						dlogger.LogOld(0, 5, "cmd data", cmdc)
						dlogger.LogOld(0, 5, cmdc, cmd)
						err := cmdarray[i].cmdFunc(arg, session, data)
						if err != nil {
							dlogger.LogOld(30, 35, "responce error", err.Error())
							data.Message.RespondString(session, "Something seems to have went wrong")
							data.Message.RespondString(session, err.Error())
						}
					}
				}
			} else {
				if cmdarray[i].cmdFirstChr == "" {
					break
				}
			}
			//dlogger.LogOld(0,5,"cmdchk", cmd)
			//if strings.HasPrefix(msg, cmd) {
			//	dta := strings.Replace(msg, cmd, "", -1)
			//	dlogger.LogOld(0,5,"command", cmd)
			//	dlogger.LogOld(0,5,"arguments", dta)
			//	go cmdcorehandler(cmd, dta , session, data)
			//	break
			//}
		}

		//responce = "hello"
		//msg.RespondString(session, responce)
	}

	//return responce, meta, err
}

func main() {

	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", confdbUser, confdbPass, confdbAddress, confdbName))

	//if err != nil {
	//	panic(err.Error())
	//}

	//defer db.Close()

	session, err := disgord.NewClient(&disgord.Config{
		BotToken: confToken,
		Logger: disgord.DefaultLogger(true),
	})
	if err != nil {
		dlogger.LogOld(50, 999999, "Failed to open discord session", "")
		dlogger.LogOld(51, 999999, err.Error(), "")
		panic(err)
	}

	myself, err := session.GetCurrentUser()
	if err != nil {
		dlogger.LogOld(50, 999999, "Discord Session error", "")
		dlogger.LogOld(51, 999999, err.Error(), "")
		panic(err)
	}

	userSelf, err = session.GetCurrentUser()
	if err != nil {
		dlogger.LogOld(30, 25, "Error getting current user", "")
		return
	}
	fmt.Println(userSelf.ID)

	session.On(disgord.EvtMessageCreate, func(session disgord.Session, data *disgord.MessageCreate) {
		msg := data.Message
		dlogger.LogOld(5, 15, "Message recived", msg.Content)

		fmt.Println(data.Message.Author)

		if userSelf != nil && data.Message.Author.ID != userSelf.ID {
			if msg.ID.String() != messagechk {
				go messageDo(session, data)
			}
		}
	})

	err = session.Connect()
	if err != nil {
		dlogger.LogOld(50, 999999, "Discord Session error", "")
		dlogger.LogOld(51, 999999, err.Error(), "")
		panic(err)
	}

	dlogger.SetLevels(confDebug)
	tst := dlogger.Check()
	dlogger.LogOld(0, 15, "debug check", strconv.Itoa(tst))
	dlogger.LogExtraInfo(15, "test", "")

	dlogger.LogOld(0, 15, "Running under user", myself.String())

	session.DisconnectOnInterrupt()
}
