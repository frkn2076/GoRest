package Models

import (
	"GoRest/Config"
	_ "fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

//GetAllCategories Fetch all category data
func GetAllCategories(category *[]Category) (err error) {
	if err = Config.DB.Find(category).Error; err != nil {
		return err
	}
	return nil
}

//GetCategoryByID ... Fetch only one user by Id
func GetCategoryByID(category *Category, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(category).Error; err != nil {
		return err
	}
	return nil
}
