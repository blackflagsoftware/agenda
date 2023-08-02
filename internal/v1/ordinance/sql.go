package ordinance

import (
	"database/sql"
	"errors"
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLOrdinance struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLOrdinance {
	db := stor.SqliteInit()
	return &SQLOrdinance{DB: db}
}

func (d *SQLOrdinance) Read(ord *Ordinance) error {
	sqlGet := `
		SELECT
			id,
			date,
			confirmations,
			blessings
		FROM ordinance WHERE date = $1`
	if errDB := d.DB.Get(ord, sqlGet, ord.Date); errDB != nil {
		if errors.Is(errDB, sql.ErrNoRows) {
			return nil
		}
		return ae.DBError("Ordinance Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLOrdinance) ReadAll(ord *[]Ordinance, param OrdinanceParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			confirmations,
			blessings
		FROM ordinance
		%s
		ORDER BY %s %s`, searchStmt, param.Sort, param.Limit)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(ord, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Ordinance Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM ordinance
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("ordinance Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLOrdinance) Create(ord *Ordinance) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	ord.Id = count
	sqlPost := `
		INSERT INTO ordinance (
			id,
			date,
			confirmations,
			blessings
		) VALUES (
			:id,
			:date,
			:confirmations,
			:blessings
		)`
	_, errDB := d.DB.NamedExec(sqlPost, ord)
	if errDB != nil {
		return ae.DBError("Ordinance Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLOrdinance) Update(ord Ordinance) error {
	sqlPatch := `
		UPDATE ordinance SET
			date = :date,
			confirmations = :confirmations,
			blessings = :blessings
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, ord); errDB != nil {
		return ae.DBError("Ordinance Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLOrdinance) Delete(ord *Ordinance) error {
	sqlDelete := `
		DELETE FROM ordinance WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, ord.Id); errDB != nil {
		return ae.DBError("Ordinance Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLOrdinance) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM ordinance"); errDB != nil {
		return 0, ae.DBError("Ordinance count: unable to get count.", errDB)
	}
	return count + 1, nil
}
