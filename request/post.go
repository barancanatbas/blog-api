package request

type PostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostUpdateReq struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostDeleteReq struct {
	ID uint32 `json:"id"`
}
