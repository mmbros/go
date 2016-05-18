package model

import "fmt"

// Transaction type
type Transaction struct {
	ID          string   `xml:"id"`
	Currency    string   `xml:"currency>id"`
	DatePosted  GncDate  `xml:"date-posted>date"`
	DateEntered GncDate  `xml:"date-entered>date"`
	Description string   `xml:"description"`
	SplitList   []*Split `xml:"splits>split"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("Transaction{DatePosted: %v, DateEntered: %v, Currency: %s}", t.DatePosted, t.DateEntered, t.Currency)
	// return fmt.Sprintf("Transaction{ID: %s, DatePosted: %v, Currency: %s}", t.ID, t.DatePosted, t.Currency)
}
