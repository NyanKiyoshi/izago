package managebot

import (
	"github.com/NyanKiyoshi/izago/izago/dispatcher"
)

type _internal struct{}

func init() {
	module := dispatcher.New(_internal{})
	module.AddCommand(
		"status", commandGetBotStatus,
		`Gets the bot's status information.`)
}
