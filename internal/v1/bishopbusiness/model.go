package bishopbusiness

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	BishopBusiness struct {
		Id      int         `db:"id" json:"id"`
		Date    null.String `db:"date" json:"date"`
		Message null.String `db:"message" json:"message"`
	}

	BishopBusinessParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const BishopBusinessConst = "bishop_business"

func InitStorage() DataBishopBusinessAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
