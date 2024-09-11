package entity

type UserStamp struct {
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
}
