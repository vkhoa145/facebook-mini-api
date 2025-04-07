package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"gorm.io/gorm"
)

func (r UserRepo) CreateVerificationCode(verificationCode *models.VerificationCode, tx *gorm.DB) (*models.VerificationCode, error) {
	createdVerificationCode := tx.Table(models.VerificationCode{}.TableName()).Create(verificationCode)
	if createdVerificationCode.Error != nil {
		return nil, createdVerificationCode.Error
	}

	return verificationCode, nil
}