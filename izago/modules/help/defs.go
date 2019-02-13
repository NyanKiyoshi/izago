package managebot

import (
	"github.com/NyanKiyoshi/izago/izago/dispatcher"
)

type _internal struct{}

func init() {
	module := dispatcher.New(_internal{})
	module.AddCommand(
		"help", commandGetHelp,
		`Gets the list of commands.`)
	module.AddListener(onMessageReceived)
}
