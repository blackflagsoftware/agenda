package roles

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Roles struct {
		Id   int         `db:"id" json:"id"`
		Name null.String `db:"name" json:"name"`
	}

	RolesParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const RolesConst = "roles"

func InitStorage() DataRolesAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
