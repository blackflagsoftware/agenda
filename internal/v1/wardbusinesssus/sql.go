package wardbusinesssus

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLWardBusinessSus struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLWardBusinessSus {
	db := stor.SqliteInit()
	return &SQLWardBusinessSus{DB: db}
}

func (d *SQLWardBusinessSus) Read(wa *WardBusinessSus) error {
	sqlGet := `
		SELECT
			id,
			date,
			name,
			calling
		FROM ward_business_sus WHERE id = $1`
	if errDB := d.DB.Get(wa, sqlGet, wa.Id); errDB != nil {
		return ae.DBError("WardBusinessSus Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLWardBusinessSus) ReadAll(wa *[]WardBusinessSus, param WardBusinessSusParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			name,
			calling
		FROM ward_business_sus
		%s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(wa, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("WardBusinessSus Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM ward_business_sus
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("ward_business_sus Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLWardBusinessSus) Create(wa *WardBusinessSus) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	wa.Id = count
	sqlPost := `
		INSERT INTO ward_business_sus (
			id,
			date,
			name,
			calling
		) VALUES (
			:id,
			:date,
			:name,
			:calling
		)`
	_, errDB := d.DB.NamedExec(sqlPost, wa)
	if errDB != nil {
		return ae.DBError("WardBusinessSus Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLWardBusinessSus) Update(wa WardBusinessSus) error {
	sqlPatch := `
		UPDATE ward_business_sus SET
			date = :date,
			name = :name,
			calling = :calling
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, wa); errDB != nil {
		return ae.DBError("WardBusinessSus Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLWardBusinessSus) Delete(wa *WardBusinessSus) error {
	sqlDelete := `
		DELETE FROM ward_business_sus WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, wa.Id); errDB != nil {
		return ae.DBError("WardBusinessSus Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLWardBusinessSus) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM ward_business_sus"); errDB != nil {
		return 0, ae.DBError("WardBusinessSus count: unable to get count.", errDB)
	}
	return count + 1, nil
}
