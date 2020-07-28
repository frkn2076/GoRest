package Models

type Step struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
}

func (b *Step) TableName() string {
	return "step"
}
