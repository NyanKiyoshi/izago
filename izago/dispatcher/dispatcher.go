package dispatcher

import "github.com/bwmarrin/discordgo"

// RegisterServerCommand registers a command for server messages.
func (mod *DiscordModule) RegisterServerCommand(name string, handler *DiscordCommandHandler) {
	mod.serverCommands[name] = handler
}

// RegisterDMCommand registers a command for direct messages.
func (mod *DiscordModule) RegisterDMCommand(name string, handler *DiscordCommandHandler) {
	mod.directMessageCommands[name] = handler
}

func (commands CommandHandlers) dispatch(
	name string, session *discordgo.Session, message *discordgo.MessageCreate) bool {

	if command, found := commands[name]; found {
		if session.SyncEvents {
			(*command)(session, message)
		} else {
			go (*command)(session, message)
		}
		return true
	}

	return false
}

// dispatchServerCommand dispatches a received guild command
// to the proper module.
func dispatchServerCommand(
	receivedCommand string, session *discordgo.Session, message *discordgo.MessageCreate) {

	// Search for a module that can handle the command.
	for _, module := range ActivatedModules {
		// If the received command was handled by the module,
		// stop searching for an handler, we found it.
		if module.serverCommands.dispatch(receivedCommand, session, message) {
			break
		}
	}
}
