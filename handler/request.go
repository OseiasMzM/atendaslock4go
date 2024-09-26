package handler

type CreateClientConnect struct {
	ID       uint   `json:"id"`
	Computer string `json:"computer"`
	User     string `json:"user"`
}
