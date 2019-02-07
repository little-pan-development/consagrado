package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/palmirinha/app/models"
)

// EmbedListItems ...
func EmbedListItems(list *models.List, bc *BotCommand) *discordgo.MessageEmbed {

	embed := &discordgo.MessageEmbed{}
	embed.Fields = []*discordgo.MessageEmbedField{}

	embed.Title = "Pedidos até o momento:"
	embed.Description = "**--** :hamburger: **--**"
	embed.Color = 0xff0000

	for _, item := range list.Items {
		var user, _ = bc.session.User(item.DiscordUserID)
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "\n\n**" + user.Username + "**",
			Value:  item.Description,
			Inline: false,
		})
	}

	return embed
}
