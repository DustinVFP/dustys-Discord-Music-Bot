package main

import (
	"flag"
	"fmt"
	//"os"
	//"os/signal"
	//"syscall"
	//"io/ioutil"
	//"encoding/json"
	"strings"
	//"time"
	"strconv"
	//"reflect"
	"./dlogger"
<<<<<<< HEAD
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
=======
>>>>>>> master

	"github.com/andersfylling/disgord"
)

<<<<<<< HEAD
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

const version = "v0.0.1.0:alpha"
const appname = "Dustys Wip Discord Bot"

var useTUI bool
var chk1 int
var messagechk string
=======
type Lowconf struct {
	Token  string
	DBGLvl string
	Prefix []string
	Name   string
}

var (
	conf_Name   string
	conf_Token  string
	conf_Debug  int
	conf_Prefix []string
)

const version = "v0.0.0.1:alpha"
const appname = "Wip Discord Bot"

var UseTUI bool // not currently used
var chk1 int = 0

var messagechk1 string = "~~~~~~" // thing for a thing to prevent duplicate message respoces (may be uneeded, not sure yet)
var messagechk2 string = "~~~~~~~"
>>>>>>> master

type cmddata struct {
	cmdFunc     func(args []string, s disgord.Session, d *disgord.MessageCreate) error
	cmdName     string
	cmdCalls    []string
	cmdMinDesc  string
	cmdFullDesc string
	cmdFirstChr string
	cmdModule   string
}

var cmdarray = make([]cmddata, 0)


// do some checks and init some stuff
func init() {
	flag.BoolVar(&useTUI, "tui", false, "Use Tui, true/false")
	flag.Parse()
<<<<<<< HEAD
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
=======
	conf_Debug = 5
	dlogger.LogOld(0, 15, "tui flag set to", strconv.FormatBool(UseTUI))

	setupConf()

	dlogger.LogOld(0, 99, "Starting up", conf_Name)
	dlogger.LogOld(1, 99, "Version", version)
	setupConf()
	dlogger.LogOld(0, 15, "Prefix is", conf_Prefix[0])
}

func prefixCheck(data string) (bool, string) {
	prearraylen := len(conf_Prefix)
	dlogger.LogOld(0, 5, "Prefix Amount", strconv.Itoa(prearraylen))
	for i := 0; i < prearraylen; i++ {
		pfx := conf_Prefix[i]
		dlogger.LogOld(0, 5, "Prefix", pfx)
		if strings.HasPrefix(data, pfx) {
			return true, pfx
>>>>>>> master
			break
		}
	}
	return false, ""
}

<<<<<<< HEAD
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
=======
func messageDo(message string, session disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message

	messagechk1 = msg.Content

	ckprfx, prefix := prefixCheck(message)

	if ckprfx {
		message2 := strings.Replace(message, prefix, "", -1)
		dlogger.LogOld(0, 5, "cmd data", message2)

		prearraylen := len(corecmdslist)
		dlogger.LogOld(0, 5, "core cmds count", strconv.Itoa(prearraylen))
		for i := 0; i < prearraylen; i++ {
			cmd := corecmdslist[i]
			if cmd == "" {
				break
			}
			dlogger.LogOld(0, 5, "cmdchk", cmd)
			if strings.HasPrefix(message2, cmd) {
				dta := strings.Replace(message2, cmd, "", -1)
				dlogger.LogOld(0, 5, "command", cmd)
				dlogger.LogOld(0, 5, "arguments", dta)
				go cmdcorehandler(cmd, dta, session, data)
				break
>>>>>>> master
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

<<<<<<< HEAD
		//responce = "hello"
		//msg.RespondString(session, responce)
	}

	//return responce, meta, err
=======
	}

	messagechk1 = "~~~~~~"
>>>>>>> master
}

func main() {

	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", confdbUser, confdbPass, confdbAddress, confdbName))

	//if err != nil {
	//	panic(err.Error())
	//}

	//defer db.Close()

	session, err := disgord.NewSession(&disgord.Config{
<<<<<<< HEAD
		Token: confToken,
=======
		Token: conf_Token,
>>>>>>> master
		Debug: true,
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

	session.On(disgord.EventMessageCreate, func(session disgord.Session, data *disgord.MessageCreate) {
		msg := data.Message
		dlogger.LogOld(5, 15, "Message recived", msg.Content)

<<<<<<< HEAD
=======
		messagechk2 = msg.Content

>>>>>>> master
		user, err := session.GetCurrentUser()
		if err != nil {
			dlogger.LogOld(30, 25, "Error getting current user", "")
		}
		fmt.Println(user.ID)
		fmt.Println(data.Message.Author)
		if data.Message.Author.ID != user.ID {
<<<<<<< HEAD
			if msg.ID.String() != messagechk {
				go messageDo(session, data)
=======
			if messagechk1 != messagechk2 {
				go messageDo(msg.Content, session, data)
>>>>>>> master
			}
		}
	})

	err = session.Connect()
	if err != nil {
		dlogger.LogOld(50, 999999, "Discord Session error", "")
		dlogger.LogOld(51, 999999, err.Error(), "")
		panic(err)
	}

<<<<<<< HEAD
	dlogger.SetLevels(confDebug)
=======
	dlogger.SetLevels(conf_Debug)
>>>>>>> master
	tst := dlogger.Check()
	dlogger.LogOld(0, 15, "debug check", strconv.Itoa(tst))
	dlogger.LogExtraInfo(15, "test", "")

	dlogger.LogOld(0, 15, "Running under user", myself.String())

	session.DisconnectOnInterrupt()
}
