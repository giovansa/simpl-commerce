package product

type FetchProductReq struct {
	Name   string `db:"name"`
	Limit  int    `db:"limit"`
	Offset int    `db:"offset"`
}

type NewLedger struct {
	RefID         string `db:"ref_id"`
	ProductID     string `db:"product_id"`
	StockMovement string `db:"stock_movement"`
	ActionType    string `db:"action_type"`
	UserID        string `db:"user_id"`
}
