package main // this is just a part of the golang build system (don't worry about it) you only need it once per file at the start

import "github.com/andersfylling/disgord" // tells the build system that this file needs the disgord library (discord communication library)
// only needed once per file so don't worry about this
// you may need to add in other imports for some things but disgord is the only one that is required in this context

// oh yea and btw // marks a comment (aka just text snippets which don't count as part of the code)
// i'll be using these to explain some things

// initialisation function, you only need one of these per file and the commands will go in here,
func init() {

	// ok so this block is everything needed to add a command

	cmdarray = append(cmdarray, (cmddata{ // this adds a command entry what follows below is all the info about the command

		// this is the function which basically contains all the logic and such to make a command work
		cmdFunc: func(args []string, session disgord.Session, data *disgord.MessageCreate) error {

			var err error // you can ignore this (its needed but you don't need to mess with it at all)

			output := "Here is an example of a command" // this creates a variable which contains the text which we will send back on the next line
			data.Message.RespondString(session, output) // responds with the message

			return err // just like the above var err error thing you can ignore it
		},

		cmdName:     "Example command",												//full command name (currently unused in code)
		cmdCalls:    []string{"excmd", "ex"},										//the names it'll use to tell a command (you can even add multiple aliases, the only thing is they all have to start with the same letter)
		cmdMinDesc:  "An Example command",											//smol description
		cmdFullDesc: "An example command for showing how commands are implemented",	//big description
		cmdFirstChr: "e",															//first letter of the command
		cmdModule:   "example",														//the module which the command exists in
	}))
}
