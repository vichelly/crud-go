package request

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int8   `json:"age" binding:"required,gte=0,lte=130"`
	Password string `json:"password" binding:"required,min=6"`
}
