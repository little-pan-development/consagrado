package main

import (
	"fmt"
	"regexp"

	"github.com/palmirinha/app/models"
)

// OpenList ...
func OpenList(bc *BotCommand) {
	channelID := bc.message.ChannelID

	// MOVE THIS TO MIDDLEWARE
	splitRegexp := regexp.MustCompile("[\n| ]")
	split := splitRegexp.Split(bc.message.Content, 2)

	if len(split) == 1 {
		_, err := bc.session.ChannelMessageSend(channelID, "Digite uma descrição para seu carrinho!")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// MOVE THIS TO MIDDLEWARE

	rows := models.CountOpenList(channelID)

	if rows > 0 {
		_, err := bc.session.ChannelMessageSend(channelID, "Existe um carrinho em aberto!")
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	id := models.OpenList(split[1], channelID)

	bc.session.UpdateStatus(0, "Faça seu pedido..")
	bc.session.ChannelMessageSend(channelID, "Carrinho `#"+id+" "+split[1]+"` criado com sucesso!")
}

// CloseList ...
func CloseList(bc *BotCommand) {
	// 	// Lista dos pedidos
	// 	embed := app.getCartContentsAsEmbed(m.ChannelID, app.Session)

	if models.CloseList(bc.message.ChannelID) {
		bc.session.UpdateStatus(0, "Ingredientes na panela.")
		// s.ChannelMessageSendEmbed(m.ChannelID, embed)
		// Avisando finalização
		bc.session.ChannelMessageSend(bc.message.ChannelID, "@here **Pedidos finalizados!**")
	}

}

// AddItem ...
func AddItem(bc *BotCommand) {

	// MOVE THIS TO MIDDLEWARE
	splitRegexp := regexp.MustCompile("[\n| ]")
	split := splitRegexp.Split(bc.message.Content, 2)

	if len(split) == 1 {
		_, err := bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+", digite seu pedido. Por exemplo, `!pedir Lentilha da vó` :heart:")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// MOVE THIS TO MIDDLEWARE

	quantityOfList := models.CountOpenList(bc.message.ChannelID)
	if quantityOfList == 0 {
		_, err := bc.session.ChannelMessageSend(bc.message.ChannelID, "Para realizar um pedido é necessário ter um carrinho aberto. `!criar [nome]`")
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	cart := models.GetOpenListByChannelID(bc.message.ChannelID)

	var item models.Item
	item.CartID = cart.ID
	item.Description = split[1]
	item.DiscordUserID = bc.message.Author.ID

	added := models.AddItem(&item)
	if added {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" seu **pedido foi realizado** com sucesso.")
		return
	}

}

// RemoveItem ...
func RemoveItem(bc *BotCommand) {
	var item models.Item
	item.DiscordUserID = bc.message.Author.ID
	item = models.GetItem(&item, bc.message.ChannelID)

	deleted := models.RemoveItem(&item)
	if deleted {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" seu pedido foi **cancelado** com sucesso!")
		return
	}
	return
}

// ListItems ...
func ListItems(bc *BotCommand) {
	list := models.GetOpenListByChannelID(bc.message.ChannelID)
	items := models.GetItemsByListID(&list)
	list.Items = items

	fmt.Println(list)
	// 	embed := app.getCartContentsAsEmbed(m.ChannelID, app.Session)
	// 	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

// RaffleList ...
func RaffleList(bc *BotCommand) {
	// 	var discordUserID string
	// 	row := app.Connection.QueryRow("SELECT i.discord_user_id FROM cart c JOIN item i ON i.cart_id = c.id WHERE c.status = 1 and c.channel_id = ? ORDER BY RAND() LIMIT 1", m.ChannelID)
	// 	err := row.Scan(&discordUserID)

	// 	// Isso pode ser aplicado melhor quando desacoplado
	// 	if err != sql.ErrNoRows {

	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		var user, _ = s.User(discordUserID)

	// 		embed := &discordgo.MessageEmbed{}

	// 		embed.Title = "Parabéns! Hoje é com..."
	// 		embed.Description = user.Mention() + " contamos com você!"
	// 		embed.Color = 0xff0000

	// 		embed.Author = &discordgo.MessageEmbedAuthor{}
	// 		embed.Author.Name = "Palmirinha!"
	// 		embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
	// 		embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

	// 		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	// 	} else {
	// 		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" **não há pedidos para sortear.**")
	// 	}
}
