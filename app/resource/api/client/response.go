package client

type RegisterClientResponse struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}
