package db

import (
	"context"
	"testing"
	"time"
	"github.com/pioz/faker"
	"github.com/stretchr/testify/require"
)

func CreateRandomAction(t *testing.T) Action {
	user1 := CreateRandomUser(t)
	arg := CreateActionParams{
		Amount:     faker.Float64(),
		UserID:     user1.ID,
	}

	action, err := testQueries.CreateAction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, action)
	require.NotZero(t, action.ID)
	require.NotZero(t, action.Amount)
	require.NotZero(t, action.UserID)
	require.NotZero(t, action.CreatedAt)

	return action
}

func TestCreateAction(t *testing.T) {
	CreateRandomAction(t)
}

func TestGetAction(t *testing.T) {
	action1 := CreateRandomBatch(t)
	action2, err := testQueries.GetBatch(context.Background(), action1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, action2)
	require.Equal(t, action1.ID, action2.ID)
	require.Equal(t, action1.Dispatched, action2.Dispatched)
	require.Equal(t, action1.Amount, action2.Amount)
	require.Equal(t, action1.UserID, action2.UserID)
	require.WithinDuration(t, action1.CreatedAt, action2.CreatedAt, time.Second)
}


func TestListAction(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAction(t)
	}

	actions, err := testQueries.ListBatches(context.Background())
	require.NoError(t, err)

	for _, action := range actions {
		require.NotEmpty(t, action)
	}

}
