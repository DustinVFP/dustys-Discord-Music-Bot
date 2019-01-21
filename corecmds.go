package main

import (
<<<<<<< HEAD
	"./dlogger"
	"fmt"
	"github.com/andersfylling/disgord"
	"time"
=======
	"fmt"
	"github.com/andersfylling/disgord"
	"./dlogger"
	"strings"
>>>>>>> master
)

func init() {
<<<<<<< HEAD

	// stest command
	cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var error error
			output := "test for possible better command handler and such\n"
			data.Message.RespondString(session, output)
			dlogger.LogOld(0, 0, "Woop", output)
			return error
		},
		cmdName:     "SpecialTest",
		cmdCalls:    []string{"stest", "st"},
		cmdMinDesc:  "Special Test for testing new command handler prototype",
		cmdFullDesc: "Special Test for testing new command handler prototype",
		cmdFirstChr: "s",
		cmdModule:   "core",
	}))

	// Help command
	cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error

			var basichelp = "``` -| Core commands |- \n - Help: The help command, \n - Version: Displays the version running and some other info \n - Ping: Pong!, \n -| Text/Test commands |- \n - Hello: Says hello back```"

			var output string

			aln := len(args)
=======
	corecmdslist[0] = "help"
	corecmdslist[1] = "ping"
	corecmdslist[2] = "version"
	corecmdslist[3] = "hello"
	corecmdslist[4] = "test"
	corecmdslist[5] = "debug"
	corecmdslist[6] = "dab"
	corecmdslist[7] = "love"

}

type helpstruct struct {
	dhname        string
	dhalts        string
	dhdescription string
	helshort      string
	cmdtitle      string
}

func cmdcorehandler(message, args string, session disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	//var responce string = "debug check ignore this"
	var err error

	switch message {
	case "help":
		err = cmdcorehelp(args, session, data)
	case "ping":
		err = cmdcoreping(args, session, data)
	case "version":
		err = cmdcoreversion(args, session, data)
	case "hello":
		err = cmdcorehello(args, session, data)
	case "test":
		err = cmdcorehello(args, session, data)
	case "debug":
		err = cmdcoredebug(args, session, data)
	case "dab":
		err = cmdcoredab(args, session, data)
	case "love":
		err = cmdcorelove(args, session, data)
	}
	if err != nil {
		dlogger.LogOld(30, 35, "responce error", err.Error())
		msg.RespondString(session, "Something seems to have went wrong")
		msg.RespondString(session, err.Error())
	}
	//msg.RespondString(session, responce)
}

func cmdcorehelp(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	//detailedconfcmd := [512]string{"Help"}
	//detailedconfdsc := [512]string{"The help command, it outputs helpy stuff"}
	var err error

	var basichelp string = "``` -| Core commands |- \n - Help: The help command, \n - Version: Displays the version running and some other info \n - Ping: Pong!, \n -| Text/Test commands |- \n - Hello: Says hello back```"

	var output string

	if cmd == "" {
		output = basichelp
	} else {
		output = "to be written"
		// to be written
	}

	data.Message.RespondString(session, output)
	return err
}

func cmdcorehello(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	content := "Hello "
	output := fmt.Sprintf("%s%s%s%s", content, "<@", data.Message.Author.ID, ">")
	data.Message.RespondString(session, output)
	return err
}

func cmdcorelove(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	var content string
	content = "Love You "
	var uid string
	if cmd == "" {
		uid = fmt.Sprint("<@", data.Message.Author.ID, ">")
	} else {
		if strings.Contains(cmd, "me") {
			uid = fmt.Sprint("<@", data.Message.Author.ID, ">")
		} else {
			uid = cmd
		}
	}
	output := fmt.Sprintf("%s%s", content, uid)
	data.Message.RespondString(session, output)
	return err
}
>>>>>>> master

			dlogger.LogOld(0, 5, "aln", fmt.Sprint(aln))

<<<<<<< HEAD
			if aln <= 1 {
				output = basichelp
			} else {
				output = "to be written"
				// to be written
			}
=======
func cmdcoreversion(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	user, err := session.GetCurrentUser()
	output := fmt.Sprint("Running: ", appname, " ", version, "\nlocally configured as: ", conf_Name, "\nRunning under user: ", user.Username, "#", user.Discriminator)
	data.Message.RespondString(session, output)
	return err
}
>>>>>>> master

			data.Message.RespondString(session, output)
			return err
		},
		cmdName:     "Help",
		cmdCalls:    []string{"help", "cmds"},
		cmdMinDesc:  "help, displays help message with basic commands",
		cmdFullDesc: "Help command, Displays and shows a list of basic commands as well as some other stuff",
		cmdFirstChr: "h",
		cmdModule:   "core",
	}))

	cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error

			var msgtime = data.Message.Timestamp
			var currenttime = time.Now()

			difference := currenttime.Sub(msgtime)

			output := fmt.Sprintf(
				"Pong!\n message sent at %s processed at %s \n Difference: %v",
				msgtime.Format("3:04:05.000 PM"), currenttime.UTC().Format("3:04:05.000 PM"), difference.Seconds(),
			)

			data.Message.RespondString(session, output)
			return err
		},

		cmdName:     "Ping",											//full command name (currently unused in code)
		cmdCalls:    []string{"ping"},									//the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
		cmdMinDesc:  "Pings and returns estimated processing time",		//smol description
		cmdFullDesc: "Pings and returns estimated processing time", 	//big description
		cmdFirstChr: "p",												//first letter of the command
		cmdModule:   "core",											//the module which the command exists in
	}))

	cmdarray = append(cmdarray, (cmddata{ // this adds a command entry what follows below is all the info about the command

		// this is the function which basically contains all the logic and such to make a command work
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error
			user, err := session.GetCurrentUser()
			output := fmt.Sprint("Running: ", appname, " ", version, "\nlocally configured as: ", confName, "\nRunning under user: ", user.Username, "#", user.Discriminator)
			data.Message.RespondString(session, output)
			return err
		},
		cmdName:     "Version",												//full command name (currently unused in code)
		cmdCalls:    []string{"version"},										//the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
		cmdMinDesc:  "Returns bot version",											//smol description
		cmdFullDesc: "Returns bot version and some other various info",	//big description
		cmdFirstChr: "v",															//first letter of the command
		cmdModule:   "core",														//the module which the command exists in
	}))

	cmdarray = append(cmdarray, (cmddata{ // this adds a command entry what follows below is all the info about the command

		// this is the function which basically contains all the logic and such to make a command work
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error
			content := "Hello "
			output := fmt.Sprintf("%s%s%s%s", content, "<@", data.Message.Author.ID, ">")
			data.Message.RespondString(session, output)
			return err
		},

<<<<<<< HEAD
		cmdName:     "Hello",												//full command name (currently unused in code)
		cmdCalls:    []string{"hello","hi"},										//the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
		cmdMinDesc:  "Hello!",											//smol description
		cmdFullDesc: "Says hello back to you",	//big description
		cmdFirstChr: "h",															//first letter of the command
		cmdModule:   "core",														//the module which the command exists in
	}))
=======
func cmdcoredebug(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	output := "u!test"
	data.Message.RespondString(session, output)
	return err
>>>>>>> master
}
