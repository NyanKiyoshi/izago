package managebot

import "github.com/NyanKiyoshi/izago/izago/dispatcher"

var Module = dispatcher.New("botManagement")

func init() {
	Module.AddCommand("?status", commandGetBotStatus)
}
