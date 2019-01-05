package main // this is just a part of the golang build system (don't worry about it) you only need it once per file at the start

import "github.com/andersfylling/disgord" // tells the build system that this file needs the disgord library (discord communication library)
// only needed once per file so don't worry about this
// you may need to add in other imports for some things but disgord is the only one that is required in this context

import (
    "strings"
    "fmt"
)

// oh yea and btw // marks a comment (aka just text snippets which don't count as part of the code)
// i'll be using these to explain some things

// initialisation function, you only need one of these per file and the commands will go in here,
func init() {

    //the gay command
	cmdarray = append(cmdarray, (cmddata{ // this adds a command entry what follows below is all the info about the command

		// this is the function which basically contains all the logic and such to make a command work
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {

			var err error // you can ignore this (its needed but you don't need to mess with it at all)

			output := " GAYYYYYYYY \n :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: :gay_pride_flag: " // this creates a variable which contains the text which we will send back on the next line
			data.Message.RespondString(session, output)                                                                                       // responds with the message

			return err // just like the above var err error thing you can ignore it
		},

		cmdName:     "Gay",                    //full command name (currently unused in code)
		cmdCalls:    []string{"gay", "gae"},   //the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
		cmdMinDesc:  "A very gay command",     //smol description
		cmdFullDesc: "The Gayest of commands", //big description
		cmdFirstChr: "g",                      //first letter of the command
		cmdModule:   "topatoes",               //the module which the command exists in
	}))

    //this command loves you
	cmdarray = append(cmdarray, (cmddata{ // this adds a command entry what follows below is all the info about the command

		// this is the function which basically contains all the logic and such to make a command work
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

		cmdName:     "Love",                    //full command name (currently unused in code)
		cmdCalls:    []string{"love"},   //the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
		cmdMinDesc:  "Loves you",     //smol description
		cmdFullDesc: "Responds back with either loving you or loving whoever you tag", //big description
		cmdFirstChr: "l",                      //first letter of the command
		cmdModule:   "topatoes",               //the module which the command exists in
	}))

    // dab on the h8ters
    cmdarray = append(cmdarray, (cmddata{ // this adds a command entry what follows below is all the info about the command

        // this is the function which basically contains all the logic and such to make a command work
        cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {

            var err error // you can ignore this (its needed but you don't need to mess with it at all)

            output := "Dabs" // this creates a variable which contains the text which we will send back on the next line
            data.Message.RespondString(session, output) // responds with the message

            return err // just like the above var err error thing you can ignore it
        },

        cmdName:     "Dab",                                             //full command name (currently unused in code)
        cmdCalls:    []string{"dab"},                                       //the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
        cmdMinDesc:  "Dab dab dab",                                          //smol description
        cmdFullDesc: "Dab", //big description
        cmdFirstChr: "d",                                                           //first letter of the command
        cmdModule:   "topatoes",                                                     //the module which the command exists in
    }))
}
