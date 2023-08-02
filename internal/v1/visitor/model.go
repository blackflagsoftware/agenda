package visitor

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Visitor struct {
		Id   int         `db:"id" json:"id"`
		Date null.String `db:"date" json:"date"`
		Name null.String `db:"name" json:"name"`
	}

	VisitorParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const VisitorConst = "visitor"

func InitStorage() DataVisitorAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
