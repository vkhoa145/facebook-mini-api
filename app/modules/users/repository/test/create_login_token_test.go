package repository_test

import (
	"errors"
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

func TestCreateLogin(t *testing.T) {
	type args struct {
		loginToken *models.LoginToken
		tx         *transaction.TransactionManager
	}

	mockedTime := time.Now()
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
		want       *models.LoginToken
		wantErr    bool
	}{
		{
			name: "success create login token",
			args: args{
				loginToken: &models.LoginToken{
					RefreshToken: "20aafe8e63f0bfe1d69f5bf131d4173c",
					UserID:       1,
					UpdatedAt:    mockedTime,
					CreatedAt:    mockedTime,
				},
				tx: transaction,
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "login_tokens" ("user_id","refresh_token","created_at","updated_at") VALUES ($1,$2,$3,$4)`,
				)).WithArgs(1, "20aafe8e63f0bfe1d69f5bf131d4173c", mockedTime, mockedTime).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mockSql.ExpectCommit()
			},
			want: &models.LoginToken{UserID: 1, RefreshToken: "20aafe8e63f0bfe1d69f5bf131d4173c", CreatedAt: mockedTime, UpdatedAt: mockedTime},
		},
		{
			name: "duplicate user id",
			args: args{
				loginToken: &models.LoginToken{
					RefreshToken: "20aafe8e63f0bfe1d69f5bf131d4173c",
					UserID:       1,
					UpdatedAt:    mockedTime,
					CreatedAt:    mockedTime,
				},
				tx: transaction,
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "login_tokens" ("user_id","refresh_token","created_at","updated_at") VALUES ($1,$2,$3,$4)`,
				)).WithArgs(1, "20aafe8e63f0bfe1d69f5bf131d4173c", mockedTime, mockedTime).WillReturnError(errors.New("duplicate key value violates unique constraint \"uni_login_tokens_user_id\""))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty refresh token",
			args: args{
				loginToken: &models.LoginToken{
					RefreshToken: "20aafe8e63f0bfe1d69f5bf131d4173c",
					UserID:       1,
					UpdatedAt:    mockedTime,
					CreatedAt:    mockedTime,
				},
				tx: transaction,
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "login_tokens" ("user_id","refresh_token","created_at","updated_at") VALUES ($1,$2,$3,$4)`,
				)).WithArgs(1, "20aafe8e63f0bfe1d69f5bf131d4173c", mockedTime, mockedTime).WillReturnError(errors.New("new row for relation \"login_tokens\" violates check constraint \"refresh_token_not_empty\""))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			u.DB.Transaction(func(tx *gorm.DB) error {
				got, err := u.CreateLoginToken(tt.args.loginToken, tx)
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
