package product

type Item struct {
	ProductID      string
	Name           string
	Description    string
	AvailableStock int64
}

type ListProductResponse struct {
	ProductList []Item `json:"product_list"`
}
