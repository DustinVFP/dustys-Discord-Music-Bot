package main

import (
<<<<<<< HEAD
	"./dlogger"
=======
	"gitea.pi.lan/dvf-productions/dlogger"
	"github.com/andersfylling/disgord"
>>>>>>> ea46c40178f42a40a7332ab2ca2005d0bc1be1fa
	"fmt"
	"github.com/andersfylling/disgord"
	"time"
)

func init() {

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

			dlogger.LogOld(0, 5, "aln", fmt.Sprint(aln))

			if aln <= 1 {
				output = basichelp
			} else {
				output = "to be written"
				// to be written
			}

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

		cmdName:     "Hello",												//full command name (currently unused in code)
		cmdCalls:    []string{"hello","hi"},										//the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
		cmdMinDesc:  "Hello!",											//smol description
		cmdFullDesc: "Says hello back to you",	//big description
		cmdFirstChr: "h",															//first letter of the command
		cmdModule:   "core",														//the module which the command exists in
	}))
}
