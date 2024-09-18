package modal

import "time"

type Product struct {
	Uuid        string
	Name        string
	Price       int64
	Count       int64
	Image       []string
	Description string
	CategoryId  int64
}

type Reviews struct {
	Uuid      string
	ProductId string
	Rating    float64
	CreatedAt time.Time
	Comment   string
	Name      string
}
