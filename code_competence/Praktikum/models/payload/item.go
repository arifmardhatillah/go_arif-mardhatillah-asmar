package payload

type CreateItem struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Stock       uint   `json:"stock" form:"stock" validate:"required"`
	Price       uint   `json:"price" form:"price" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
}

type UpdateItem struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Stock       uint   `json:"stock" form:"stock" validate:"required"`
	Price       uint   `json:"price" form:"price" validate:"required"`
}
