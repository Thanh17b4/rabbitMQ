package repo_test

import (
	"errors"
	"testing"

	"github.com/Thanh17b4/practice/repo"
	"gotest.tools/assert"

	"github.com/Thanh17b4/practice/model"
	dbtest "github.com/Thanh17b4/practice/tests/docker_test"
)

var (
	util *dbtest.TestUtil

	userID = 3
	user   = model.User{
		ID:       userID,
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "abc-xyz",
		Username: "abc-xyz",
	}
)

func init() {
	dbTest := dbtest.New()
	if err := dbTest.InitDB(); err != nil {
		dbTest.Log.Panicf("testutil.initDB(): %v", err)
	}
	util = dbTest
}

func TestUser_Create(t *testing.T) {
	if err := util.SetupDB(); err != nil {
		util.Log.Panicf("util.SetupDB(): %v", err)
	}
	defer util.CleanAndClose()

	tests := []struct {
		name    string
		user    model.User
		wantErr error
		want    *model.User
	}{
		{
			name: "error: wrong user id",
			user: model.User{
				ID:       -1,
				Name:     "Test",
				Email:    "test@gmail.com",
				Password: "abc-xyz",
				Username: "abc-xyz",
			},
			wantErr: errors.New(`gsd`),
			want:    nil,
		},
		{
			name: "error: when duplicate email",
			user: model.User{
				ID:       2,
				Email:    "test@gmail.com",
				Password: "abc-xyz",
				Username: "abc-xyz",
			},
			wantErr: errors.New(`sgsdkkh"`),
			want:    nil,
		},
		{
			name: "success",
			user: model.User{
				ID:       2,
				Email:    "test2@gmail.com",
				Password: "abc-xyz",
				Username: "abc-xyz",
			},
			want: &user,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(util.DB)
			actual, err := userRepo.Create(&tc.user)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("taskRepo.Create got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}
