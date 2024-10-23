package repository_test

import (
	"errors"
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
	type args struct {
		user *models.User
		tx   *transaction.TransactionManager
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
		want       *models.User
		wantErr    bool
	}{
		{
			name: "success create user",
			args: args{
				user: &models.User{
					Name:      "Khoa",
					Email:     "test22@gmail.com",
					Password:  "12345678",
					Birthday:  "10/12/1991",
					Phone:     "12345678",
					CreatedAt: mockedTime,
					UpdatedAt: mockedTime,
				},
				tx: transaction,
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "users" ("name","email","password","birthday","phone","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`,
				)).WithArgs("Khoa", "test22@gmail.com", "12345678", "10/12/1991", "12345678", mockedTime, mockedTime).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mockSql.ExpectCommit()
			},
			want: &models.User{ID: 1, Name: "Khoa", Email: "test22@gmail.com", Password: "12345678", Birthday: "10/12/1991", Phone: "12345678", CreatedAt: mockedTime, UpdatedAt: mockedTime},
		},
		{
			name: "Duplicate email create user",
			args: args{
				user: &models.User{
					Name:      "Khoa",
					Email:     "test22@gmail.com",
					Password:  "12345678",
					Birthday:  "10/12/1991",
					Phone:     "12345678",
					CreatedAt: mockedTime,
					UpdatedAt: mockedTime,
				},
				tx: transaction,
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "users" ("name","email","password","birthday","phone","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`,
				)).WithArgs("Khoa", "test22@gmail.com", "12345678", "10/12/1991", "12345678", mockedTime, mockedTime).WillReturnError(errors.New("duplicate key value violates unique constraint \"users_email_key\""))

				mockSql.ExpectRollback()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty email create user",
			args: args{
				user: &models.User{
					Name:      "Khoa",
					Email:     "khoa123@gmail.com",
					Password:  "12345678",
					Birthday:  "10/12/1991",
					Phone:     "12345678",
					CreatedAt: mockedTime,
					UpdatedAt: mockedTime,
				},
				tx: transaction,
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectBegin()
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "users" ("name","email","password","birthday","phone","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`,
				)).WithArgs("Khoa", "khoa123@gmail.com", "12345678", "10/12/1991", "12345678", mockedTime, mockedTime).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mockSql.ExpectCommit()
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

			got, err := u.CreateUser(tt.args.user, tt.args.tx.Begin())
			t.Logf("UserRepo.Create() got1111 = %v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Create() got = %v\n", got)
				t.Errorf("UserRepo.Create() error = %v\n, wantErr: %v\n", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
