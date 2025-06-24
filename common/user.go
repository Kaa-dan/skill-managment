package common

// create user
type UserCreationInput struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

// update user
type UserUpdateInput struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUserUpdateInput() *UserCreationInput {
	return &UserCreationInput{}
}
