package Models

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (b *User) TableName() string {
	return "users"
}
