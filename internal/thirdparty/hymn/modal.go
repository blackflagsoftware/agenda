package hymn

import "github.com/blackflagsoftware/agenda/config"

type (
	HymnAdapter interface {
		LookupSheet(string) (HymnSheet, error)
	}

	HymnSheet struct {
		Opening      string
		Sacrament    string
		Intermediate string
		Closing      string
	}
)

func NewHymnAdapter() HymnAdapter {
	if config.GoogleSheet != "" {
		return &GoogleSheet{SheetName: config.GoogleSheet}
	}
	return &Local{}
}
