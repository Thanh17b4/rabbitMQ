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
	require.NoError(t, err)
	defer cleanUpDB(db)

	// case 1: There is no record
	listUsers := []*model.User(nil)

	t.Run("if no record", func(t *testing.T) {
		userRepo := repo.NewUser(db)
		got, _ := userRepo.ListUser(1, 10)
		want := listUsers

		assert.Equal(t, want, got)
	})
	// case 2: There are some records, add some records to db
	t.Run("if have some records but can not get list users", func(t *testing.T) {
		db.Exec("INSERT INTO users VALUES (1, '', '', 0, 0, 0, '', '', '')")
		//db.Exec("INSERT INTO users VALUES (1, '', '', 0, 0, 0, '', '', '')")

		listUsers = []*model.User{nil}

		userRepo := repo.NewUser(db)
		got, err := userRepo.ListUser(1, 10)
		want := listUsers
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
	t.Run("if have some records", func(t *testing.T) {
		db.Exec("INSERT INTO users VALUES (1, '', '', 0, 0, 0, '', '', '')")
		//db.Exec("INSERT INTO users VALUES (1, '', '', 0, 0, 0, '', '', '')")

		user1 := &model.User{
			ID:        1,
			Name:      "",
			Email:     "",
			Protected: 0,
			Banned:    0,
			Activated: 0,
			Address:   "",
			Password:  "",
			Username:  "",
		}
		//user2 := &model.User{
		//	ID:        2,
		//	Name:      "",
		//	Email:     "",
		//	Protected: 0,
		//	Banned:    0,
		//	Activated: 0,
		//	Address:   "",
		//	Password:  "",
		//	Username:  "",
		//}

		listUsers = []*model.User{}
		listUsers := append(listUsers, user1)

		userRepo := repo.NewUser(db)
		got, _ := userRepo.ListUser(1, 10)
		want := listUsers
		fmt.Println("abc", got)

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
		user    model.User
		wantErr error
		want    *model.User
	}{
		{
			name: "error: userID invalid",
			user: model.User{
				ID:       4,
				Name:     "Test",
				Email:    "test@gmail.com",
				Password: "abc-test",
				Username: "abc-kka",
			},
			wantErr: errors.New(`userID is not correct: sql: no rows in result set`),
			want:    nil,
		},
		{
			name: "success: insert user successfully",
			user: model.User{
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
			actual, err := userRepo.DetailUser(int64(tc.user.ID))
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.Detail got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)
		})
	}
}

//func TestUser_UpdateUser(t *testing.T) {
//	db, err := SetupDB()
//	require.NoError(t, err)
//	defer cleanUpDB(db)
//	//db.Exec("INSERT INTO users VALUES (1, '', '', 0, 0, 0, '', '', 'thanh17')")
//
//	tests := []struct {
//		name    string
//		user    model.User
//		wantErr error
//		want    *model.User
//	}{
//		{
//			name: "id invalid",
//			user: model.User{
//				ID:        4,
//				Name:      "",
//				Email:     "",
//				Protected: 0,
//				Banned:    0,
//				Activated: 0,
//				Address:   "",
//				Password:  "",
//				Username:  "",
//			},
//			wantErr: errors.New(`userID is not correct: sql: no rows in result set`),
//			want:    nil,
//		},
//		//{
//		//	name: "One of the unique field is duplicate",
//		//	user: model.User{
//		//		ID:        1,
//		//		Name:      "",
//		//		Email:     "",
//		//		Protected: 0,
//		//		Banned:    0,
//		//		Activated: 0,
//		//		Address:   "",
//		//		Password:  "",
//		//		Username:  "",
//		//	},
//		//	wantErr: errors.New(`could not update user`),
//		//	want:    nil,
//		//},
//		//{
//		//	name: "success: insert user successfully",
//		//	user: model.User{
//		//		ID:        1,
//		//		Name:      "thanh",
//		//		Email:     "abc@gmail.com",
//		//		Protected: 0,
//		//		Banned:    0,
//		//		Activated: 0,
//		//		Address:   "china",
//		//		Password:  "221293",
//		//		Username:  "thanh17b4",
//		//	},
//		//	want: &model.User{
//		//		ID:        1,
//		//		Name:      "thanh",
//		//		Email:     "abc@gmail.com",
//		//		Protected: 0,
//		//		Banned:    0,
//		//		Activated: 0,
//		//		Address:   "china",
//		//		Password:  "221293",
//		//		Username:  "thanh17b4",
//		//	},
//		//},
//	}
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			userRepo := repo.NewUser(db)
//			actual, err := userRepo.UpdateUser(&tc.user)
//			if err != nil && err.Error() != tc.wantErr.Error() {
//				t.Errorf("userRepo.Update got: %v, but expected: %v", err, tc.wantErr)
//				return
//			}
//			assert.Equal(t, tc.want, actual)
//		})
//	}
//}

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
			if err != nil {
				fmt.Println("eeeeeee: ", err.Error())
			}
			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("userRepo.Detail got: %v, but expected: %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, actual)

		})
	}
}
