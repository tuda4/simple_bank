package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all func to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes  a func within a db transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rollBackErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Tranfers `json:"transfer"`
	FromAccount Accounts `json:"from_account"`
	ToAccount   Accounts `json:"to_account"`
	FromEntry   Entries  `json:"from_entry"`
	ToEntry     Entries  `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        int32(arg.Amount),
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -int64(arg.Amount),
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    int64(arg.Amount),
		})

		if err != nil {
			return err
		}

		// ToDo: update accounts balance

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: +arg.Amount,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
