package model

import "fmt"

// Transaction type
type Transaction struct {
	ID          string   `xml:"id"`
	Currency    string   `xml:"currency>id"`
	DatePosted  Date     `xml:"date-posted>date"`
	DateEntered Date     `xml:"date-entered>date"`
	Description string   `xml:"description"`
	Splits      []*Split `xml:"splits>split"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("Transaction{DatePosted: %v, DateEntered: %v, Currency: %s}", t.DatePosted, t.DateEntered, t.Currency)
	// return fmt.Sprintf("Transaction{ID: %s, DatePosted: %v, Currency: %s}", t.ID, t.DatePosted, t.Currency)
}

// used to sort Transactions
type byDatePosted []*Transaction

func (t byDatePosted) Len() int           { return len(t) }
func (t byDatePosted) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byDatePosted) Less(i, j int) bool { return t[i].DatePosted.Before(t[j].DatePosted) }
