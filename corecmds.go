package main

import (
	"gitea.pi.lan/DVF-Productions/DustysDBMB/dlogger"
	"fmt"
	"github.com/andersfylling/disgord"
	"time"
	"math"
	"strings"
)

var helpcache string

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

			var output string

			if helpcache == "" {
				arraylen := len(cmdarray)

				var helpsarray = make(map[string][]string)
				var helpsmods= make([]string, 0)

				var mod string

				for i := 0; i < arraylen; i++ {
					if cmdarray[i].cmdShowHelp || !cmdarray[i].cmdShowHelp {
						mod = cmdarray[i].cmdModule
						helpsarray[mod] = append(helpsarray[mod], (fmt.Sprint(cmdarray[i].cmdName, " : ", cmdarray[i].cmdMinDesc)))
					}
				}
				fmt.Println(helpsarray)

				for k := range helpsarray {
					helpsmods = append(helpsmods, fmt.Sprint("---| " , k , " |---\n```"))
					helpsmods = append(helpsmods, strings.Join(helpsarray[k], "\n"))
					helpsmods = append(helpsmods, "```")
				}

				fmt.Println(helpsmods)
				helpcache = strings.Join(helpsmods, "\n")
			}

			aln := len(args)

			dlogger.LogOld(0, 5, "aln", fmt.Sprint(aln))

			if aln <= 1 {
				output = "to be written hand written help menu, append 'all' or 'cmds' to the end of the command to view auto generated command list \n or use the [to be programmed] cmds command"
			} else if args[1] == "all" || args[1] == "cmds" {
				output = helpcache
			} else {
//				if cmdc == cmdarray[i].cmdCalls[1] | cmdc == cmdarray[i].cmdName {
//
//					}
				// to be written
				output = "to be codded"
			}

			data.Message.RespondString(session, output)
			return err
		},
		cmdName:     "NewHelp",
		cmdCalls:    []string{"newhelp", "nh"},
		cmdMinDesc:  "newhelp, displays help message with basic commands",
		cmdFullDesc: "newHelp command, Displays and shows a list of basic commands as well as some other stuff",
		cmdFirstChr: "n",
		cmdModule:   "core",
	}))

	// ping command
	cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error

			var msgtime = data.Message.Timestamp
			var currenttime = time.Now()

			difference := currenttime.Sub(msgtime.Time)

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
			output := fmt.Sprint(	"Running: ", appname, " ", version, "\n",
									"locally configured as: ", confName, "\n",
									"Running under user: ", userSelf.Username, "#", userSelf.Discriminator, "\n\n",
									"Started on: ", starttime.Format("2006 Jan 2 at 3:04pm"),
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

		cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
			var err error
			var currenttime = time.Now()
			var difference = currenttime.Sub(starttime)
			var timest string
			tstmpday := math.Floor((difference.Hours()/24))
			tstmphrs := (math.Floor(difference.Hours()) - (tstmpday * 24))
			tstmpmin := (math.Floor(difference.Minutes() - (tstmphrs * 60)))
			tstmpsec := (math.Floor(difference.Seconds() - (tstmpmin * 60)))
			if difference.Hours() > 24 {
				timest = fmt.Sprintf("%.0f Days, %.0f Hours, %.0f Minutes", tstmpday, tstmphrs, tstmpmin)
			} else if difference.Hours() > 1 {
				timest = fmt.Sprintf("%.0f Hours, %.0f Minutes, %.0f Seconds", tstmphrs, tstmpmin, tstmpsec)
			} else if difference.Hours() < 1 {
				timest = fmt.Sprintf("%.0f Minutes, %.0f Seconds", tstmpmin, tstmpsec)
			}
			content := fmt.Sprint(	"Has Ran for: ", timest, "\n",
									"Started at: ", starttime.Format("2006 Jan-2 > 3:04:05PM"), "\n",
									"Current time: ", currenttime.Format("2006 Jan-2 > 3:04:05PM"), "\n",
									)
			output := fmt.Sprintf("%s%s%s%s", content, "<@", data.Message.Author.ID, ">")
			data.Message.RespondString(session, output)
			return err
		},

		cmdName:     "Uptime",
		cmdCalls:    []string{"uptime"},
		cmdMinDesc:  "How long the bot has been online for",
		cmdFullDesc: "Shows how long the bot has been online for aswell as the date it was started",
		cmdFirstChr: "u",
		cmdModule:   "core",
	}))
}
