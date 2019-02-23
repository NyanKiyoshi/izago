package managebot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dispatchReceivedMessageTests = []struct {
	message     *discordgo.Message
	returnValue string
}{
	// Invalid commands
	{&discordgo.Message{Content: ">help"}, ""},
	{&discordgo.Message{Content: ">something"}, ""},
	{&discordgo.Message{Content: ">help    "}, ""},
	{&discordgo.Message{Content: ">"}, ""},
	{&discordgo.Message{Content: ""}, ""},

	// Non existing command
	{&discordgo.Message{Content: ">help invalid"}, "help: no such command (invalid)."},

	// Existing command
	{
		// DM command
		&discordgo.Message{Content: ">help dmcommand"},
		"```>dmcommand:\nThis test passes.```",
	},
	{
		// Guild command
		&discordgo.Message{Content: ">help servercommand", GuildID: "123"},
		"```>servercommand:\nThis test passes as well.```",
	},
}

func Test_dispatchReceivedMessage(t *testing.T) {
	for _, tt := range dispatchReceivedMessageTests {
		t.Run(tt.message.Content, func(t *testing.T) {
			assert.Equal(t, tt.returnValue, dispatchReceivedMessage(tt.message))
		})
	}
}
