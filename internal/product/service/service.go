package service

import "shop-product-service/internal/product/modal"

type ProductStorage interface {
	FindById(id int) modal.Product
}
