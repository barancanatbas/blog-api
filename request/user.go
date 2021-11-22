package request

type UserLogin struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}
type UserDelRequest struct {
	ID uint `validate:"required"`
}
