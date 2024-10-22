package repository_test

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	"github.com/vkhoa145/facebook-mini-api/app/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		user *models.User
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       models.User
		wantErr    bool
	}{
		{
			name: "success create user",
			args: args{
				user: &models.User{
					Name:     "Khoa",
					Email:    "test22@gmail.com",
					Birthday: "10/12/1991",
					Password: "12345678",
				},
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectExec(regexp.QuoteMeta(
					`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email","password","birthday","phone") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
				)).WithArgs(time.Now().Format("2006-01-02 15:04:05.000000"), time.Now().Format("2006-01-02 15:04:05.000000"), nil, "Khoa", "test22@gmail.com", "12345678", "10/12/1991", "").WillReturnResult(sqlmock.NewResult(1, 1))

				mockSql.ExpectCommit()
			},
			want: models.User{Name: "Khoa", Email: "test22@gmail.com", Birthday: "10/12/1991"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			dialector := postgres.New(postgres.Config{
				Conn:       mockDB,
				DriverName: "postgres",
			})
			db, _ := gorm.Open(dialector, &gorm.Config{})
			u := &repository.UserRepo{
				DB: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}
			transaction := transaction.NewTransactionManager(db)

			got, err := u.CreateUser(tt.args.user, transaction.Begin())
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Create() error = %v, wantErr: %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
