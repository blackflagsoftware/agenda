package newmember

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	NewMember struct {
		Id         int         `db:"id" json:"id"`
		Date       null.String `db:"date" json:"date"`
		FamilyName null.String `db:"family_name" json:"family_name"`
		Names      null.String `db:"names" json:"names"`
	}

	NewMemberParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const NewMemberConst = "new_member"

func InitStorage() DataNewMemberAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
