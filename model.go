package main

// EXAMPLE FUNCTION
// func getItems() *[]item_struct {
// 	rows, err := db.Query("select item_id, title, artist, release_date, price, image, quality_name, format_name, category_name from inventory as i join qualities as q on q.quality_id = i.quality_id join formats as f on f.format_id = i.format_id join categories as c on c.category_id = i.category_id")

// 	checkError(err, "Query failed")

// 	items := []item_struct{}
// 	for rows.Next() {
// 		var r item_struct
// 		err = rows.Scan(&r.Item_id, &r.Title, &r.Artist, &r.Release_date, &r.Price, &r.Image, &r.Quality, &r.Format, &r.Category)
// 		checkError(err, "No items were found.")
// 		items = append(items, r)
// 	}

// 	return &items
// }
