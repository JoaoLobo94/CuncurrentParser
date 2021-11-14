package db

import (
	"context"
	"testing"
	"time"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	name := faker.Username()
	user, err := testQueries.CreateUser(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
