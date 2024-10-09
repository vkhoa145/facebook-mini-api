package repository

import "github.com/vkhoa145/facebook-mini-api/app/models"

func (r UserRepo) CreateLoginToken(jwt *models.JwtResponse) (*models.JwtResponse, error) {
	LoginToken := &models.LoginToken{
		RefreshToken: jwt.RefreshToken,
	}
	createRefreshToken := r.DB.Table(models.LoginToken{}.TableName()).Create(LoginToken)
	if createRefreshToken.Error != nil {
		return nil, createRefreshToken.Error
	}

	return jwt, nil
}
