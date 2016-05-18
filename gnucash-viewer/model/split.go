package model

import (
	"fmt"

	"github.com/mmbros/go/gnucash-viewer/numeric"
)

// Split type
type Split struct {
	ID              string          `xml:"id"`
	ReconciledState string          `xml:"reconciled-state"`
	ReconcileDate   GncDate         `xml:"reconcile-date>date"`
	Value           numeric.Numeric `xml:"value"`
	Memo            string          `xml:"memo"`
	Quantity        numeric.Numeric `xml:"quantity"`
	AccountID       string          `xml:"account"`

	// ReconcileDate time.Time
	Account *Account
}

func (s *Split) String() string {
	return fmt.Sprintf("Split{Value: %v, ReconcileDate: %s, Memo: %s}", s.Value, s.ReconcileDate, s.Memo)
}
