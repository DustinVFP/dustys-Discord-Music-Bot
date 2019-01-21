package main

import "github.com/andersfylling/disgord"

import (
    "strings"
    "fmt"
)

// some memeish commands

func init() {

    //the gay command
	cmdarray = append(cmdarray, (cmddata{ // this adds a command entry what follows below is all the info about the command
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {

			var err error

			output := " GAYYYYYYYY \n :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: "
			data.Message.RespondString(session, output)

			return err
		},

		cmdName:     "Gay",
		cmdCalls:    []string{"gay", "gae"},
		cmdMinDesc:  "A very gay command",
		cmdFullDesc: "The Gayest of commands",
		cmdFirstChr: "g",
		cmdModule:   "topatoes",
	}))

    //this command loves you
	cmdarray = append(cmdarray, (cmddata{
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
            var err error
        	var uid string

            aln := len(args)

            if aln <= 1 {
                uid = fmt.Sprint("<@", data.Message.Author.ID, ">")
            } else {
                if strings.Contains(args[1], "me") {
        			uid = fmt.Sprint("<@", data.Message.Author.ID, ">")
        		} else {
        			uid = args[1]
        		}
            }
        	output := fmt.Sprintf("Love You %s!", uid)
        	data.Message.RespondString(session, output)
        	return err
		},

		cmdName:     "Love",
		cmdCalls:    []string{"love"},
		cmdMinDesc:  "Loves you",
		cmdFullDesc: "Responds back with either loving you or loving whoever you tag",
		cmdFirstChr: "l",
		cmdModule:   "topatoes",
	}))

    // dab on the h8ters
    cmdarray = append(cmdarray, (cmddata{
        cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {
            var err error

            output := "Dabs"
            data.Message.RespondString(session, output)

            return err
        },

        cmdName:     "Dab", 
        cmdCalls:    []string{"dab"},
        cmdMinDesc:  "Dab dab dab",
        cmdFullDesc: "Dab", //big description
        cmdFirstChr: "d",
        cmdModule:   "topatoes",
    }))
}
