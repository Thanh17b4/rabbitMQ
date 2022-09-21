package repo_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	db, err := SetupDB()
	fmt.Println(err)
	require.NoError(t, err)
	defer cleanUpDB(db)
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
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestUser_ListUser(t *testing.T) {
	db, err := SetupDB()
	fmt.Println(err)
	require.NoError(t, err)
	defer cleanUpDB(db)
	// no have record
	t.Run("no have record", func(t *testing.T) {
		//db.Exec("INSERT INTO users VALUES (1, 'thanh', 'abc@gmail.com', 0, 0, 0, 'china', '22121992', 'thanh17')")
		listUsers := []*model.User(nil)
		userRepo := repo.NewUser(db)
		got, err := userRepo.ListUser(1, 10)
		want := listUsers
		assert.Nil(t, err)
		assert.Equal(t, want, got)

	})
	//add some records
	t.Run("have some records", func(t *testing.T) {
		db.Exec("INSERT INTO users VALUES (1, 'thanh', 'abc@gmail.com', 0, 0, 0, 'china', '22121992', 'thanh17')")
		db.Exec("INSERT INTO users VALUES (2, 'thanh', 'cdf@gmail.com', 0, 0, 0, 'china', '22121992', 'thanh18')")
		user1 := &model.User{ID: 1, Name: "thanh", Email: "abc@gmail.com", Protected: 0, Banned: 0, Activated: 0, Address: "china", Password: "", Username: "thanh17"}
		user2 := &model.User{ID: 2, Name: "thanh", Email: "cdf@gmail.com", Protected: 0, Banned: 0, Activated: 0, Address: "china", Password: "", Username: "thanh18"}
		var listUsers []*model.User
		listUsers = append(listUsers, user1, user2)
		userRepo := repo.NewUser(db)
		got, err := userRepo.ListUser(1, 10)
		want := listUsers
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})

}

func TestUser_DetailUser(t *testing.T) {
	db, err := SetupDB()
	require.NoError(t, err)
	defer cleanUpDB(db)
	db.Exec("INSERT INTO users VALUES (1, '', '', 0, 0, 0, '', '', '')")

	tests := []struct {
		name    string
		userID  int
		wantErr error
		want    *model.User
	}{
		{
			name:    "error: userID invalid",
			userID:  100,
			wantErr: errors.New(`userID is not correct: sql: no rows in result set`),
			want:    nil,
		},
		{
			name:   "success: insert user successfully",
			userID: 1,
			want: &model.User{
				ID:        1,
				Name:      "",
				Email:     "",
				Protected: 0,
				Banned:    0,
				Activated: 0,
				Address:   "",
				Password:  "",
				Username:  "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(db)
			actual, err := userRepo.DetailUser(int64(tc.userID))
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.Detail got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestUser_UpdateUser(t *testing.T) {
	db, err := SetupDB()
	require.NoError(t, err)
	defer cleanUpDB(db)
	db.Exec("INSERT INTO users VALUES (1, 'thanh', 'abc@gmail.com', 0, 0, 0, 'china', '22121992', 'thanh17')")

	tests := []struct {
		name    string
		user    model.User
		wantErr error
		want    *model.User
	}{
		{
			name: "id invalid",
			user: model.User{
				ID:        4,
				Name:      "thanh",
				Email:     "abc@gmail.com",
				Protected: 0,
				Banned:    0,
				Activated: 0,
				Address:   "china",
				Password:  "22121992",
				Username:  "thanh17",
			},
			wantErr: errors.New(`userID is not correct: sql: no rows in result set`),
			want:    nil,
		},

		{
			name: "success: insert user successfully",
			user: model.User{
				ID:        1,
				Name:      "thanh",
				Email:     "abc@gmail.com",
				Protected: 0,
				Banned:    0,
				Activated: 0,
				Address:   "china",
				Password:  "221293",
				Username:  "thanh17b4",
			},
			want: &model.User{
				ID:        1,
				Name:      "thanh",
				Email:     "abc@gmail.com",
				Protected: 0,
				Banned:    0,
				Activated: 0,
				Address:   "china",
				Password:  "221293",
				Username:  "thanh17b4",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(db)
			actual, err := userRepo.UpdateUser(&tc.user)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.Update got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestUser_Delete(t *testing.T) {
	db, err := SetupDB()
	require.NoError(t, err)
	defer cleanUpDB(db)
	db.Exec("INSERT INTO users VALUES (1, 'Test', 'test@gmail.com', 0, 0, 0, '', 'abc-test', 'abc-kka')")

	tests := []struct {
		name    string
		userID  int64
		wantErr error
		want    int64
	}{
		{
			name:    "error: userID invalid",
			userID:  100,
			wantErr: errors.New(`userID is not correct: sql: no rows in result set`),
			want:    0,
		},
		{
			name:   "success: delete user successfully",
			userID: 1,
			want:   int64(1),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(db)
			actual, err := userRepo.Delete(tc.userID)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.Detail got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)

		})
	}
}

func TestUser_GetUserByEmail(t *testing.T) {
	db, err := SetupDB()
	require.NoError(t, err)
	defer cleanUpDB(db)
	db.Exec("INSERT INTO users VALUES (1, 'thanh', 'abc@gmail.com', 0, 0, 0, '', '', '')")

	tests := []struct {
		name    string
		email   string
		wantErr error
		want    *model.User
	}{
		{
			name:    "error: email invalid",
			email:   "cdf@gmail.com",
			wantErr: errors.New(`email is not correct, please try again: sql: no rows in result set`),
			want:    nil,
		},
		{
			name:  "success: get detail user successfully",
			email: "abc@gmail.com",
			want: &model.User{
				ID:        1,
				Name:      "thanh",
				Email:     "abc@gmail.com",
				Protected: 0,
				Banned:    0,
				Activated: 0,
				Address:   "",
				Password:  "",
				Username:  "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(db)
			actual, err := userRepo.GetUserByEmail(tc.email)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.DetailByEmail got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestUser_GetUserByUsername(t *testing.T) {
	db, err := SetupDB()
	require.NoError(t, err)
	defer cleanUpDB(db)
	db.Exec("INSERT INTO users VALUES (1, 'thanh', 'abc@gmail.com', 0, 0, 0, '', '', 'thanh17')")

	tests := []struct {
		name     string
		username string
		wantErr  error
		want     *model.User
	}{
		{
			name:     "error: username invalid",
			username: "thanh18",
			wantErr:  errors.New(`username is not correct, please try again: sql: no rows in result set`),
			want:     nil,
		},
		{
			name:     "success: get detail user successfully",
			username: "thanh17",
			want: &model.User{
				ID:        1,
				Name:      "thanh",
				Email:     "abc@gmail.com",
				Protected: 0,
				Banned:    0,
				Activated: 0,
				Address:   "",
				Password:  "",
				Username:  "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(db)
			actual, err := userRepo.GetUserByUsername(tc.username)
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.DetailByUsername got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestUser_CountUsers(t *testing.T) {
	db, err := SetupDB()
	require.NoError(t, err)
	defer cleanUpDB(db)
	db.Exec("INSERT INTO users VALUES (1, 'thanh', 'abc@gmail.com', 0, 0, 0, '', '', 'thanh17')")

	tests := []struct {
		name    string
		wantErr error
		want    int64
	}{
		{
			name:    "error: could not count users",
			wantErr: errors.New(`could not count users`),
			want:    1,
		},
		{
			name:    "success: count user successfully",
			wantErr: nil,
			want:    1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := repo.NewUser(db)
			actual, err := userRepo.CountUsers()
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.Count got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}
