package models

// Item of the list
type Item struct {
	ID            uint
	Description   string
	DiscordUserID string
}

// AddItem ...
func AddItem() {

	// stmt, err := app.Connection.Prepare("INSERT item SET description = ?, cart_id = ?, discord_user_id = ?")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// _, err = stmt.Exec(split[1], cart.ID, m.Author.ID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

// RemoveItem ...
func RemoveItem() {

	// var item Item
	// row := app.Connection.QueryRow("select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = ? and c.channel_id = ?", m.Author.ID, m.ChannelID)
	// err := row.Scan(&item.ID)

	// // select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = "186909290475290624";
	// stmt, err := app.Connection.Prepare("delete from item where id = ?")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// _, err = stmt.Exec(item.ID)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
