package Models

type Detail struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Info       string `json:"info"`
	Details	   []Step `json:"info"`
	FoodId 	   uint   `json:"categoryId"`
}

func (b *Detail) TableName() string {
	return "detail"
}
