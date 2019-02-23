package managebot

import "github.com/NyanKiyoshi/izago/izago/dispatcher"

func init() {
	module := dispatcher.New(_internal{})
	module.AddDMCommand("dmcommand", nil, "This test passes.")
	module.AddCommand("servercommand", nil, "This test passes as well.")
	module.FlagEnabled()
}
