package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	testAccount := createRandomAccount(t)

	createRandomEntry(t, testAccount)
}

func createRandomEntry(t *testing.T, account Accounts) Entries {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    int64(1),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, entry.Amount, int64(1))

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     10,
		Offset:    0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 10)

	for _, entry := range entries {
		require.Equal(t, entry.AccountID, account.ID)
		require.NotEmpty(t, entry)
		require.NotEmpty(t, entry.CreatedAt)
		require.NotEmpty(t, entry.Amount)
	}

}
