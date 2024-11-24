package data

type CreateUserDTO struct {
	FirstName            string
	LastName             *string
	Phone                *string
	Password             string
	PasswordConfirmation string
	Email                string
}
