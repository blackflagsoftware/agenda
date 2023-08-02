package defaultcalling

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	DefaultCalling struct {
		Id         int         `db:"id" json:"id"`
		Organist   null.String `db:"organist" json:"organist"`
		Chorister  null.String `db:"chorister" json:"chorister"`
		Newsletter null.String `db:"newsletter" json:"newsletter"`
		Stake      null.String `db:"stake" json:"stake"`
	}

	DefaultCallingParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const DefaultCallingConst = "default_calling"

func InitStorage() DataDefaultCallingAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
