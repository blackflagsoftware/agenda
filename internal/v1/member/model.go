package member

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Member struct {
		Id         int         `db:"id" json:"id"`
		FirstName  null.String `db:"first_name" json:"first_name"`
		LastName   null.String `db:"last_name" json:"last_name"`
		Gender     null.String `db:"gender" json:"gender"`
		LastPrayed null.String `db:"last_prayed" json:"last_prayed"`
		LastTalked null.String `db:"last_talked" json:"last_talked"`
		Active     null.Bool   `db:"active" json:"active"`
		NoPrayer   null.Bool   `db:"no_prayer" json:"no_prayer"`
		NoTalk     null.Bool   `db:"no_talk" json:"no_talk"`
	}

	MemberParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const MemberConst = "member"

func InitStorage() DataMemberAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
