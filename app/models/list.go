package models

import "fmt"

// OpenList ...
func OpenList(app *App) {

}

// CloseList ...
func CloseList(app *App, channelID string) bool {

	query := `UPDATE cart SET status = ? WHERE status = ? AND channel_id = ?`
	stmt, err := app.Connection.Prepare(query)
	if err != nil {
		fmt.Println("Model CloseList [prepare]: ", err)
		return false
	}

	_, err = stmt.Exec(0, 1, channelID)
	if err != nil {
		fmt.Println("Model CloseList [exec]: ", err)
		return false
	}

	return true
}
