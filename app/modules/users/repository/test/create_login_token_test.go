package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	"github.com/vkhoa145/facebook-mini-api/app/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateLogin(t *testing.T) {
	type args struct {
		jwt    *models.JwtResponse
		UserID uint
		tx     *transaction.TransactionManager
	}

	// mockedTime := time.Now()
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

	transaction := transaction.NewTransactionManager(db)

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       *models.JwtResponse
		wantErr    bool
	}{
		{
			name: "success create login token",
			args: args{
				jwt: &models.JwtResponse{
					RefreshToken: "akjdkfjaskdjf",
				},
				UserID: 1,
				tx:     transaction,
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "login_tokens" ("user_id","refresh_token") VALUES ($1$,$2$)`,
				)).WithArgs(1, "asdfasdf").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mockSql.ExpectCommit()
			},
			want: &models.JwtResponse{RefreshToken: "akjdkfjaskdjf", AccessToken: "adsfasdfsf"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			u.DB.Transaction(func(tx *gorm.DB) error {
				got, err := u.CreateLoginToken(tt.args.jwt, tt.args.UserID, tx)
				t.Logf("CreateLoginTokengot = %v, err: %v", got, err)

				if (err != nil) != tt.wantErr {
					t.Errorf("CreateLoginToken error = %v, wantErr %v", err, tt.wantErr)
					return err
				}

				return nil
			})

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			} else {
				t.Logf("Fullfilled Expectations")
			}
		})
	}
}
