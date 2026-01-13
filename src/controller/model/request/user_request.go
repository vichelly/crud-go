package request

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int8   `json:"age"`
	Password string `json:"password"`
}
