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

// EmbedHelpList ...
func EmbedHelpList() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{}

	embed.Title = "Calma, vou te ajudar"
	embed.Description = "**--** :heart: **--**"
	embed.Color = 0x00ff00

	embed.Author = &discordgo.MessageEmbedAuthor{}
	embed.Author.Name = "Palmirinha!"
	embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
	embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

	embed.Fields = []*discordgo.MessageEmbedField{}

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!ajuda**",
		Value:  "Exibe esta tela de ajuda",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!criar**",
		Value:  "Cria um carrinho para você e todos do canal colocarem seus pedidos.",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!pedir Marmitex**",
		Value:  "Faz o pedido `Marmitex` em seu nome no carrinho criado neste canal.",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!adicionar Batata Extra**",
		Value:  "Adiciona `Batata Extra` ao seu último pedido aberto.",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!cancelar**",
		Value:  "Cancela o seu pedido no carrinho deste canal.",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!pedidos**",
		Value:  "Lista todos pedidos do carrinho aberto neste canal.",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!sortear**",
		Value:  "Só pode antes de finalizar. Seleciona uma pessoa aleatória dentre os pedidos do carrinho aberto para pedir hoje!",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!finalizar**",
		Value:  "Finaliza carrinho aberto no canal e lista todos os pedidos do mesmo.",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "**!chegou**",
		Value:  "Avisa no canal que a comida chegou.",
		Inline: false,
	})

	return embed
}
