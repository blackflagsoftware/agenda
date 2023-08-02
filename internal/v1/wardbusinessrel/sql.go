package wardbusinessrel

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLWardBusinessRel struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLWardBusinessRel {
	db := stor.SqliteInit()
	return &SQLWardBusinessRel{DB: db}
}

func (d *SQLWardBusinessRel) Read(war *WardBusinessRel) error {
	sqlGet := `
		SELECT
			id,
			date,
			name,
			calling
		FROM ward_business_rel WHERE id = $1`
	if errDB := d.DB.Get(war, sqlGet, war.Id); errDB != nil {
		return ae.DBError("WardBusinessRel Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLWardBusinessRel) ReadAll(war *[]WardBusinessRel, param WardBusinessRelParam) (int, error) {
	fmt.Println("******", param.Search)
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			name,
			calling
		FROM ward_business_rel
		%s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	fmt.Println(sqlSearch)
	if errDB := d.DB.Select(war, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("WardBusinessRel Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM ward_business_rel
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("ward_business_rel Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLWardBusinessRel) Create(war *WardBusinessRel) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	war.Id = count
	sqlPost := `
		INSERT INTO ward_business_rel (
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
	_, errDB := d.DB.NamedExec(sqlPost, war)
	if errDB != nil {
		return ae.DBError("WardBusinessRel Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLWardBusinessRel) Update(war WardBusinessRel) error {
	sqlPatch := `
		UPDATE ward_business_rel SET
			date = :date,
			name = :name,
			calling = :calling
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, war); errDB != nil {
		return ae.DBError("WardBusinessRel Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLWardBusinessRel) Delete(war *WardBusinessRel) error {
	sqlDelete := `
		DELETE FROM ward_business_rel WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, war.Id); errDB != nil {
		return ae.DBError("WardBusinessRel Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLWardBusinessRel) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM ward_business_rel"); errDB != nil {
		return 0, ae.DBError("WardBusinessRel count: unable to get count.", errDB)
	}
	return count + 1, nil
}
