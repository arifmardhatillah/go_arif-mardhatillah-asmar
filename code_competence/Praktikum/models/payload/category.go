package payload

type CreateCategory struct {
	NameCategory string `json:"name" form:"name" validate:"required"`
}
