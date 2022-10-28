package user

type Userformater struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupatiom string `json:"occupatiom"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) Userformater {
	formaterr := Userformater{
		ID:         user.ID,
		Name:       user.Name,
		Occupatiom: user.Occuraption,
		Email:      user.Email,
		Token:      token,
	}
	return formaterr
}
