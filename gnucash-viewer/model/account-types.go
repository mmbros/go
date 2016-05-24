package model

import (
	"encoding/xml"
	"fmt"
)

// AccountType ...
type AccountType int

// AccountType constants
const (
	AccountTypeRoot AccountType = iota
	AccountTypeLiability
	AccountTypeAsset
	AccountTypeReceivible
	AccountTypeExpense
	AccountTypeIncome
	AccountTypeEquity
	AccountTypeBank
	AccountTypeCash
	AccountTypeCredit
)

// AccountTypeInfo type
type AccountTypeInfo struct {
	label        string
	root         bool
	invertValues bool
	plusLabel    string
	minusLabel   string
}

// AccountTypeInfos is the map of all AccountType
var AccountTypeInfos = map[AccountType]AccountTypeInfo{
	AccountTypeRoot: AccountTypeInfo{
		label: "Root",
		root:  true},
	AccountTypeLiability: AccountTypeInfo{
		label:        "Liability",
		invertValues: true,
		plusLabel:    "Decrease",
		minusLabel:   "Increase",
	},
	AccountTypeAsset: AccountTypeInfo{
		label:      "Asset",
		plusLabel:  "Increase",
		minusLabel: "Decrease",
	},
	AccountTypeReceivible: AccountTypeInfo{
		label:      "Receivible",
		plusLabel:  "Increase",
		minusLabel: "Decrease",
	},
	AccountTypeExpense: AccountTypeInfo{
		label:      "Expense",
		plusLabel:  "Expense",
		minusLabel: "Rebate",
	},
	AccountTypeIncome: AccountTypeInfo{
		label:        "Income",
		invertValues: true,
		plusLabel:    "Charge",
		minusLabel:   "Income",
	},
	AccountTypeEquity: AccountTypeInfo{
		label:        "Equity",
		invertValues: true,
		plusLabel:    "Decrease",
		minusLabel:   "Increase",
	},
	AccountTypeBank: AccountTypeInfo{
		label:      "Bank",
		plusLabel:  "Deposit",
		minusLabel: "Withdrawal",
	},
	AccountTypeCash: AccountTypeInfo{
		label:      "Cash",
		plusLabel:  "Receive",
		minusLabel: "Spend",
	},
	AccountTypeCredit: AccountTypeInfo{
		label:      "Credit",
		plusLabel:  "Increase",
		minusLabel: "Decrease",
	},
}

// UnmarshalXML ..
func (at *AccountType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	var s2i = map[string]AccountType{
		"ROOT":       AccountTypeRoot,
		"LIABILITY":  AccountTypeLiability,
		"ASSET":      AccountTypeAsset,
		"RECEIVABLE": AccountTypeReceivible,
		"EXPENSE":    AccountTypeExpense,
		"INCOME":     AccountTypeIncome,
		"EQUITY":     AccountTypeEquity,
		"BANK":       AccountTypeBank,
		"CASH":       AccountTypeCash,
		"CREDIT":     AccountTypeCredit,
	}

	accType, ok := s2i[content]
	if !ok {
		return fmt.Errorf("Invalid AccountType: %s", content)
	}
	*at = accType

	return nil
}

// Info returns the AccountType's information object: AccountTypeInfo.
func (at AccountType) Info() AccountTypeInfo {
	info, ok := AccountTypeInfos[at]
	if !ok {
		panic("Invalid AccountType")
	}
	return info
}

func (at AccountType) String() string {
	return at.Info().label
}
