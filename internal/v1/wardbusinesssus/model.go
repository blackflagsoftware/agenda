package wardbusinesssus

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	WardBusinessSus struct {
		Id      int         `db:"id" json:"id"`
		Date    null.String `db:"date" json:"date"`
		Name    null.String `db:"name" json:"name"`
		Calling null.String `db:"calling" json:"calling"`
	}

	WardBusinessSusParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const WardBusinessSusConst = "ward_business_sus"

func InitStorage() DataWardBusinessSusAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
