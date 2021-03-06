package request

type UserLogin struct {
	UserName string `validate:"required" json:"username"`
	Password string `validate:"required"`
}
type UserDelRequest struct {
	ID uint `validate:"required"`
}

type UserInsert struct {
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required" json:"password"`
	Surname  string `validate:"required" json:"surname"`
	Age      uint   `validate:"required" json:"age"`
	Job      string `validate:"required" json:"job"`
	UserName string `validate:"required" json:"username"`
}
