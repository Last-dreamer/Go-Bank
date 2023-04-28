// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createEntriesStmt, err = db.PrepareContext(ctx, createEntries); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEntries: %w", err)
	}
	if q.createTransferStmt, err = db.PrepareContext(ctx, createTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTransfer: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteEntryStmt, err = db.PrepareContext(ctx, deleteEntry); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEntry: %w", err)
	}
	if q.deleteTransferStmt, err = db.PrepareContext(ctx, deleteTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTransfer: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getEntriesStmt, err = db.PrepareContext(ctx, getEntries); err != nil {
		return nil, fmt.Errorf("error preparing query GetEntries: %w", err)
	}
	if q.getEntryStmt, err = db.PrepareContext(ctx, getEntry); err != nil {
		return nil, fmt.Errorf("error preparing query GetEntry: %w", err)
	}
	if q.getListAccountStmt, err = db.PrepareContext(ctx, getListAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetListAccount: %w", err)
	}
	if q.getListTransfersStmt, err = db.PrepareContext(ctx, getListTransfers); err != nil {
		return nil, fmt.Errorf("error preparing query GetListTransfers: %w", err)
	}
	if q.getTransferStmt, err = db.PrepareContext(ctx, getTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransfer: %w", err)
	}
	if q.updateAccountStmt, err = db.PrepareContext(ctx, updateAccount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccount: %w", err)
	}
	if q.updateEntryStmt, err = db.PrepareContext(ctx, updateEntry); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateEntry: %w", err)
	}
	if q.updateTransferStmt, err = db.PrepareContext(ctx, updateTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTransfer: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createEntriesStmt != nil {
		if cerr := q.createEntriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEntriesStmt: %w", cerr)
		}
	}
	if q.createTransferStmt != nil {
		if cerr := q.createTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTransferStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteEntryStmt != nil {
		if cerr := q.deleteEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEntryStmt: %w", cerr)
		}
	}
	if q.deleteTransferStmt != nil {
		if cerr := q.deleteTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTransferStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getEntriesStmt != nil {
		if cerr := q.getEntriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEntriesStmt: %w", cerr)
		}
	}
	if q.getEntryStmt != nil {
		if cerr := q.getEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEntryStmt: %w", cerr)
		}
	}
	if q.getListAccountStmt != nil {
		if cerr := q.getListAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListAccountStmt: %w", cerr)
		}
	}
	if q.getListTransfersStmt != nil {
		if cerr := q.getListTransfersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListTransfersStmt: %w", cerr)
		}
	}
	if q.getTransferStmt != nil {
		if cerr := q.getTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransferStmt: %w", cerr)
		}
	}
	if q.updateAccountStmt != nil {
		if cerr := q.updateAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountStmt: %w", cerr)
		}
	}
	if q.updateEntryStmt != nil {
		if cerr := q.updateEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateEntryStmt: %w", cerr)
		}
	}
	if q.updateTransferStmt != nil {
		if cerr := q.updateTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTransferStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   DBTX
	tx                   *sql.Tx
	createAccountStmt    *sql.Stmt
	createEntriesStmt    *sql.Stmt
	createTransferStmt   *sql.Stmt
	deleteAccountStmt    *sql.Stmt
	deleteEntryStmt      *sql.Stmt
	deleteTransferStmt   *sql.Stmt
	getAccountStmt       *sql.Stmt
	getEntriesStmt       *sql.Stmt
	getEntryStmt         *sql.Stmt
	getListAccountStmt   *sql.Stmt
	getListTransfersStmt *sql.Stmt
	getTransferStmt      *sql.Stmt
	updateAccountStmt    *sql.Stmt
	updateEntryStmt      *sql.Stmt
	updateTransferStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createAccountStmt:    q.createAccountStmt,
		createEntriesStmt:    q.createEntriesStmt,
		createTransferStmt:   q.createTransferStmt,
		deleteAccountStmt:    q.deleteAccountStmt,
		deleteEntryStmt:      q.deleteEntryStmt,
		deleteTransferStmt:   q.deleteTransferStmt,
		getAccountStmt:       q.getAccountStmt,
		getEntriesStmt:       q.getEntriesStmt,
		getEntryStmt:         q.getEntryStmt,
		getListAccountStmt:   q.getListAccountStmt,
		getListTransfersStmt: q.getListTransfersStmt,
		getTransferStmt:      q.getTransferStmt,
		updateAccountStmt:    q.updateAccountStmt,
		updateEntryStmt:      q.updateEntryStmt,
		updateTransferStmt:   q.updateTransferStmt,
	}
}
