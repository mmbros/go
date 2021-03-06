package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mmbros/go/gnucash-viewer/model"
	"github.com/mmbros/go/gnucash-viewer/report"
)

var gnucashPath = flag.String("gnucash-file", "data/data.gnucash", "GnuCash file path")

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n--\nfunction %s took %v\n", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "task duration:")

	gnc, err := model.LoadFromXMLFile(*gnucashPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("books: %d\n", len(gnc.Books))

	if len(gnc.Books) != 1 {
		return
	}
	book := gnc.Books[0]
	accounts := book.Accounts

	fmt.Printf("accounts: %d\n", accounts.Len())
	fmt.Printf("accounts.ROOT: %v\n", accounts.Root)

	fmt.Printf("tansactions: %d\n", len(book.Transactions))

	// acc := book.AccountByName("Conto Arancio")
	acc := book.AccountByName("Stipendio")
	fmt.Printf("ACCOUNT %v\n", acc)
	acc.PrintAccTrxList()

	// for j, at := range acc.AccountTransactionList {
	// 	fmt.Printf("%02d) %s %s %5.2f %7.2f %7.0f\n",
	t := book.Transactions[0]
	fmt.Printf("%v\n", t)

	fmt.Printf("splits: %d\n", len(t.Splits))

	s := t.Splits[0]
	fmt.Println(s)

	// accounts.PrintTree("    ")

	dateFrom := model.NewDate(2016, 1, 2, false)
	dateTo := model.NewDate(2016, 1, 2, true)

	report.Uscite(book, dateFrom, dateTo)
}
