package types

type User struct {
    ID    string  `json:"id" form:"id"`
    Name  string  `json:"name" form:"name" binding:"required"`
    Email string  `json:"email" form:"email" binding:"required,email"`
}