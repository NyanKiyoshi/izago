package dispatcher

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordCommandHandler func(
	session *discordgo.Session, message *discordgo.MessageCreate)

type CommandHandlers map[string]DiscordCommandHandler

type DiscordModule struct {
	name                  string
	listeners             []*func()
	serverCommands        CommandHandlers
	directMessageCommands CommandHandlers
}

// createdModules contains all the created modules,
// but they may not be activated.
var createdModules []*DiscordModule

// activatedModules contains all the modules that are loaded
// into the discord session.
var activatedModules []*DiscordModule

// New creates a new base module from a given name.
func New(name string) *DiscordModule {
	newModule := &DiscordModule{
		name: name,
	}

	createdModules = append(createdModules, newModule)
	return newModule
}

// LoadModules loads every available and activated modules
// into a given discord session.
func LoadModules(session *discordgo.Session) {
	// Register global handlers into the discord session
	session.AddHandler(onMessageReceived)

	for _, module := range createdModules {
		// TODO - feature:
		//  	we should check whether the module is enabled or not
		module.ActivateModule(session)
	}
}

// ActivateModule adds every registered listeners
// into a discord session.
func (mod *DiscordModule) ActivateModule(session *discordgo.Session) {
	session.Lock()
	defer session.Unlock()

	for _, listener := range mod.listeners {
		session.AddHandler(listener)
	}

	activatedModules = append(activatedModules, mod)
}

// RegisterListener appends a function handler
// into a module's listeners list.
func (mod *DiscordModule) RegisterListener(handler *func()) {
	mod.listeners = append(mod.listeners, handler)
}
