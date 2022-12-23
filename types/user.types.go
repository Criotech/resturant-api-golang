package types

type LoginRequest struct {
	Password *string `json:"Password" validate:"required"`
	Email    *string `json:"email" validate:"email,required"`
}
