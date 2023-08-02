package ordinance

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Ordinance struct {
		Id            int         `db:"id" json:"id"`
		Date          null.String `db:"date" json:"date"`
		Confirmations null.String `db:"confirmations" json:"confirmations"`
		Blessings     null.String `db:"blessings" json:"blessings"`
	}

	OrdinanceParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const OrdinanceConst = "ordinance"

func InitStorage() DataOrdinanceAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
