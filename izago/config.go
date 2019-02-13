package izago

import (
	"flag"
	"fmt"
	"os"
)

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
