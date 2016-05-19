package model

import (
	"fmt"
)

// Split type
type Split struct {
	ID              string  `xml:"id"`
	ReconciledState string  `xml:"reconciled-state"`
	ReconcileDate   Date    `xml:"reconcile-date>date"`
	Value           Numeric `xml:"value"`
	Memo            string  `xml:"memo"`
	Quantity        Numeric `xml:"quantity"`
	AccountID       string  `xml:"account"`

	Account *Account
}

func (s *Split) String() string {
	return fmt.Sprintf("Split{Value: %v, ReconcileDate: %s, AccountName: %s, Memo: %s}", s.Value, s.ReconcileDate, s.Account.Name, s.Memo)
}
