package modal

type CreateCategoryDto struct {
	Name           string `json:"Name"`
	ParentCategory int64  `json:"ParentCategory"`
}
