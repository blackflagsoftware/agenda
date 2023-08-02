package roleuser

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	RoleUser struct {
		Id     int         `db:"id" json:"id"`
		RoleId null.Int    `db:"role_id" json:"role_id"`
		Name   null.String `db:"name" json:"name"`
		Pwd    null.String `db:"pwd" json:"pwd"`
	}

	RoleUserParam struct {
		// TODO: add any other custom params here
		util.Param
	}

	RoleLogin struct {
		Role string `db:"name" json:"role"`
	}
)

const RoleUserConst = "role_user"

func InitStorage() DataRoleUserAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
