package entity

type User struct {
	ID        string      `json:"id"`
	Password  string      `json:"password"`
	Contacts  ContactInfo `json:"contacts"`
	Info      UserInfo    `json:"info"`
	UserStamp UserStamp   `json:"user_stamp"`
	Timestamp Timestamp   `json:"timestamp"`
}

type UserInfo struct {
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
}
