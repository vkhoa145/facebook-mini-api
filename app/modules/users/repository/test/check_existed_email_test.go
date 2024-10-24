package repository_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCheckExistedEmail(t *testing.T) {
	type args struct {
		email string
	}

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

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       bool
		wantErr    bool
	}{
		{
			name: "email existed",
			args: args{
				email: "khoa@gmail.com",
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE email = $1`,
				)).WithArgs("khoa@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow(1, "khoa@gmail.com"))
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "email not existed",
			args: args{
				email: "khoa@gmail.com",
			},
			beforeTest: func(mockSql sqlmock.Sqlmock) {
				mockSql.MatchExpectationsInOrder(true)
				mockSql.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE email = $1`,
				)).WithArgs("khoa@gmail.com").WillReturnRows(sqlmock.NewRows([]string{}))
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.CheckExistedEmail(tt.args.email)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.Create() = %v, want %v", got, tt.want)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			} else {
				t.Logf("Fullfilled Expectations")
			}
		})
	}
}
