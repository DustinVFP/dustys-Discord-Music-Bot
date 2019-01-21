package main

import (
	"./dlogger"
	"fmt"
	"github.com/andersfylling/disgord"
	"time"
	//"math"
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
	
	// ping command
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

		cmdName:     "Ping",
		cmdCalls:    []string{"ping"},
		cmdMinDesc:  "Pings and returns estimated processing time",
		cmdFullDesc: "Pings and returns estimated processing time",
		cmdFirstChr: "p",
		cmdModule:   "core",
	}))

	// version command
	cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error
			user, err := session.GetCurrentUser().Execute()
			if err == nil {
				dlogger.LogOld(30, 25, err.Error(), "")
			}
			output := fmt.Sprint(	"Running: ", appname, " ", version, "\n",
									"locally configured as: ", confName, "\n",
									"Running under user: ", user.Username, "#", user.Discriminator, "\n\n",
									"Started at: ", starttime.Format("3:04:05.000 PM"),
									)
			data.Message.RespondString(session, output)
			return err
		},
		cmdName:     "Version",
		cmdCalls:    []string{"version"},
		cmdMinDesc:  "Returns bot version",
		cmdFullDesc: "Returns bot version and some other various info",
		cmdFirstChr: "v",
		cmdModule:   "core",
	}))

	// hello command
	cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error
			content := "Hello "
			output := fmt.Sprintf("%s%s%s%s", content, "<@", data.Message.Author.ID, ">")
			data.Message.RespondString(session, output)
			return err
		},

		cmdName:     "Hello",
		cmdCalls:    []string{"hello","hi"},
		cmdMinDesc:  "Hello!",
		cmdFullDesc: "Says hello back to you",
		cmdFirstChr: "h",
		cmdModule:   "core",
	}))
	
	// in development shenanigans (don't mind me i'm just being stupid and not knowing how golang time works and trying to figure it out the hard way)
	
		//cmdarray = append(cmdarray, (cmddata{
		//cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			//var err error
			//var currenttime = time.Now()
			//var difference = currenttime.Sub(starttime)
			//if difference.Hours() > 24.0 {
				//timeout := fmt.Sprint(math.Round(difference.Hours / 24.0))
			//}
			//content := fmt.Sprint(	"Has Ran for: ", timeout, "/n",
									//"Started at: ", starttime.Format("2006 Jan 2 3:04:05.000 PM"), "/n",
									//"Current time: ", currenttime.Format("2006 Jan 2 3:04:05.000 PM"), "/n",
									//)
			//output := fmt.Sprintf("%s%s%s%s", content, "<@", data.Message.Author.ID, ">")
			//data.Message.RespondString(session, output)
			//return err
		//},

		//cmdName:     "Uptime",
		//cmdCalls:    []string{"uptime"},
		//cmdMinDesc:  "How long the bot has been online for",
		//cmdFullDesc: "Shows how long the bot has been online for aswell as the date it was started",
		//cmdFirstChr: "u",
		//cmdModule:   "core",
	//}))
}
