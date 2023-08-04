package role

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLRole struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLRole {
	db := stor.SqliteInit()
	return &SQLRole{DB: db}
}

func (d *SQLRole) Read(rol *Role) error {
	sqlGet := `
		SELECT
			id,
			name
		FROM role WHERE id = $1`
	if errDB := d.DB.Get(rol, sqlGet, rol.Id); errDB != nil {
		return ae.DBError("Role Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLRole) ReadAll(rol *[]Role, param RoleParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			name
		FROM role
		%s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(rol, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Role Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM role
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("role Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLRole) Create(rol *Role) error {
	sqlPost := `
		INSERT INTO role (
			id,
			name
		) VALUES (
			:id,
			:name
		)`
	_, errDB := d.DB.NamedExec(sqlPost, rol)
	if errDB != nil {
		return ae.DBError("Role Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLRole) Update(rol Role) error {
	sqlPatch := `
		UPDATE role SET
			name = :name
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, rol); errDB != nil {
		return ae.DBError("Role Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLRole) Delete(rol *Role) error {
	sqlDelete := `
		DELETE FROM role WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, rol.Id); errDB != nil {
		return ae.DBError("Role Delete: unable to delete record.", errDB)
	}
	return nil
}
