package dispatcher

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"reflect"
)

type DiscordCommandHandler func(
	session *discordgo.Session, message *discordgo.MessageCreate)

type CommandDefinition struct {
	Handler   DiscordCommandHandler
	ShortHelp string
	LongHelp  string
}

type CommandHandlers map[string]*CommandDefinition

type DiscordModule struct {
	Name                  string
	Listeners             []interface{}
	ServerCommands        CommandHandlers
	DirectMessageCommands CommandHandlers
}

// CreatedModules contains all the created modules,
// but they may not be activated.
var CreatedModules []*DiscordModule

// ActivatedModules contains all the modules that are loaded
// into the discord session.
var ActivatedModules []*DiscordModule

// New creates a new base module from a given Name.
func New(initFunc interface{}) *DiscordModule {
	newModule := &DiscordModule{
		Name:                  reflect.TypeOf(initFunc).PkgPath(),
		ServerCommands:        map[string]*CommandDefinition{},
		DirectMessageCommands: map[string]*CommandDefinition{},
	}

	CreatedModules = append(CreatedModules, newModule)
	return newModule
}

// ActivateModules loads every available and activated modules
// into a given discord session.
func ActivateModules(session *discordgo.Session) {
	// Register global handlers into the discord session
	session.AddHandler(onMessageReceived)

	for _, module := range CreatedModules {
		log.Print("Activating: ", module.Name)

		// TODO - feature:
		//  	we should check whether the module is enabled or not
		module.Activate(session)
	}
}

// FlagEnabled marks the modules as enabled. This doesn't load the module.
func (mod *DiscordModule) FlagEnabled() {
	ActivatedModules = append(ActivatedModules, mod)
}

// Activate adds every registered Listeners
// into a discord session.
func (mod *DiscordModule) Activate(session *discordgo.Session) {
	session.Lock()
	defer session.Unlock()

	for _, listener := range mod.Listeners {
		session.AddHandler(listener)
	}

	mod.FlagEnabled()
}

// AddListener appends a function Handler
// into a module's Listeners list.
func (mod *DiscordModule) AddListener(handler interface{}) {
	mod.Listeners = append(mod.Listeners, handler)
}
