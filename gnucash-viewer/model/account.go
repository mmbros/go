package model

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
)

// Accounts type
type Accounts struct {
	Map  AccountMap
	Root *Account
}

// AccountMap type
type AccountMap map[string]*Account

// AccountList type
type AccountList []*Account

// Account type
type Account struct {
	ID          string      `xml:"id"`
	Type        AccountType `xml:"type"`
	Name        string      `xml:"name"`
	Description string      `xml:"description"`
	ParentID    string      `xml:"parent"`
	Currency    string      `xml:"commodity>id"`

	Parent   *Account
	Children []*Account
	//	AccountTransactionList []*AccountTransaction
}

func (a *Account) String() string {
	return fmt.Sprintf("Account{Name: %s, Type: %v}", a.Name, a.Type)
}

// UnmarshalXML function unmarshal an AccountMap variable
// func (am *AccountMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//
// 	if *am == nil {
// 		*am = AccountMap{}
// 	}
//
// 	a := Account{}
// 	d.DecodeElement(&a, &start)
// 	(*am)[a.ID] = &a
//
// 	return nil
// }

// UnmarshalXML function unmarshal an AccountMap variable
func (accounts *Accounts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	if accounts.Map == nil {
		accounts.Map = AccountMap{}
		// accounts.Root = nil
	}

	a := Account{}
	d.DecodeElement(&a, &start)
	accounts.Map[a.ID] = &a

	return nil
}

// Len return the number of accounts
func (accounts *Accounts) Len() int {
	if (accounts == nil) || (accounts.Map == nil) {
		return 0
	}
	return len(accounts.Map)
}

func (accounts *Accounts) postLoadXML() error {

	// step 1: initilize root account and parent/children fields
	for _, a := range accounts.Map {

		if len(a.ParentID) == 0 {
			// found root account
			if a.Type != AccountTypeRoot {
				return fmt.Errorf("Account of type ROOT can't have parent: Account.ID = %s", a.ID)
			}
			if accounts.Root != nil {
				return fmt.Errorf("Not Implemented: multiple ROOT account")
			}

			accounts.Root = a
			fmt.Printf("*************** ROOT: %v\n", a)

		} else {
			// not root account: set parent and children
			parent := accounts.Map[a.ParentID]
			if parent == nil {
				return fmt.Errorf("Parent account not found: ParentID = %s", a.ParentID)
			}

			a.Parent = parent
			parent.Children = append(parent.Children, a)
		}
	}

	// step 2: sort each account.children by name
	for _, a := range accounts.Map {
		sort.Sort(byAccountName(a.Children))
	}

	return nil
}

// used to sort each Account.children list
type byAccountName []*Account

func (a byAccountName) Len() int           { return len(a) }
func (a byAccountName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAccountName) Less(i, j int) bool { return strings.Compare(a[i].Name, a[j].Name) < 0 }

// auxPrintTree is a PrintTree auxiliary function
func auxPrintTree(a *Account, level int, indent string) {
	fmt.Printf("%s[%s] %s (%s)\n", strings.Repeat(indent, level), a.Type, a.Name, a.Currency)

	for _, child := range a.Children {
		auxPrintTree(child, level+1, indent)
	}
}

// PrintTree prints account tree
func (accounts *Accounts) PrintTree(indent string) {
	if indent == "" {
		indent = "  "
	}

	if (accounts == nil) || (accounts.Root == nil) {
		fmt.Println("<nil>")
		return
	}

	auxPrintTree(accounts.Root, 0, indent)
}
