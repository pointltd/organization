package entity

type Point struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	OrganizationId string    `json:"organization_id"`
	Timestamp      Timestamp `json:"timestamp"`
}
