package report

import (
	"fmt"

	"github.com/mmbros/go/gnucash-viewer/model"
)

// Uscite ...
func Uscite(book model.Book, dateFrom, dateTo model.Date) {

	fmt.Printf("REPORT.USCITE from %v to %v\n", dateFrom, dateTo)

	for _, t := range book.Transactions {

		if t.DatePosted.Before(dateFrom) {
			continue
		}
		if t.DatePosted.After(dateTo) {
			break
		}

		fmt.Printf("%v\n", t)

	}
}
