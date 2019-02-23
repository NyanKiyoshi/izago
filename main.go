package main

import (
	"github.com/NyanKiyoshi/izago/izago/dispatcher"
	"github.com/NyanKiyoshi/izago/izago/globals"
	// This is required to import all the modules of the bot
	_ "github.com/NyanKiyoshi/izago/izago/modules"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func createBot() *discordgo.Session {
	session, err := discordgo.New("Bot " + globals.Config.Token)

	if err != nil {
		log.Panic("error creating Discord session: ", err.Error())
	}

	dispatcher.ActivateModules(session)
	return session
}

// runForever waits here until CTRL-C or other term signal is received.
func runForever() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func main() {
	session := createBot()

	// Open a websocket connection to Discord and begin listening.
	if err := session.Open(); err != nil {
		log.Panic("error opening connection: ", err.Error())
	}

	// Print invite link
	log.Printf(
		"Bot is ready. Press CTRL-C to exit."+
			"\n\tInvite: https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot\n",
		session.State.User.ID)

	// Wait for kill for ever
	runForever()

	// Safely close the session
	_ = session.Close()
}
