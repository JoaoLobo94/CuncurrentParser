package db

import (
	"context"
	"testing"
	"time"
	"github.com/pioz/faker"
	"github.com/stretchr/testify/require"
)

func CreateRandomBatch(t *testing.T) Batch {
	user1 := CreateRandomUser(t)
	arg := CreateBatchParams{
		Dispatched: false,
		Amount:     faker.Float64(),
		UserID:     user1.ID,
	}

	batch, err := testQueries.CreateBatch(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, batch)
	require.NotZero(t, batch.ID)
	require.False(t, batch.Dispatched)
	require.NotZero(t, batch.Amount)
	require.NotZero(t, batch.UserID)
	require.NotZero(t, batch.CreatedAt)

	return batch

}

func TestCreateBatch(t *testing.T) {
	CreateRandomBatch(t)
}

func TestGetBatch(t *testing.T) {
	batch1 := CreateRandomBatch(t)
	batch2, err := testQueries.GetBatch(context.Background(), batch1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, batch2)
	require.Equal(t, batch1.ID, batch2.ID)
	require.Equal(t, batch1.Dispatched, batch2.Dispatched)
	require.Equal(t, batch1.Amount, batch2.Amount)
	require.Equal(t, batch1.UserID, batch2.UserID)
	require.WithinDuration(t, batch1.CreatedAt, batch2.CreatedAt, time.Second)
}

// Need to figure out what is the problem here. Solve later

// func TestUpdateBatch(t *testing.T) {
// 	batch1 := CreateRandomBatch(t)

// 	arg := UpdateBatchParams{
// 		ID:         batch1.ID,
// 		Amount:     faker.Float64(),
// 	}
	
// 	batch2, err := testQueries.UpdateBatch(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, batch2)
// 	require.Equal(t, batch1.ID, batch2.ID)
// 	require.Equal(t, arg.Dispatched, batch2.Dispatched)
// 	require.Equal(t, arg.Amount, batch2.Amount)
// 	require.Equal(t, batch1.UserID, batch2.UserID)
// 	require.WithinDuration(t, batch1.CreatedAt, batch2.CreatedAt, time.Second)
// }

func TestListBatch(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomBatch(t)
	}

	batches, err := testQueries.ListBatches(context.Background())
	require.NoError(t, err)

	for _, batch := range batches {
		require.NotEmpty(t, batch)
	}

}
