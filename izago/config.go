package izago

import (
	"flag"
	"fmt"
	_ "github.com/NyanKiyoshi/izago/izago/modules"
	"os"
)

// Config contains the basic bot's configuration
// passed through command line arguments (flags).
var Config = struct {
	Token string
}{}

func init() {
	flag.StringVar(&Config.Token, "token", "", "The bot's login token.")
	flag.Parse()

	if Config.Token == "" {
		fmt.Println("Error: missing login token")
		flag.Usage()
		os.Exit(1)
	}
}
