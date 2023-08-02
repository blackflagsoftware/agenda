package agenda

import (
	"github.com/blackflagsoftware/agenda/config"
	"github.com/blackflagsoftware/agenda/internal/util"
	"gopkg.in/guregu/null.v3"
)

type (
	Agenda struct {
		Date             string      `db:"date" json:"date"`
		Presiding        null.String `db:"presiding" json:"presiding"`
		Conducting       null.String `db:"conducting" json:"conducting"`
		Organist         null.String `db:"organist" json:"organist"`
		Chorister        null.String `db:"chorister" json:"chorister"`
		Newsletter       null.String `db:"newsletter" json:"newsletter"`
		OpeningHymn      null.Int    `db:"opening_hymn" json:"opening_hymn"`
		SacramentHymn    null.Int    `db:"sacrament_hymn" json:"sacrament_hymn"`
		IntermediateHymn null.Int    `db:"intermediate_hymn" json:"intermediate_hymn"`
		MusicalNumber    null.String `db:"musical_number" json:"musical_number"`
		ClosingHymn      null.Int    `db:"closing_hymn" json:"closing_hymn"`
		Invocation       null.String `db:"invocation" json:"invocation"`
		Benediction      null.String `db:"benediction" json:"benediction"`
		WardBusiness     null.Bool   `db:"ward_business" json:"ward_business"`
		BishopBusiness   null.Bool   `db:"bishop_business" json:"bishop_business"`
		LetterRead       null.Bool   `db:"letter_read" json:"letter_read"`
		StakeBusiness    null.Bool   `db:"stake_business" json:"stake_business"`
		Stake            null.String `db:"stake" json:"stake"`
		NewMembers       null.Bool   `db:"new_members" json:"new_members"`
		Ordinance        null.Bool   `db:"ordinance" json:"ordinance"`
		Fastsunday       null.Bool   `db:"fast_sunday" json:"fast_sunday"`
		AgendaPublished  null.Bool   `db:"agenda_published" json:"agenda_published"`
		ProgramPublished null.Bool   `db:"program_published" json:"program_published"`
	}

	AgendaParam struct {
		// TODO: add any other custom params here
		util.Param
	}
)

const AgendaConst = "agenda"

func InitStorage() DataAgendaAdapter {
	if config.StorageSQL {
		return InitSQL()
	}
	return nil
}
