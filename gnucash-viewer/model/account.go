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

	Parent     *Account
	Children   []*Account
	AccTrxList []*AccountTransaction
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

func (accounts *Accounts) initAccountTree() error {

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

func (accounts *Accounts) initAccountTransactionList(transactions []*Transaction) {
	// initialize each Split.Account variable and creates each Account.AccTrxList field
	// NOTE: transaction must be already ordered by DatePosted
	//       so that also AccTrxList will be ordered by DatePosted
	for _, t := range transactions {
		for _, s := range t.Splits {
			a := accounts.Map[s.AccountID]
			s.Account = a
			at := AccountTransaction{Transaction: t, Split: s}
			a.AccTrxList = append(a.AccTrxList, &at)
		}
	}

	// initialize account balance
	for _, a := range accounts.Map {
		var balance Numeric

		// recupera, in funzione dell'account type, se occorre invertire i valori
		invertValues := a.Type.Info().invertValues

		for _, at := range a.AccTrxList {

			v := at.Split.Value
			if invertValues {
				v.NegEqual()
			}

			balance.AddEqual(&v)
			at.Balance.Set(&balance)

			if v.Sign() >= 0 {
				at.PlusValue.Set(&v)
			} else {
				v.NegEqual()
				at.MinusValue.Set(&v)
			}
		}
	}
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

// PrintAccTrxList ...
func (a *Account) PrintAccTrxList() {
	info := a.Type.Info()

	fmt.Printf("  # Date       %s       %s  Balance\n", info.minusLabel, info.plusLabel)
	fmt.Println("--- ---------- -------- --------")

	for j, at := range a.AccTrxList {
		fmt.Printf("%02d) %v %8.2f %8.2f %9.2f %s\n",
			j+1,
			at.Transaction.DatePosted,
			at.PlusValue.Float64(),
			at.MinusValue.Float64(),
			at.Balance.Float64(),
			at.Description(),
		)
	}
}
