package Models

type User struct {
	Id       uint    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (b *User) TableName() string {
	return "users"
}
