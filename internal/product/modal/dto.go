package modal

type CreateProductDTO struct {
	Name        string
	Price       int64
	Count       int64
	Image       []string
	Description string
	CategoryId  int64
}
