package db

import (
	"context"
	"testing"
	"github.com/pioz/faker"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransaction(t *testing.T) Transaction {
	user1 := CreateRandomUser(t)
	arg := CreateTransactionParams{
		Amount:     faker.Float64(),
		UserID:     user1.ID,
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.Amount)
	require.NotZero(t, transaction.UserID)
	require.NotZero(t, transaction.CreatedAt)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	CreateRandomAction(t)
}

func TestListTransaction(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransaction(t)
	}

	transactions, err := testQueries.ListBatches(context.Background())
	require.NoError(t, err)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
	}

}
