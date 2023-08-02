package speaker

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Speaker struct {
		Id       int         `db:"id" json:"id"`
		Date     null.String `db:"date" json:"date"`
		Position null.String `db:"position" json:"position"`
		Name     null.String `db:"name" json:"name"`
	}

	SpeakerParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const SpeakerConst = "speaker"

func InitStorage() DataSpeakerAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
