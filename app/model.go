package main

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

// App struct
type App struct {
	Connection *sql.DB
	Session    *discordgo.Session
	Message    *discordgo.MessageCreate
}

// Cart is our itens wrapper
type Cart struct {
	ID          uint
	Description string
	Item        []Item
}

// Item is any order made
type Item struct {
	ID            uint
	Description   string
	DiscordUserID string
}
