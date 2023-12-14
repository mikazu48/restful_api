package web

type CategoryUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `validate:"required,max=200,min=5" json:"name"`
}
