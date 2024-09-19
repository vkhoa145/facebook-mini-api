package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
)

func (r UserRepo) CheckExistedEmail(email string) bool {
	var user models.User
	result := r.DB.Table(user.TableName()).Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return false
	}

	return result.RowsAffected > 0
}
