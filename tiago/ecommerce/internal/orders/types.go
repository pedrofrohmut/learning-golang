package orders

type OrderItem struct {
	ProductId int64 `json:"productId"`
	Quantity int32 `json:"quantity"`
}

type CreateOrderParams struct {
	CustomerId int64 `json:"customerId"`
	Items []OrderItem `json:"items"`
}
