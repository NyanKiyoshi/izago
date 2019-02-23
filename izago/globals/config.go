package globals

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

// Configuration contains all the different
// settings that the bot can have.
type Configuration struct {
	// The bot's discord token. Get one at: https://discordapp.com/developers/applications/.
	Token string

	// The bot's commands prefix.
	Prefix string

	// The bot disabled modules
	// (e.g.: `github.com/NyanKiyoshi/izago/izago/modules/playingstatus`)
	DisabledModules []string
}

// Config stores the loaded bot's globals.
var Config = Configuration{
	Prefix: ">",
}

func init() {
	configFilename := os.Getenv("CONFIG_PATH")

	// Retrieve the globals file path from the arguments
	flag.StringVar(&configFilename, "c", configFilename, "The config file path.")
	flag.Parse()

	// If not globals path was supplied, show the usage
	if configFilename == "" {
		flag.Usage()
		os.Exit(1)
	}

	viper.SetConfigFile(configFilename)

	// Attempt to read the globals file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read the config file: ", err)
	}

	// Attempt to decode the read globals into a struct
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal("failed to decode the globals into a struct: ", err)
	}

	// Check if the bot token was supplied
	if Config.Token == "" {
		log.Fatal("the globals is requiring 'token' to be set")
	}
}
