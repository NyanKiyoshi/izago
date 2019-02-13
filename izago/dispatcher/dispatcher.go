package dispatcher

import "github.com/bwmarrin/discordgo"

// AddCommand registers a command for server messages.
func (mod *DiscordModule) AddCommand(name string, handler DiscordCommandHandler) {
	mod.serverCommands[name] = handler
}

// AddDMCommand registers a command for direct messages.
func (mod *DiscordModule) AddDMCommand(name string, handler DiscordCommandHandler) {
	mod.directMessageCommands[name] = handler
}

func (commands CommandHandlers) dispatch(
	name string, session *discordgo.Session, message *discordgo.MessageCreate) bool {

	if command, found := commands[name]; found {
		if session.SyncEvents {
			(command)(session, message)
		} else {
			go (command)(session, message)
		}
		return true
	}

	return false
}

// dispatchCommand dispatches a received guild command
// to the proper module.
func dispatchCommand(
	receivedCommand string, session *discordgo.Session, message *discordgo.MessageCreate) {

	// Search for a module that can handle the command.
	for _, module := range activatedModules {
		// If the received command was sent in a guild, dispatch
		// it to the guild commands handlers.
		//
		// If it was handled by the module,
		// we stop searching for an handler as we found it.
		if message.GuildID != "" {
			if module.serverCommands.dispatch(receivedCommand, session, message) {
				break
			}
		} else {
			// Dispatch the DM command
			if module.directMessageCommands.dispatch(receivedCommand, session, message) {
				break
			}
		}
	}
}
