package main

import (
	"dlogger"
	"github.com/andersfylling/disgord"
	"fmt"
	"strings"
)

var corecmdslist [255]string

func init() {
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
	dhname			string
	dhalts 			string
	dhdescription 	string
	helshort		string
	cmdtitle		string
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
		dlogger.LogOld(30,35,"responce error", err.Error())
		msg.RespondString(session, "Something seems to have went wrong")
		msg.RespondString(session, err.Error())
	}
	//msg.RespondString(session, responce)
}

func cmdcorehelp(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	//detailedconfcmd := [512]string{"Help"}
	//detailedconfdsc := [512]string{"The help command, it outputs helpy stuff"}
	var err error

	var basichelp = "``` -| Core commands |- \n - Help: The help command, \n - Version: Displays the version running and some other info \n - Ping: Pong!, \n -| Text/Test commands |- \n - Hello: Says hello back```"

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
	output := fmt.Sprintf("%s%s%s%s", content, "<@",data.Message.Author.ID,">")
	data.Message.RespondString(session, output)
	return err
}

func cmdcorelove(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	var content string
	content = "Love You "
	var uid string
	if cmd == "" {
		uid = fmt.Sprint("<@",data.Message.Author.ID,">")
	} else {
		if strings.Contains(cmd, "me") {
			uid = fmt.Sprint("<@",data.Message.Author.ID,">")
		} else {
			uid = cmd
		}
	}
	output := fmt.Sprintf("%s%s", content, uid)
	data.Message.RespondString(session, output)
	return err
}

func cmdcoreping(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	output := "Pong!"
	data.Message.RespondString(session, output)
	return err
}

func cmdcoreversion(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	user, err := session.GetCurrentUser()
	output := fmt.Sprint("Running: ", appname," ", version, "\nlocally configured as: ", conf_Name, "\nRunning under user: ", user.Username,"#",user.Discriminator)
	data.Message.RespondString(session, output)
	return err
}

func cmdcoredab(cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	output := "Dabs"
	data.Message.RespondString(session, output)
	return err
}

func cmdcoredebug (cmd string, session disgord.Session, data *disgord.MessageCreate) error {
	var err error
	output := "u!test"
	data.Message.RespondString(session, output)
	return err
}
