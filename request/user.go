package request

type UserLogin struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}
