package model

import (
	"fmt"
	"time"

	"github.com/mmbros/go/gnucash-viewer/numeric"
)

// Split type
type Split struct {
	ID               string          `xml:"id"`
	ReconciledState  string          `xml:"reconciled-state"`
	ReconcileDateStr string          `xml:"reconcile-date>date"`
	Value            numeric.Numeric `xml:"value"`
	Memo             string          `xml:"memo"`
	Quantity         numeric.Numeric `xml:"quantity"`
	AccountID        string          `xml:"account"`

	ReconcileDate time.Time
	Account       *Account
}

func (s *Split) String() string {
	return fmt.Sprintf("Split{Value: %v, Memo: %s}", s.Value, s.Memo)
}
