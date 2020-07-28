package Models

type Category struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Info  string `json:"info"`
}

func (b *Category) TableName() string {
	return "category"
}
