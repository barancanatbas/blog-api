package request

type PostReq struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryFK uint   `json:"categoryfk" validate:"required"`
}

type PostUpdateReq struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostDeleteReq struct {
	ID uint32 `json:"id"`
}

type PostSearchReq struct {
	Key string `param:"key"`
}
