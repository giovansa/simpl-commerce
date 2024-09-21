package product

import (
	"database/sql"
	"time"
)

type Product struct {
	ProductID      string       `json:"product_id" db:"product_id"`
	Name           string       `json:"name" db:"name"`
	Description    string       `json:"description" db:"description"`
	AvailableStock int64        `json:"available_stock" db:"available_stock"`
	LockStock      int64        `json:"lock_stock" db:"lock_stock"`
	PendingStock   int64        `json:"pending_stock" db:"pending_Stock"`
	Sold           int64        `json:"sold" db:"sold_qty"`
	ShopID         string       `json:"shop_id" db:"shop_id"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at" db:"updated_at"`
}

func (p *Product) toDTO() Item {
	return Item{
		ProductID:      p.ProductID,
		Name:           p.Name,
		Description:    p.Description,
		AvailableStock: p.AvailableStock,
	}
}

type ListProduct []Product

func (lp *ListProduct) toDTO() ListProductResponse {
	result := make([]Item, len(*lp))
	for i := 0; i < len(*lp); i++ {
		result[i] = Item{
			ProductID:      (*lp)[i].ProductID,
			Name:           (*lp)[i].Name,
			Description:    (*lp)[i].Description,
			AvailableStock: (*lp)[i].AvailableStock,
		}
	}
	return ListProductResponse{ProductList: result}
}
