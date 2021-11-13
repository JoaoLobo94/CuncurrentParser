package db

import (
	"context"
	"testing"
	"database/sql"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T){
	var name = sql.NullString{String: "joao", Valid: true}
	user, err := testQueries.CreateUser(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}