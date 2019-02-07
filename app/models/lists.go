package models

// OpenList ...
// func OpenList(description, channelID string) string {
// 	query := `INSERT cart SET description = ?, status = ?, channel_id = ?`
// 	stmt, err := app.Connection.Prepare(query)
// 	if err != nil {
// 		fmt.Println("Model OpenList [prepare]: ", err)
// 	}

// 	res, err := stmt.Exec(description, 1, channelID)
// 	if err != nil {
// 		fmt.Println("Model OpenList [exec]: ", err)
// 	}

// 	id, _ := res.LastInsertId()
// 	idToString := strconv.FormatInt(int64(id), 10)

// 	return idToString
// }

// // CloseList ...
// func CloseList(channelID string) bool {

// 	query := `UPDATE cart SET status = ? WHERE status = ? AND channel_id = ?`
// 	stmt, err := app.Connection.Prepare(query)
// 	if err != nil {
// 		fmt.Println("Model CloseList [prepare]: ", err)
// 		return false
// 	}

// 	_, err = stmt.Exec(0, 1, channelID)
// 	if err != nil {
// 		fmt.Println("Model CloseList [exec]: ", err)
// 		return false
// 	}

// 	return true
// }
