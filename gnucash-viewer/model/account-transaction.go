package model

import "fmt"

// AccountTransaction type
type AccountTransaction struct {
	Transaction *Transaction
	Split       *Split
	PlusValue   Numeric
	MinusValue  Numeric
	Balance     Numeric
}

// Description returns Split.Memo if not null, else Transaction.Description.
func (at *AccountTransaction) Description() string {
	if at == nil || at.Transaction == nil {
		return "<nil>"
	}
	if at.Split != nil && len(at.Split.Memo) > 0 {
		return at.Split.Memo
	}
	return at.Transaction.Description
}

func (at *AccountTransaction) String() string {
	return fmt.Sprintf("AccTrx{Plus:%0.2f, Minus:%0.2f, Balance:%0.2f}", at.PlusValue.Float64(), at.MinusValue.Float64(), at.Balance.Float64())
}

// Account returns Split.Account of the AccountTransaction object
func (at *AccountTransaction) Account() *Account {
	return at.Split.Account
}
