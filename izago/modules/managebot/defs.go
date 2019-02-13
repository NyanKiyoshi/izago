package managebot

import "izago/izago/dispatcher"

var module = dispatcher.New("botManagement")

func init() {
	module.AddCommand("?status", commandGetBotStatus)
}
