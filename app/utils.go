package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/palmirinha/app/models"
)

// EmbedListItems ...
func EmbedListItems(list *models.List, bc *BotCommand) *discordgo.MessageEmbed {

	embed := &discordgo.MessageEmbed{}
	embed.Author = &discordgo.MessageEmbedAuthor{}
	embed.Fields = []*discordgo.MessageEmbedField{}

	embed.Author.Name = "Palmirinha!"
	embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
	embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

	embed.Title = "**Comanda:** __" + list.Description + "__"
	embed.Description = "**" + strconv.Itoa(len(list.Items)) + "** pedido(s) até o momento:"
	embed.Color = 0xff0000

	for _, item := range list.Items {
		var user, _ = bc.session.User(item.DiscordUserID)
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "\n\n** " + user.Username + "**",
			Value:  item.Description,
			Inline: false,
		})
	}

	return embed
}

// EmbedRaffleListItems ...
func EmbedRaffleListItems(Chosen string, bc *BotCommand) *discordgo.MessageEmbed {
	var user, _ = bc.session.User(Chosen)

	embed := &discordgo.MessageEmbed{}
	embed.Author = &discordgo.MessageEmbedAuthor{}

	embed.Author.Name = "Palmirinha!"
	embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
	embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

	embed.Title = "Parabéns! Hoje é com..."
	embed.Description = user.Mention() + " contamos com você!"
	embed.Color = 0xff0000

	return embed
}
