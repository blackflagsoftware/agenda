package announcement

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Announcement struct {
		Id      int         `db:"id" json:"id"`
		Date    null.String `db:"date" json:"date"`
		Message null.String `db:"message" json:"message"`
		Pulpit  null.Bool   `db:"pulpit" json:"pulpit"`
	}

	AnnouncementParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const AnnouncementConst = "announcement"

func InitStorage() DataAnnouncementAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
