package newmember

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLNewMember struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLNewMember {
	db := stor.SqliteInit()
	return &SQLNewMember{DB: db}
}

func (d *SQLNewMember) Read(new *NewMember) error {
	sqlGet := `
		SELECT
			id,
			date,
			family_name,
			names
		FROM new_member WHERE id = $1`
	if errDB := d.DB.Get(new, sqlGet, new.Id); errDB != nil {
		return ae.DBError("NewMember Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLNewMember) ReadAll(new *[]NewMember, param NewMemberParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			family_name,
			names
		FROM new_member
		%s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(new, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("NewMember Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM new_member
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("new_member Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLNewMember) Create(new *NewMember) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	new.Id = count
	sqlPost := `
		INSERT INTO new_member (
			id,
			date,
			family_name,
			names
		) VALUES (
			:id,
			:date,
			:family_name,
			:names
		)`
	_, errDB := d.DB.NamedExec(sqlPost, new)
	if errDB != nil {
		return ae.DBError("NewMember Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLNewMember) Update(new NewMember) error {
	sqlPatch := `
		UPDATE new_member SET
			date = :date,
			family_name = :family_name,
			names = :names
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, new); errDB != nil {
		return ae.DBError("NewMember Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLNewMember) Delete(new *NewMember) error {
	sqlDelete := `
		DELETE FROM new_member WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, new.Id); errDB != nil {
		return ae.DBError("NewMember Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLNewMember) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM new_member"); errDB != nil {
		return 0, ae.DBError("NewMember count: unable to get count.", errDB)
	}
	return count + 1, nil
}
