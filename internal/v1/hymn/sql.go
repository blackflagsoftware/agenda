package hymn

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLHymn struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLHymn {
	db := stor.SqliteInit()
	return &SQLHymn{DB: db}
}

func (d *SQLHymn) Read(hym *Hymn) error {
	sqlGet := `
		SELECT
			id,
			name,
			pdf_name
		FROM hymn WHERE id = $1`
	if errDB := d.DB.Get(hym, sqlGet, hym.Id); errDB != nil {
		return ae.DBError("Hymn Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLHymn) ReadAll(hym *[]Hymn, param HymnParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			id || ' - ' || name AS name,
			pdf_name
		FROM hymn
		%s
		ORDER BY id`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(hym, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Hymn Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM hymn
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("hymn Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLHymn) Create(hym *Hymn) error {
	sqlPost := `
		INSERT INTO hymn (
			id,
			name,
			pdf_name
		) VALUES (
			:id,
			:name,
			:pdf_name
		)`
	_, errDB := d.DB.NamedExec(sqlPost, hym)
	if errDB != nil {
		return ae.DBError("Hymn Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLHymn) Update(hym Hymn) error {
	sqlPatch := `
		UPDATE hymn SET
			name = :name,
			pdf_name = :pdf_name
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, hym); errDB != nil {
		return ae.DBError("Hymn Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLHymn) Delete(hym *Hymn) error {
	sqlDelete := `
		DELETE FROM hymn WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, hym.Id); errDB != nil {
		return ae.DBError("Hymn Delete: unable to delete record.", errDB)
	}
	return nil
}
