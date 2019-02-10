package models

import (
	"database/sql"
	"errors"
	"fmt"
)

// Item of the list
type Item struct {
	ID            uint
	CartID        uint
	DiscordUserID string
	Description   string
}

// AddItem ...
func AddItem(item *Item) bool {
	query := `
		INSERT item 
		SET description = ?, cart_id = ?, discord_user_id = ?
	`
	stmt, err := Connection.Mysql.Prepare(query)
	if err != nil {
		fmt.Println("Model AddItem [prepare]: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.Description, item.CartID, item.DiscordUserID)
	if err != nil {
		fmt.Println("Model AddItem [exec]: ", err)
		return false
	}

	return true
}

// RepeatItem ...
func RepeatItem(DiscordUserID, ChannelID string) (added bool, userErr error) {
	item, err := GetLastItem(DiscordUserID, ChannelID)
	if err != nil {
		fmt.Println("Model RepeatItem / GetLastItem: ", err)
		return false, err
	}

	list, err := GetOpenListByChannelID(ChannelID)
	if err != nil {
		fmt.Println("Model RepeatItem / GetOpenListByChannelID: ", err)
		return false, err
	}

	item.CartID = list.ID
	added = AddItem(&item)
	return added, userErr
}

// GetLastItem ...
func GetLastItem(discordUserID, channelID string) (item Item, userErr error) {
	query := `
		SELECT item.id, item.description, item.discord_user_id 
		FROM item
		INNER JOIN cart ON item.cart_id = cart.id
		WHERE cart.channel_id = ?
		AND item.discord_user_id = ? 
		ORDER BY item.created_at DESC
	`
	row := Connection.Mysql.QueryRow(query, channelID, discordUserID)
	err := row.Scan(&item.ID, &item.Description, &item.DiscordUserID)

	if err == sql.ErrNoRows {
		userErr = errors.New("Este é seu primeiro pedido. Portanto não encontramos pedidos anteriores")
	} else if err != nil {
		userErr = errors.New("Erro ao encontrar último pedido")
		fmt.Println("Model GetLastItem [scan]: ", err, sql.ErrNoRows)
	}

	return item, userErr
}

// GetLastActiveItem ...
func GetLastActiveItem(list *List, author string) (itemID int, err error) {
	query := `
		SELECT item.id
		FROM item
		INNER JOIN cart ON item.cart_id = cart.id
		WHERE cart.status = TRUE
		AND cart.channel_id = ?
		AND item.discord_user_id = ?
		LIMIT 1
	`
	row := Connection.Mysql.QueryRow(query, list.channelID, author)
	err = row.Scan(&itemID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Model getLastActiveItem [scan]: ", err)
	}

	return itemID, err
}

// UpdateItem ...
func UpdateItem(itemID int, description string) bool {
	query := `
		UPDATE item 
		SET description = CONCAT(description, ?)
		WHERE id = ?
	`
	stmt, err := Connection.Mysql.Prepare(query)
	if err != nil {
		fmt.Println("Model UpdateItem [prepare]: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(description, itemID)
	if err != nil {
		fmt.Println("Model UpdateItem [exec]: ", err)
		return false
	}

	return true
}

// RemoveItem ...
func RemoveItem(item *Item) bool {
	query := `
		DELETE 
		FROM item 
		WHERE id = ?
	`
	stmt, err := Connection.Mysql.Prepare(query)
	if err != nil {
		fmt.Println("Model RemoveItem [prepare]: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(&item.ID)
	if err != nil {
		fmt.Println("Model RemoveItem [exec]: ", err)
		return false
	}

	return true
}

// GetItem ...
func GetItem(item *Item, channelID string) (getItem Item, userErr error) {
	query := `
		SELECT item.id, item.description, item.discord_user_id 
		FROM item
		INNER JOIN cart ON item.cart_id = cart.id
		WHERE cart.status = TRUE
		AND cart.channel_id = ?
		AND item.discord_user_id = ?
	`
	row := Connection.Mysql.QueryRow(query, channelID, item.DiscordUserID)
	err := row.Scan(&getItem.ID, &getItem.Description, &getItem.DiscordUserID)

	if err == sql.ErrNoRows {
		userErr = errors.New("Você ainda não realizou nenhum pedido. Digite `!pedir [pedido]`")
	} else if err != nil {
		userErr = errors.New("Erro ao buscar último carrinho criado")
		fmt.Println("Model GetItem [scan]: ", err, sql.ErrNoRows)
	}

	return getItem, userErr
}

// GetItemsByListID ...
func GetItemsByListID(list *List) []Item {
	var items []Item
	query := `
		SELECT description, discord_user_id 
		FROM item 
		WHERE cart_id = ?
	`
	rows, err := Connection.Mysql.Query(query, list.ID)
	if err != nil {
		fmt.Println("Model GetItemsByListID [query]: ", err)
	}

	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Description, &item.DiscordUserID)
		if err != nil {
			fmt.Println("Model GetItemsByListID [next]: ", err)
		}

		items = append(items, item)
	}

	return items
}

// HasItem ...
func HasItem(list *List, author string) bool {
	var hasItem bool
	query := `
		SELECT IF(COUNT(id) > 0, true, false) as hasItem
		FROM item 
		WHERE discord_user_id = ? 
		AND cart_id = ?
		LIMIT 1
	`
	row := Connection.Mysql.QueryRow(query, author, list.ID)
	err := row.Scan(&hasItem)
	if err != nil {
		fmt.Println("Model HasItem [scan]: ", err)
	}

	if !hasItem {
		return false
	}

	return true
}
