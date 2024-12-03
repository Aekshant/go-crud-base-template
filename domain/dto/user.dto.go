package dto

type GetUserDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
