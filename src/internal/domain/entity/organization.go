package entity

type Organization struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerId   string    `json:"owner_id"`
	Timestamp Timestamp `json:"timestamp"`
}
