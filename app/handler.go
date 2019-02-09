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
	ListItems(bc)
	if models.CloseList(bc.message.ChannelID) {
		bc.session.UpdateStatus(0, "Ingredientes na panela.")
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

	list, _ := models.GetOpenListByChannelID(bc.message.ChannelID)
	userHasItemInList := models.HasItem(&list, bc.message.Author.ID)

	if userHasItemInList {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" você já realizou seu pedido. Para **cancelar** digite `!cancelar`")
		return
	}

	var item models.Item
	item.CartID = list.ID
	item.Description = split[1]
	item.DiscordUserID = bc.message.Author.ID

	added := models.AddItem(&item)
	if added {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" seu **pedido foi realizado** com sucesso.")
		return
	}

	bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" por algum motivo seu pedido não foi realizado. Entre em contato com um administrador.")
	return

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

	bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" **não** foi possivel **cancelar** seu pedido!")
	return
}

// ListItems ...
func ListItems(bc *BotCommand) {
	list, err := models.GetOpenListByChannelID(bc.message.ChannelID)
	if err != nil {
		bc.session.ChannelMessageSend(bc.message.ChannelID, "Não há um carrinho aberto, crie com o comando `!criar [nome]`")
		return
	}
	items := models.GetItemsByListID(&list)
	list.Items = items

	embedListItems := EmbedListItems(&list, bc)
	bc.session.ChannelMessageSendEmbed(bc.message.ChannelID, embedListItems)
}

// RaffleListItems ...
func RaffleListItems(bc *BotCommand) {
	Chosen := models.RaffleList(bc.message.ChannelID)
	embedRaffleListItems := EmbedRaffleListItems(Chosen, bc)
	bc.session.ChannelMessageSendEmbed(bc.message.ChannelID, embedRaffleListItems)
}
