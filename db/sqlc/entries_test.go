package db

import (
	"bank/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createEntry(t *testing.T) Entries {

	account := createRandomAccount(t)

	args := CreateEntriesParams{
		AccountID: sql.NullInt64{Int64: account.ID, Valid: true},
		Money:     util.RandomMoney(),
	}

	entry, err := testQuries.CreateEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.Money, args.Money)
	require.Equal(t, entry.AccountID, args.AccountID)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createEntry(t)
}

func TestGetEntry(t *testing.T) {

	oldEntry := createEntry(t)

	entry, err := testQuries.GetEntry(context.Background(), int32(oldEntry.ID))
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, oldEntry.AccountID)
	require.Equal(t, entry.ID, oldEntry.ID)
	require.Equal(t, entry.Money, oldEntry.Money)
	require.WithinDuration(t, entry.CreatedAt, oldEntry.CreatedAt, time.Second)

}
