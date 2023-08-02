package visitor

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLVisitor struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLVisitor {
	db := stor.SqliteInit()
	return &SQLVisitor{DB: db}
}

func (d *SQLVisitor) Read(vis *Visitor) error {
	sqlGet := `
		SELECT
			id,
			date,
			name
		FROM visitor WHERE id = $1`
	if errDB := d.DB.Get(vis, sqlGet, vis.Id); errDB != nil {
		return ae.DBError("Visitor Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLVisitor) ReadAll(vis *[]Visitor, param VisitorParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			name
		FROM visitor
		%s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(vis, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Visitor Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM visitor
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("visitor Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLVisitor) Create(vis *Visitor) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	vis.Id = count
	sqlPost := `
		INSERT INTO visitor (
			id,
			date,
			name
		) VALUES (
			:id,
			:date,
			:name
		)`
	fmt.Println(sqlPost)
	fmt.Printf("%+v\n", vis)
	_, errDB := d.DB.NamedExec(sqlPost, vis)
	if errDB != nil {
		return ae.DBError("Visitor Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLVisitor) Update(vis Visitor) error {
	sqlPatch := `
		UPDATE visitor SET
			date = :date,
			name = :name
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, vis); errDB != nil {
		return ae.DBError("Visitor Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLVisitor) Delete(vis *Visitor) error {
	sqlDelete := `
		DELETE FROM visitor WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, vis.Id); errDB != nil {
		return ae.DBError("Visitor Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLVisitor) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM visitor"); errDB != nil {
		return 0, ae.DBError("Visitor count: unable to get count.", errDB)
	}
	return count + 1, nil
}
