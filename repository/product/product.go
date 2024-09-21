package product

import (
	"context"
	"database/sql"
	"fmt"
	"simpl-commerce/model/product"

	"github.com/google/uuid"
)

func (r *Repository) FetchProducts(ctx context.Context, input FetchProductReq) ([]product.Product, error) {
	rows, err := r.Db.QueryContext(ctx,
		qGetProduct,
		sql.NullString{
			String: input.Name,
			Valid:  input.Name != "",
		},
		input.Limit,
		input.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()

	var products []product.Product

	for rows.Next() {
		var prod product.Product
		err = rows.Scan(
			&prod.ProductID,
			&prod.Name,
			&prod.Description,
			&prod.AvailableStock,
			&prod.LockStock,
			&prod.PendingStock,
			&prod.ShopID,
			&prod.Sold,
			&prod.CreatedAt,
			&prod.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, prod)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return products, nil
}

func (r *Repository) CreateLedger(ctx context.Context, input NewLedger) error {
	_, err := r.Db.ExecContext(ctx, qInsertLedger,
		uuid.New().String(),
		input.RefID,
		input.ProductID,
		input.StockMovement,
		input.ActionType,
		input.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to create ledger entry: %w", err)
	}

	return nil
}
