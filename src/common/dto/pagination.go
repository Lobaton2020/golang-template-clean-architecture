package common

type PaginationDto struct {
	Page  int `json:"page" validate:"required,number"`
	Limit int `json:"limit" validate:"required,number"`
}