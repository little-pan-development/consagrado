package main

import (
	"fmt"
	"regexp"
)

// OpenList ...
func OpenList(bc *BotCommand) {
	channelID := bc.message.ChannelID

	// ADD THIS TO MIDDLEWARE
	splitRegexp := regexp.MustCompile("[\n| ]")
	split := splitRegexp.Split(bc.message.Content, 2)

	if len(split) == 1 {
		_, err := bc.session.ChannelMessageSend(channelID, "Digite uma descrição para seu carrinho!")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// ADD THIS TO MIDDLEWARE

	// rows := app.countOpenOrderByChannelId(channelID)

	// if rows > 0 {
	// 	_, err := s.ChannelMessageSend(channelID, "Existe um carrinho em aberto!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	return
	// }

	// id := app.createOrderByChannel(split[1], channelID)

	// b.Session.UpdateStatus(0, "Faça seu pedido..")
	// s.ChannelMessageSend(channelID, "Carrinho `#"+id+" "+split[1]+"` criado com sucesso!")
	// b.Session.ChannelMessageSend(channelID, "Carrinho `#lalala criado com sucesso!")
}

// CloseList ...
// func CloseList(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	// Lista dos pedidos
// 	embed := app.getCartContentsAsEmbed(m.ChannelID, app.Session)

// 	if models.CloseList(m.ChannelID) {
// 		s.UpdateStatus(0, "Ingredientes na panela.")
// 		s.ChannelMessageSendEmbed(m.ChannelID, embed)
// 		// Avisando finalização
// 		s.ChannelMessageSend(m.ChannelID, "@here **Pedidos finalizados!**")
// 	}

// }

// // AddItem ...
// func AddItem(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	splitRegexp := regexp.MustCompile("[\n| ]")
// 	split := splitRegexp.Split(m.Content, 2)

// 	if len(split) == 1 {
// 		_, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+", digite seu pedido. Por exemplo, `!pedir Lentilha da vó` :heart:")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		return
// 	}

// 	var cart Cart
// 	row := app.Connection.QueryRow("SELECT id, description FROM cart WHERE status = 1 and channel_id = ?", m.ChannelID)
// 	err := row.Scan(&cart.ID, &cart.Description)

// 	switch err {
// 	case sql.ErrNoRows:
// 		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+", antes de pedirem, utilize `!criar nome do carrinho` para **criar um novo carrinho**.")
// 		return
// 	default:
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}

// 	_, err = app.Connection.Query("SELECT COUNT(*) FROM item WHERE discord_user_id = ? AND cart_id = ?", m.Author.ID, cart.ID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// if checkCount(rows) > 0 {
// 	// 	_, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" você já realizou seu pedido. Para **cancelar** digite `!cancelar`")
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// 	return
// 	// }

// 	stmt, err := app.Connection.Prepare("INSERT item SET description = ?, cart_id = ?, discord_user_id = ?")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	_, err = stmt.Exec(split[1], cart.ID, m.Author.ID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" seu **pedido foi realizado** com sucesso.")
// }

// // RemoveItem ...
// func RemoveItem(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	var item Item
// 	row := app.Connection.QueryRow("select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = ? and c.channel_id = ?", m.Author.ID, m.ChannelID)
// 	err := row.Scan(&item.ID)

// 	// select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = "186909290475290624";
// 	stmt, err := app.Connection.Prepare("delete from item where id = ?")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	_, err = stmt.Exec(item.ID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" seu pedido foi **cancelado** com sucesso!")
// }

// // ItemsByList ...
// func ItemsByList(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	embed := app.getCartContentsAsEmbed(m.ChannelID, app.Session)
// 	s.ChannelMessageSendEmbed(m.ChannelID, embed)
// }

// // raffle
// func raffle(s *discordgo.Session, m *discordgo.MessageCreate) {

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
// }
