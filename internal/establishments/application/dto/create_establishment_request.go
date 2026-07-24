package dto

type CreateEstablishmentRequest struct {
	Name string `json:"name" binding:"required"`
	WabaId string `json:"waba_id" binding:"required"`
	PhoneNumberId string `json:"phone_number_id" binding:"required"`
}