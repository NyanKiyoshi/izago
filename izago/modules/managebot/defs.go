package managebot

import "izago/izago/dispatcher"

var Module = dispatcher.New("botManagement")

func init() {
	Module.AddCommand("?status", commandGetBotStatus)
}
