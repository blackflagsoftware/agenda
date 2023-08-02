package hymn

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Hymn struct {
		Id   int         `db:"id" json:"id"`
		Name null.String `db:"name" json:"name"`
	}

	HymnParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const HymnConst = "hymn"

func InitStorage() DataHymnAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
