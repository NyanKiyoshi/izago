package managebot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var getHelpExpectedHelpText = "```Usage:\n" +
	">servercommand  This test passes as well.```"

func Test_getHelp(t *testing.T) {
	session := &discordgo.Session{}

	t.Run("Non cached test", func(t *testing.T) {
		assert.Empty(t, helpText)
		assert.Equal(t, getHelpExpectedHelpText, getHelp(session))
	})

	t.Run("Cached test", func(t *testing.T) {
		assert.NotEmpty(t, helpText)
		assert.Equal(t, getHelpExpectedHelpText, getHelp(session))
	})
}

var getCommandHelpTextTests = []struct {
	isDm        bool
	commandName string
	returnValue string
}{
	// Non existing command
	{false, "hello", "help: no such command (hello)."},

	// Existing command
	{
		// DM command
		true,
		"dmcommand",
		"```>dmcommand:\nThis test passes.```",
	},
	{
		// Guild command
		false,
		"servercommand",
		"```>servercommand:\nThis test passes as well.```",
	},
}

func Test_getCommandHelpText(t *testing.T) {
	for _, tt := range getCommandHelpTextTests {
		t.Run(tt.commandName, func(t *testing.T) {
			assert.Equal(t, tt.returnValue, getCommandHelpText(tt.commandName, tt.isDm))
		})
	}
}
