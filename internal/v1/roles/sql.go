package roles

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLRoles struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLRoles {
	db := stor.SqliteInit()
	return &SQLRoles{DB: db}
}

func (d *SQLRoles) Read(rol *Roles) error {
	sqlGet := `
		SELECT
			id,
			name
		FROM roles WHERE id = $1`
	if errDB := d.DB.Get(rol, sqlGet, rol.Id); errDB != nil {
		return ae.DBError("Roles Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLRoles) ReadAll(rol *[]Roles, param RolesParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			name
		FROM roles
		%s
		ORDER BY %s %s`, searchStmt, param.Sort, param.Limit)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(rol, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Roles Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM roles
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("roles Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLRoles) Create(rol *Roles) error {
	sqlPost := `
		INSERT INTO roles (
			id,
			name
		) VALUES (
			:id,
			:name
		)`
	_, errDB := d.DB.NamedExec(sqlPost, rol)
	if errDB != nil {
		return ae.DBError("Roles Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLRoles) Update(rol Roles) error {
	sqlPatch := `
		UPDATE roles SET
			name = :name
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, rol); errDB != nil {
		return ae.DBError("Roles Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLRoles) Delete(rol *Roles) error {
	sqlDelete := `
		DELETE FROM roles WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, rol.Id); errDB != nil {
		return ae.DBError("Roles Delete: unable to delete record.", errDB)
	}
	return nil
}
