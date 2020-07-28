package Models

import (
	"GoRest/Config"

	_ "github.com/denisenkom/go-mssqldb"
)

//GetAllFoods Fetch all user data
func GetAllFoodsByCategoryId(food *[]Food, categoryId string) (err error) {
	if err = Config.DB.Where("categoryId = ?", categoryId).Find(food).Error; err != nil {
		return err
	}
	return nil
}
