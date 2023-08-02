package defaultcalling

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	DefaultCalling struct {
		Id         int         `db:"id" json:"id"`
		Bishop     null.String `db:"bishop" json:"bishop"`
		B1st       null.String `db:"b_1st" json:"b_1st"`
		B2nd       null.String `db:"b_2nd" json:"b_2nd"`
		SPres      null.String `db:"s_pres" json:"s_pres"`
		S1st       null.String `db:"s_1st" json:"s_1st"`
		S2nd       null.String `db:"s_2nd" json:"s_2nd"`
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
