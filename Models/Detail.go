package Models

import (
	"GoRest/Config"

	_ "github.com/denisenkom/go-mssqldb"
)

//GetAllFoods Fetch all user data
func GetDetailByFoodId(food *[]Food, categoryId string) (err error) {
	if err = Config.DB.Where("foodId = ?", categoryId).Find(food).Error; err != nil {
		return err
	}
	return nil
}
