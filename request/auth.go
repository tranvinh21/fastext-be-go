package request

type SignupRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,passwd,min=8,max=20"`
}

type SigninRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,passwd,min=8,max=20"`
}
