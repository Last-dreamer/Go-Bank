package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all func to exec db queries & transaction
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore create new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx execute a func within db transactions
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%v, rb err:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// transactiontxParams  the input params ....
type TransactiontxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// transactionTxResult , is the result of transfer transaction ...
type TransferTxResult struct {
	Transfers   Transfers `json:"transfer"`
	FromAccount Account   `json:"from_account"`
	ToAccount   Account   `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, args TransactiontxParams) (TransferTxResult, error) {

	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error
		result.Transfers, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: sql.NullInt64{Int64: args.FromAccountId, Valid: true},
			ToAccountID:   sql.NullInt64{Int64: args.ToAccountId, Valid: true},
			Money:         args.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: sql.NullInt64{Int64: args.FromAccountId, Valid: true},
			Money:     args.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: sql.NullInt64{Int64: args.ToAccountId, Valid: true},
			Money:     args.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err

}
