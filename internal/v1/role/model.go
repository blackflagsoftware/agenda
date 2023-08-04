package role

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Role struct {
		Id   int         `db:"id" json:"id"`
		Name null.String `db:"name" json:"name"`
	}

	RoleParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const RoleConst = "role"

func InitStorage() DataRoleAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
