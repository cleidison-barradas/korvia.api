package dto

import "github.com/cleidison-barradas/korvia.api/internal/users/domain/model"

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Document string `json:"document" binding:"required"`
	DocumentType model.DocumentType `json:"document_type" binding:"required"`
}