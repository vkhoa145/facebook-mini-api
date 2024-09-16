package repository

import (
	"fmt"

	"github.com/vkhoa145/facebook-mini-api/app/models"
)

func (r UserRepo) CheckExistedEmail(email string) bool {
	var user models.User
	result := r.DB.Table(user.TableName()).Where("email = ?", email).First(&user)

	if result.Error == nil {
		fmt.Println("Error while querying:", result.Error)
		return false
	}

	return result.RowsAffected > 0
}
