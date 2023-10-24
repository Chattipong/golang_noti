package dto

type CreateUser struct {
	FirstName string ` json:"firstName"  binding:"required"`
	LastName  string ` json:"lastName"  binding:"required"`
	Email     string `json:"email"  binding:"required,email"`
	Phone     string ` json:"phone" binding:"required"`
	Password  string ` json:"password"  binding:"required"`
	Role      int    ` json:"role"  binding:"required" min:0 max:100`
	Status    bool   ` json:"status"  binding:"required"`
}

type UpdateUser struct {
	FirstName string ` json:"firstName"  binding:"required"`
	LastName  string ` json:"lastName"  binding:"required"`
	Email     string `json:"email"  binding:"required,email"`
	Phone     string ` json:"phone" binding:"required"`
}
