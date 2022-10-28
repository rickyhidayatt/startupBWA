package user

// Ini buat mapping halaman login yang di butuhin halaman loggin ini disini
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" bindng:"required"`
}

// ini buat halaman login user
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" bindng:"required"`
}
