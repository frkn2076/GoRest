package Models

type Food struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Info       string `json:"info"`
	CategoryId uint   `json:"categoryId"`
}

func (b *Food) TableName() string {
	return "food"
}
