package userdomain

type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	HashedPassword string `json:"-"`
	PhoneNumber    string `json:"phone_number"`
	Role           Role   `json:"role"`
}
