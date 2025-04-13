package dto

type Ticket struct {
	Id       string `json:"ticket-id"`
	Catagory string `json:"catagory"`
	Price    int    `json:"price"`
	Status   string `json:"status"`
	IsActive bool   `json:"is-active"`
}
