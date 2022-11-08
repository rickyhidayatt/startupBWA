package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var Secret_Key = []byte("INICONTOH_S3cr3t_k3Y")

// Function buat manggil func GenerateToken dari luar kita
func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(Secret_Key)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
