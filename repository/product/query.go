package product

var (
	qInsertLedger = `
		INSERT INTO product_inventory_ledgers(
			id,
			ref_id,
			product_id,
			stock_movement,
			action_type,
			user_id
		) VALUES ($1, $2, $3, $4, $5, $6);`

	qGetProduct = `
		SELECT
			product_id,
			name,
			description,
			available_stock,
			lock_stock,
			pending_stock,
			shop_id,
			sold_qty,
			created_at,
			updated_at
		FROM products
		WHERE ($1 IS NULL OR name ILIKE '%' || $1 || '%')
		LIMIT $2
		OFFSET $3;`
)
