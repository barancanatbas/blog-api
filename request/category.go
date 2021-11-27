package request

type CategorySave struct {
	Name string `json:"name"`
}

type CategoryDelete struct {
	Id uint `json:"id"`
}

type CategoryUpdate struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
