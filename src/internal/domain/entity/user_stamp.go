package entity

type UserStamp struct {
	CreatedById *string `json:"created_by_id"`
	UpdatedById *string `json:"updated_by_id"`
	DeletedById *string `json:"deleted_by_id"`
}
