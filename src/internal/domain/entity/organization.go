package entity

type Organization struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	Timestamp Timestamp `json:"timestamp"`
}
