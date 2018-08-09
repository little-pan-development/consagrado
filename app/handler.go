package main

import (
	"database/sql"
	"regexp"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

// Create a order / cart
func (app *App) createOrder(client *Client) {

	channelID := app.Message.ChannelID

	splitRegexp := regexp.MustCompile("[\n| ]")
	split := splitRegexp.Split(app.Message.Content, 2)

	if len(split) == 1 {
		_, err := app.Session.ChannelMessageSend(channelID, "Digite uma descrição para seu carrinho!")
		checkErr(err)
		return
	}

	rows := app.countOpenOrderByChannelId(channelID)

	if rows > 0 {
		_, err := app.Session.ChannelMessageSend(channelID, "Existe um carrinho em aberto!")
		checkErr(err)
		return
	}

	id := app.createOrderByChannel(split[1], channelID)

	app.Session.UpdateStatus(0, "Faça seu pedido..")
	app.Session.ChannelMessageSend(channelID, "Carrinho `#"+id+" "+split[1]+"` criado com sucesso!")
}

// close a order / cart
func (app *App) closeOrder(client *Client) {
	// Lista dos pedidos
	embed := app.getCartContentsAsEmbed(app.Message.ChannelID, app.Session)

	// Desabilita carrinho
	stmt, err := app.Connection.Prepare("update cart set status = ? where status = ? and channel_id = ?")
	checkErr(err)

	_, err = stmt.Exec(0, 1, app.Message.ChannelID)
	checkErr(err)

	app.Session.UpdateStatus(0, "Ingredientes na panela.")

	app.Session.ChannelMessageSendEmbed(app.Message.ChannelID, embed)

	// Avisando finalização
	app.Session.ChannelMessageSend(app.Message.ChannelID, "@here **Pedidos finalizados!**")
}

// add item
func (app *App) addItem(client *Client) {
	splitRegexp := regexp.MustCompile("[\n| ]")
	split := splitRegexp.Split(app.Message.Content, 2)

	if len(split) == 1 {
		_, err := app.Session.ChannelMessageSend(app.Message.ChannelID, app.Message.Author.Mention()+", digite seu pedido. Por exemplo, `!pedir Lentilha da vó` :heart:")
		checkErr(err)
		return
	}

	var cart Cart
	row := app.Connection.QueryRow("SELECT id, description FROM cart WHERE status = 1 and channel_id = ?", app.Message.ChannelID)
	err := row.Scan(&cart.ID, &cart.Description)

	switch err {
	case sql.ErrNoRows:
		app.Session.ChannelMessageSend(app.Message.ChannelID, app.Message.Author.Mention()+", antes de pedirem, utilize `!criar nome do carrinho` para **criar um novo carrinho**.")
		return
	default:
		checkErr(err)
	}

	rows, err := app.Connection.Query("SELECT COUNT(*) FROM item WHERE discord_user_id = ? AND cart_id = ?", app.Message.Author.ID, cart.ID)
	checkErr(err)

	if checkCount(rows) > 0 {
		_, err := app.Session.ChannelMessageSend(app.Message.ChannelID, app.Message.Author.Mention()+" você já realizou seu pedido. Para **cancelar** digite `!cancelar`")
		checkErr(err)
		return
	}

	stmt, err := app.Connection.Prepare("INSERT item SET description = ?, cart_id = ?, discord_user_id = ?")
	checkErr(err)

	_, err = stmt.Exec(split[1], cart.ID, app.Message.Author.ID)
	checkErr(err)

	app.Session.ChannelMessageSend(app.Message.ChannelID, app.Message.Author.Mention()+" seu **pedido foi realizado** com sucesso.")
}

// remove item
func (app *App) removeItem(client *Client) {
	var item Item
	row := app.Connection.QueryRow("select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = ? and c.channel_id = ?", app.Message.Author.ID, app.Message.ChannelID)
	err := row.Scan(&item.ID)

	// select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = "186909290475290624";
	stmt, err := app.Connection.Prepare("delete from item where id = ?")
	checkErr(err)

	_, err = stmt.Exec(item.ID)
	checkErr(err)

	app.Session.ChannelMessageSend(app.Message.ChannelID, app.Message.Author.Mention()+" seu pedido foi **cancelado** com sucesso!")
}

// list items
func (app *App) listItems(client *Client) {
	embed := app.getCartContentsAsEmbed(app.Message.ChannelID, app.Session)
	app.Session.ChannelMessageSendEmbed(app.Message.ChannelID, embed)
}

// raffle
func (app *App) raffle(client *Client) {

	var discordUserID string
	row := app.Connection.QueryRow("SELECT i.discord_user_id FROM cart c JOIN item i ON i.cart_id = c.id WHERE c.status = 1 and c.channel_id = ? ORDER BY RAND() LIMIT 1", app.Message.ChannelID)
	err := row.Scan(&discordUserID)

	// Isso pode ser aplicado melhor quando desacoplado
	if err != sql.ErrNoRows {

		checkErr(err)

		var user, _ = app.Session.User(discordUserID)

		embed := &discordgo.MessageEmbed{}

		embed.Title = "Parabéns! Hoje é com..."
		embed.Description = user.Mention() + " contamos com você!"
		embed.Color = 0xff0000

		embed.Author = &discordgo.MessageEmbedAuthor{}
		embed.Author.Name = "Palmirinha!"
		embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
		embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

		app.Session.ChannelMessageSendEmbed(app.Message.ChannelID, embed)
	} else {
		app.Session.ChannelMessageSend(app.Message.ChannelID, app.Message.Author.Mention()+" **não há pedidos para sortear.**")
	}
}
