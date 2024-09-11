package entity

type User struct {
	UUID      string    `json:"uuid"`
	Info      UserInfo  `json:"info"`
	UserStamp UserStamp `json:"user_stamp"`
	Timestamp Timestamp `json:"timestamp"`
}

type UserInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
