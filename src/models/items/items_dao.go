package items

const (
	//Basic CRUD queries
	queryInsertItem = "INSERT INTO items (code, title, description, picture, price, " +
		"internal_price, available_quantity, sold_quantity) VALUES (?,?,?,?,?,?,?,?);"
	queryGetItemByCode = "SELECT code, title, description, picture, price, internal_price, available_quantity, sold_quantity FROM items WHERE id=?;"
	queryUpdateItem = "UPDATE items SET title=?, description=?, picture=?, price=?, internal_price=?, available_quantity=?, sold_quantity=? WHERE code=?;"
	queryDeleteItem = "DELETE FROM items WHERE code=?;"
	)

