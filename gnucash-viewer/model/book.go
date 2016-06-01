package model

import (
	"encoding/xml"
	"sort"
)

// Book type
type Book struct {
	XMLName      xml.Name     `xml:"book"`
	ID           string       `xml:"id"`
	Accounts     *Accounts    `xml:"account"`
	Transactions Transactions `xml:"transaction"`
}

// AccountByName return the Book's Account with given name
// Returns nil if Account doesn't exist
func (b *Book) AccountByName(name string) *Account {
	for _, acc := range b.Accounts.Map {
		if acc.Name == name {
			return acc
		}
	}
	return nil
}

func (b *Book) postLoadXML() error {

	if err := b.Accounts.initAccountTree(); err != nil {
		return err
	}

	// sort Transactions by DatePosted
	sort.Sort(byDatePosted(b.Transactions))

	b.Accounts.initAccountTransactionList(b.Transactions)

	return nil
}
