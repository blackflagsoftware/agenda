package speaker

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLSpeaker struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLSpeaker {
	db := stor.SqliteInit()
	return &SQLSpeaker{DB: db}
}

func (d *SQLSpeaker) Read(spe *Speaker) error {
	sqlGet := `
		SELECT
			id,
			date,
			position,
			speaker_type,
			name
		FROM speaker WHERE id = $1`
	if errDB := d.DB.Get(spe, sqlGet, spe.Id); errDB != nil {
		return ae.DBError("Speaker Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLSpeaker) ReadAll(spe *[]Speaker, param SpeakerParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			position,
			speaker_type,
			name
		FROM speaker
		%s ORDER BY position`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(spe, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Speaker Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM speaker
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("speaker Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLSpeaker) Create(spe *Speaker) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	spe.Id = count
	sqlPost := `
		INSERT INTO speaker (
			id,
			date,
			position,
			speaker_type,
			name
		) VALUES (
			:id,
			:date,
			:position,
			:speaker_type,
			:name
		)`
	_, errDB := d.DB.NamedExec(sqlPost, spe)
	if errDB != nil {
		return ae.DBError("Speaker Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLSpeaker) Update(spe Speaker) error {
	sqlPatch := `
		UPDATE speaker SET
			date = :date,
			position = :position,
			speaker_type = :speaker_type,
			name = :name
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, spe); errDB != nil {
		return ae.DBError("Speaker Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLSpeaker) Delete(spe *Speaker) error {
	sqlDelete := `
		DELETE FROM speaker WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, spe.Id); errDB != nil {
		return ae.DBError("Speaker Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLSpeaker) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM speaker"); errDB != nil {
		return 0, ae.DBError("Speaker count: unable to get count.", errDB)
	}
	return count + 1, nil
}
