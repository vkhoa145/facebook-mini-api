package usecase_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/modules/users/repository/mocks"
	"github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
	u := usecase.UserUseCase{
		UserRepo: mockUserRepo,
	}

	type args struct {
		user *models.User
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *mocks.MockUserRepoInterface)
		want       *models.UserResponse
		wantErr    bool
	}{
		{
			name: "success sign up",
			args: args{
				user: &models.User{
					Email:    "khoa@gmail.com",
					Name:     "khoa",
					Birthday: "10/10/1991",
					Password: "12345678",
				},
			},
			beforeTest: func(userRepo *mocks.MockUserRepoInterface) {
				userRepo.EXPECT().CheckExistedEmail("khoa@gmail.com").Return(false)
				userRepo.EXPECT().CreateUser(
					models.User{
						Email:    "khoa@gmail.com",
						Name:     "khoa",
						Birthday: "10/10/1991",
						Password: "12345678",
					},
					u.Tx.Begin(),
				).Return(
					models.User{
						Email:    "khoa@gmail.com",
						Name:     "khoa",
						Birthday: "10/10/1991",
						Password: "12345678",
					}, nil,
				)
				userRepo.EXPECT().CreateLoginToken(
					&models.LoginToken{
						UserID: 1,
					},
					u.Tx.Begin(),
				).Return(
					&models.LoginToken{
						UserID: 1,
					}, nil,
				)
			},
			want: &models.UserResponse{
				ID:           1,
				Email:        "khoa@gmail.com",
				Birthday:     "10/10/1991",
				Name:         "khoa",
				AccessToken:  "adkfjlasdjflaljdf",
				RefreshToken: "laekdlfkldsfllk",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mockUserRepo)
			}

			got, err := u.SignUp(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase SignUp error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase SignUp = %v, want %v", got, tt.want)
			}
		})
	}
}
