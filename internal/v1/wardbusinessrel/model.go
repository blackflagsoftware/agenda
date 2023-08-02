package wardbusinessrel

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	WardBusinessRel struct {
		Id      int         `db:"id" json:"id"`
		Date    null.String `db:"date" json:"date"`
		Name    null.String `db:"name" json:"name"`
		Calling null.String `db:"calling" json:"calling"`
	}

	WardBusinessRelParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const WardBusinessRelConst = "ward_business_rel"

func InitStorage() DataWardBusinessRelAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
