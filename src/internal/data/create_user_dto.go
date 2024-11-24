package data

type CreateUserDTO struct {
	FirstName            string
	LastName             *string
	Password             string
	PasswordConfirmation string
	Email                string
}
