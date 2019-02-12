package main

import (
	"database/sql"
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
	bc.session.ChannelMessageSend(channelID, "@here Um novo carrinho (`#"+id+" "+split[1]+"`) foi criado, faça seu pedido :yum:")
}

// CloseList ...
func CloseList(bc *BotCommand) {
	list, err := models.GetOpenListByChannelID(bc.message.ChannelID)
	if err != nil {
		bc.session.ChannelMessageSend(bc.message.ChannelID, err.Error())
		return
	}

	closed := models.CloseList(&list)
	if closed {
		bc.session.UpdateStatus(0, "Ingredientes na panela.")
		bc.session.ChannelMessageSend(bc.message.ChannelID, "@here **Pedidos finalizados!**")
		return
	}

	bc.session.ChannelMessageSend(bc.message.ChannelID, "@here Por algum motivo o carrinho não pode ser finalizado.")
	return
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

// UpdateItem ...
func UpdateItem(bc *BotCommand) {
	// MOVE THIS TO MIDDLEWARE
	splitRegexp := regexp.MustCompile("[\n| ]")
	split := splitRegexp.Split(bc.message.Content, 2)

	if len(split) == 1 {
		_, err := bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+", digite seu texto. Por exemplo, `!atualizar Feijão extra` :heart:")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// MOVE THIS TO MIDDLEWARE

	lastActiveItem, err := models.GetLastActiveItem(bc.message.Author.ID, bc.message.ChannelID)

	if err == sql.ErrNoRows {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" você não tem um pedido aberto para atualizar.")
		return
	}

	if err != sql.ErrNoRows && err != nil {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" erro ao atualizar ao pedido.")
		return
	}

	updated := models.UpdateItem(lastActiveItem, "\n"+split[1])
	if updated {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" **pedido atualizado** com sucesso.")
		return
	}

	bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" por algum motivo seu pedido não foi atualizado. Entre em contato com um administrador.")
	return
}

// RemoveItem ...
func RemoveItem(bc *BotCommand) {
	var item models.Item
	item.DiscordUserID = bc.message.Author.ID
	item, userErr := models.GetItem(&item, bc.message.ChannelID)
	if userErr != nil {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+": "+userErr.Error())
		return
	}

	deleted := models.RemoveItem(&item)
	if deleted {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" seu pedido foi **cancelado** com sucesso!")
		return
	}

	bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" **não** foi possivel **cancelar** seu pedido!")
	return
}

// RepeatItem ...
func RepeatItem(bc *BotCommand) {
	list, _ := models.GetOpenListByChannelID(bc.message.ChannelID)
	userHasItemInList := models.HasItem(&list, bc.message.Author.ID)

	if userHasItemInList {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" você já realizou seu pedido. Para **cancelar** digite `!cancelar`")
		return
	}

	repeated, userErr := models.RepeatItem(bc.message.Author.ID, bc.message.ChannelID)
	if repeated {
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" seu **pedido foi realizado** com sucesso.")
		return
	}

	bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+": "+userErr.Error())
	return
}

// ListItems ...
func ListItems(bc *BotCommand) {
	list, err := models.GetOpenListByChannelID(bc.message.ChannelID)
	if err != nil {
		bc.session.ChannelMessageSend(bc.message.ChannelID, err.Error())
		return
	}
	items := models.GetItemsByListID(&list)
	list.Items = items

	embedListItems := EmbedListItems(&list, bc)
	bc.session.ChannelMessageSendEmbed(bc.message.ChannelID, embedListItems)
}

// RaffleListItems ...
func RaffleListItems(bc *BotCommand) {
	Chosen, err := models.RaffleList(bc.message.ChannelID)
	if err != nil {
		bc.session.ChannelMessageSend(bc.message.ChannelID, err.Error())
		return
	}
	embedRaffleListItems := EmbedRaffleListItems(Chosen, bc)
	bc.session.ChannelMessageSendEmbed(bc.message.ChannelID, embedRaffleListItems)
}

// RevertListItems ...
func RevertListItems(bc *BotCommand) {
	list, _ := models.GetLastList(bc.message.ChannelID)
	if !list.Status {
		list.Status = true
		updated := models.UpdateList(&list, bc.message.ChannelID)
		if updated {
			bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" seu carrinho foi **re-aberto** com sucesso!")
			return
		}
		bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" não foi possivel **reabrir** seu carrinho!")
		return
	}

	bc.session.ChannelMessageSend(bc.message.ChannelID, bc.message.Author.Mention()+" seu carrinho já esta aberto!")
	return
}
