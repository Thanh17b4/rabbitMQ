package repo_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Thanh17b4/practice/model"
	"github.com/Thanh17b4/practice/repo"
)

var (
	user = model.User{
		ID:       3,
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "abc-xyz",
		Username: "abc-xyz",
	}
)

func TestUser_Create(t *testing.T) {
	resource, pool, db := SetupDB()
	defer closeContainer(resource, pool)

	tests := []struct {
		name    string
		user    model.User
		wantErr error
		want    *model.User
	}{
		{
			name: "success: insert user successfully",
			user: user,
			want: &model.User{
				ID:       3,
				Name:     "Test",
				Email:    "test@gmail.com",
				Password: "abc-xyz",
				Username: "abc-xyz",
			},
		},
		{
			name: "error: email userID",
			user: model.User{
				ID:       4,
				Name:     "Test",
				Email:    "test@gmail.com",
				Password: "abc-test",
				Username: "abc-kka",
			},
			wantErr: errors.New(`could not creat new user: Error 1062: Duplicate entry 'test@gmail.com' for key 'users_email_unique'`),
			want:    nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(db)
			actual, err := userRepo.Create(&tc.user)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.Create got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			if actual != nil {
				assert.Equal(t, tc.want, actual)
			}
		})
	}
}
