package models

type IdRequest struct {
	ID string `json:"id" binding:"required"`
}
