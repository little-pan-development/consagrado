package main

import "github.com/bwmarrin/discordgo"

// FindHandler ...
type FindHandler func(string) (Handler, bool)

// BotCommand ...
type BotCommand struct {
	findHandler FindHandler
	session     *discordgo.Session
	message     *discordgo.MessageCreate
}

// NewBotCommand ...
func NewBotCommand(findHandler FindHandler, s *discordgo.Session, m *discordgo.MessageCreate) *BotCommand {
	return &BotCommand{
		findHandler: findHandler,
		session:     s,
		message:     m,
	}
}
