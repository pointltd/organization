package entity

type User struct {
	Id        string      `json:"id"`
	Password  string      `json:"password"`
	Contacts  ContactInfo `json:"contacts"`
	Info      UserInfo    `json:"info"`
	Timestamp Timestamp   `json:"timestamp"`
}

type UserInfo struct {
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
}
