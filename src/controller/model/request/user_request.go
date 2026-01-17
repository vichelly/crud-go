package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Age      int8   `json:"age" binding:"required"`
	Password string `json:"password" binding:"required,containsany=!@#$%^&*(),min=4,max=20"`
}
