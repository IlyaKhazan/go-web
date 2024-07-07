package db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func CreateUser(t *testing.T) User {
	generatedUUID, err := uuid.NewUUID()
	arg := CreateUserParams{
		ID:        generatedUUID,
		Firstname: "Olga",
		Lastname:  "Jane",
		Email:     RandomEmail(),
		Age:       31,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Firstname, user.Firstname)
	require.Equal(t, arg.Lastname, user.Lastname)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Age, user.Age)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Firstname, user2.Firstname)
	require.Equal(t, user1.Lastname, user2.Lastname)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Age, user2.Age)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateUser(t)

	arg := UpdateUserParams{
		ID:        user1.ID,
		Firstname: user1.Firstname,
		Lastname:  "Ivano",
		Email:     user1.Email,
		Age:       49,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user1.Firstname, user2.Firstname)
	require.Equal(t, arg.Lastname, user2.Lastname)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.Age, user2.Age)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.Empty(t, user2)
	require.EqualError(t, err, ErrRecordNotFound.Error())
}

//func TestListUsers(t *testing.T) {
//	for i := 0; i < 10; i++ {
//		CreateUser(t)
//	}
//
//	arg := ListUsersParams{
//
//		Offset: 5,
//	}
//
//	users, err := testQueries.ListUsers(context.Background(), arg)
//	require.NoError(t, err)
//	require.Len(t, users, 5)
//
//	for _, user := range users {
//		require.NotEmpty(t, user)
//
//	}
//
//}
