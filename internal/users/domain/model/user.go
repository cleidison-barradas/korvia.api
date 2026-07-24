package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string
type DocumentType string

const (
	CPF DocumentType = "CPF"
	CNPJ DocumentType = "CNPJ"
)

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleOwner UserRole = "owner"
	UserRoleEstablishment UserRole = "establishment"
	UserRoleProfessional UserRole = "professional"
)

type User struct {
	Id uuid.UUID `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Enabled bool `json:"enabled"`
	Role UserRole `json:"role"`
	PasswordHash string `json:"password"`
	Document string `json:"document"`
	DocumentType DocumentType `json:"document_type"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewUser(fistName string, lastName string, email string, phoen string, passwordHash string, document string, documentType DocumentType) *User {
	now := time.Now()

	return &User{
		FirstName: fistName,
		LastName: lastName,
		Email: email,
		Phone: phoen,
		Enabled: false,
		Role: UserRoleProfessional,
		PasswordHash: passwordHash,
		Document: document,
		DocumentType: documentType,
		VerifiedAt: nil,
		CreatedAt: now,
		UpdatedAt: nil,
	}
	
}