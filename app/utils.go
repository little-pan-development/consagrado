package main

import (
	"github.com/bwmarrin/discordgo"
)

func (app *App) getCartContentsAsEmbed(channelID string, s *discordgo.Session) *discordgo.MessageEmbed {
	var cart Cart
	row := app.Connection.QueryRow("SELECT id, description FROM cart WHERE status = 1 and channel_id = ?", channelID)
	err := row.Scan(&cart.ID, &cart.Description)

	rows, err := app.Connection.Query("SELECT description, discord_user_id FROM item WHERE cart_id = ?", cart.ID)

	embed := &discordgo.MessageEmbed{}

	embed.Title = "Pedidos até o momento:"
	embed.Description = "**--** :hamburger: **--**"
	embed.Color = 0xff0000

	embed.Author = &discordgo.MessageEmbedAuthor{}
	embed.Author.Name = "Palmirinha!"
	embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
	embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

	embed.Fields = []*discordgo.MessageEmbedField{}

	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Description, &item.DiscordUserID)
		checkErr(err)

		var user, _ = s.User(item.DiscordUserID)

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "\n\n**" + user.Username + "**",
			Value:  item.Description,
			Inline: false,
		})
	}

	return embed
}
