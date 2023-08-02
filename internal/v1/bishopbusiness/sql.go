package bishopbusiness

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLBishopBusiness struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLBishopBusiness {
	db := stor.SqliteInit()
	return &SQLBishopBusiness{DB: db}
}

func (d *SQLBishopBusiness) Read(bis *BishopBusiness) error {
	sqlGet := `
		SELECT
			id,
			date,
			message
		FROM bishop_business WHERE id = $1`
	if errDB := d.DB.Get(bis, sqlGet, bis.Id); errDB != nil {
		return ae.DBError("BishopBusiness Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLBishopBusiness) ReadAll(bis *[]BishopBusiness, param BishopBusinessParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			message
		FROM bishop_business
		%s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(bis, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("BishopBusiness Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM bishop_business
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("bishop_business Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLBishopBusiness) Create(bis *BishopBusiness) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	bis.Id = count
	sqlPost := `
		INSERT INTO bishop_business (
			id,
			date,
			message
		) VALUES (
			:id,
			:date,
			:message
		)`
	_, errDB := d.DB.NamedExec(sqlPost, bis)
	if errDB != nil {
		return ae.DBError("BishopBusiness Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLBishopBusiness) Update(bis BishopBusiness) error {
	sqlPatch := `
		UPDATE bishop_business SET
			date = :date,
			message = :message
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, bis); errDB != nil {
		return ae.DBError("BishopBusiness Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLBishopBusiness) Delete(bis *BishopBusiness) error {
	sqlDelete := `
		DELETE FROM bishop_business WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, bis.Id); errDB != nil {
		return ae.DBError("BishopBusiness Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLBishopBusiness) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM bishop_business"); errDB != nil {
		return 0, ae.DBError("BishopBusiness count: unable to get count.", errDB)
	}
	return count + 1, nil
}
