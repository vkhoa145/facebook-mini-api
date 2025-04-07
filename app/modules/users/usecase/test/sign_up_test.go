package usecase_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/modules/users/repository/mocks"
	"github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
	"github.com/vkhoa145/facebook-mini-api/app/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()

	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})

	transaction := transaction.NewTransactionManager(db)
	u := usecase.UserUseCase{
		UserRepo: mockUserRepo,
		Tx:       *transaction,
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
				userRepo.EXPECT().CheckExistedEmail(gomock.Eq("khoa@gmail.com")).Return(false).Times(1)
				userRepo.EXPECT().CreateUser(
					&models.User{
						Email:    "khoa@gmail.com",
						Name:     "khoa",
						Birthday: "10/10/1991",
						Password: "12345678",
					},
					gomock.Any(),
				).Return(
					&models.User{
						ID:       1,
						Email:    "khoa@gmail.com",
						Name:     "khoa",
						Birthday: "10/10/1991",
						Password: "12345678",
					}, nil,
				)
				userRepo.EXPECT().CreateLoginToken(
					gomock.AssignableToTypeOf(&models.LoginToken{}), // Chỉ cần đúng kiểu dữ liệu
					gomock.Any(),
				).Return(
					&models.LoginToken{
						UserID:       1,
						RefreshToken: gomock.Any().String(),
					}, nil,
				)
			},
			want: &models.UserResponse{
				ID:           1,
				Email:        "khoa@gmail.com",
				Birthday:     "10/10/1991",
				Name:         "khoa",
				AccessToken:  gomock.Any().String(),
				RefreshToken: gomock.Any().String(),
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

			if got.ID != tt.want.ID || got.Email != tt.want.Email || got.Name != tt.want.Name || got.Birthday != tt.want.Birthday {
				t.Errorf("UseCase SignUp = %v, want %v", got, tt.want)
			}
		})
	}
}
