package model

type Product struct {
	ProductID    *int `json:"productId"`
	ProductName  string `json:"productName"`
	ProductPrice string `json:"productPrice"`
}
