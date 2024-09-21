package product

type InventoryLedger struct {
	ID            string `json:"id" db:"id"`
	RefID         string `json:"ref_id" db:"ref_id"`
	ProductID     string `json:"product_id" db:"product_id"`
	StockMovement string `json:"stock_movement" db:"stock_movement"`
	ActionType    string `json:"action_type" db:"action_type"`
	UserID        string `json:"user_id" db:"user_id"`
	CreatedAt     string `json:"created_at" db:"created_at"`
}
