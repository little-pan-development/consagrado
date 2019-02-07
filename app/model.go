package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"strconv"

// 	"github.com/bwmarrin/discordgo"
// )

// // App struct
// type App struct {
// 	Connection *sql.DB
// 	Session    *discordgo.Session
// 	Message    *discordgo.MessageCreate
// }

// // Cart is our itens wrapper
// type Cart struct {
// 	ID          uint
// 	Description string
// 	Item        []Item
// }

// // Item is any order made
// type Item struct {
// 	ID            uint
// 	Description   string
// 	DiscordUserID string
// }

// func (app *App) countOpenOrderByChannelId(channelID string) uint {

// 	var count uint
// 	query := `SELECT COUNT(*) FROM cart WHERE status = 1 and channel_id = ?`
// 	rows, err := app.Connection.Query(query, app.Message.ChannelID)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if rows.Next() {
// 		rows.Scan(&count)
// 	}

// 	return count
// }

// func (app *App) createOrderByChannel(description, channelID string) string {

// 	query := `INSERT cart SET description = ?, status = ?, channel_id = ?`
// 	stmt, _ := app.Connection.Prepare(query)
// 	res, _ := stmt.Exec(description, 1, channelID)

// 	id, _ := res.LastInsertId()
// 	idToString := strconv.FormatInt(int64(id), 10)

// 	return idToString
// }
