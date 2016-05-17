package model

// Transaction type
type Transaction struct {
	ID          string  `xml:"id"`
	Currency    string  `xml:"currency>id"`
	DatePosted  string  `xml:"date-posted>date"`
	DateEntered string  `xml:"date-entered>date"`
	Description string  `xml:"description"`
	SplitList   []Split `xml:"splits>split"`
}
