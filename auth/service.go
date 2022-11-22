package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidasiToken(token string) (*jwt.Token, error) //kenapa balikin jwt? nanti butuh method dari bawaan jwt.token
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

func (s *jwtService) ValidasiToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(Secret_Key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil

	// validasi terhadap token intinya apakah benar secret key yang dimiliki user sama dengan yang tadi kita kasih sebelumnya jika ya maju jika no stop ini kayanya bisa pake go Validator buat mempersingkat kodingan ini
}
