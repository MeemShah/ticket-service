package dto

type Ticket struct {
	Id          string `json:"ticket-id"      db:"id"`
	Category    string `json:"category"       db:"category"`
	Price       int    `json:"price"          db:"price"`
	PurchasedBy string `json:"purchased_by"   db:"purchased_by"`
	Status      string `json:"status"         db:"status"`
	IsActive    bool   `json:"is-active"      db:"is_active"`
}