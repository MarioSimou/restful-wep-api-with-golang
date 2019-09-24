package models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
}
type Users []User

type WrappedResponseSuccessOne struct {
	Status  int16 `json:"status"`
	Success bool  `json:"success"`
	Data    User  `json:"data"`
}

type WrappedResponseSuccessMany struct {
	Status  int16  `json:"status"`
	Success bool   `json:"success"`
	Data    []User `json:"data"`
}

type WrappedResponseError struct {
	Status  int16  `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
