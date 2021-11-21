package request

type DenemeReq struct {
	Name string `query:"name" json:"name" validate:"required"`
}
