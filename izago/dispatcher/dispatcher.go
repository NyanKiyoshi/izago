package dispatcher

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func createCommand(handler DiscordCommandHandler, helpText string) *CommandDefinition {
	return &CommandDefinition{
		Handler:   handler,
		ShortHelp: strings.SplitN(helpText, "\n", 1)[0],
		LongHelp:  helpText,
	}
}

// AddCommand registers a command for server messages.
func (mod *DiscordModule) AddCommand(name string, handler DiscordCommandHandler, helpText string) {
	mod.ServerCommands[name] = createCommand(handler, helpText)
}

// AddDMCommand registers a command for direct messages.
func (mod *DiscordModule) AddDMCommand(name string, handler DiscordCommandHandler, helpText string) {
	mod.DirectMessageCommands[name] = createCommand(handler, helpText)
}

// Get retrieves the command definition of the given command name.
// Returns nil if the command was not found.
func (commands CommandHandlers) Get(name string) *CommandDefinition {
	if command, found := commands[name]; found {
		return command
	}

	return nil
}

// FindCommand looks for a given command name
// of a given scope (DMs or Guid only command) in the loaded modules.
func FindCommand(name string, isDm bool) *CommandDefinition {
	var command *CommandDefinition

	// Search for a module that can handle the command.
	for _, module := range ActivatedModules {
		// If the received command was sent in a guild, dispatch
		// it to the guild commands handlers.
		if !isDm {
			command = module.ServerCommands.Get(name)
		} else {
			command = module.DirectMessageCommands.Get(name)
		}

		// If the command was found,
		// we stop searching and we return it.
		if command != nil {
			return command
		}
	}

	return nil
}

// dispatchCommand dispatches a received guild command
// to the proper module.
func dispatchCommand(
	receivedCommand string, session *discordgo.Session, message *discordgo.MessageCreate) {

	if command := FindCommand(receivedCommand, message.GuildID == ""); command != nil {
		if session.SyncEvents {
			command.Handler(session, message)
		} else {
			go command.Handler(session, message)
		}
	}
}
