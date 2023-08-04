package roleuser

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLRoleUser struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLRoleUser {
	db := stor.SqliteInit()
	return &SQLRoleUser{DB: db}
}

func (d *SQLRoleUser) Login(ro RoleUser) (RoleLogin, error) {
	roleLogin := RoleLogin{}
	sqlGet := `
		SELECT
			r.name
		FROM role_user AS ru
		INNER JOIN role AS r ON ru.role_id = r.id
		WHERE ru.name = $1
			AND ru.pwd = $2`
	if errDB := d.DB.Get(&roleLogin, sqlGet, ro.Name.String, ro.Pwd.String); errDB != nil {
		return roleLogin, ae.DBError("RoleUser Get: unable to get record.", errDB)
	}
	return roleLogin, nil
}

func (d *SQLRoleUser) Read(ro *RoleUser) error {
	sqlGet := `
		SELECT
			id,
			role_id,
			name,
			pwd
		FROM role_user WHERE id = $1`
	if errDB := d.DB.Get(ro, sqlGet, ro.Id); errDB != nil {
		return ae.DBError("RoleUser Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLRoleUser) ReadAll(ro *[]RoleUser, param RoleUserParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			role_id,
			name,
			pwd
		FROM role_user
		%s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(ro, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("RoleUser Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM role_user
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("role_user Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLRoleUser) Create(ro *RoleUser) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	ro.Id = count
	sqlPost := `
		INSERT INTO role_user (
			id,
			role_id,
			name,
			pwd
		) VALUES (
			:id,
			:role_id,
			:name,
			:pwd
		)`
	_, errDB := d.DB.NamedExec(sqlPost, ro)
	if errDB != nil {
		return ae.DBError("RoleUser Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLRoleUser) Update(ro RoleUser) error {
	sqlPatch := `
		UPDATE role_user SET
			role_id = :role_id,
			name = :name,
			pwd = :pwd
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, ro); errDB != nil {
		return ae.DBError("RoleUser Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLRoleUser) Delete(ro *RoleUser) error {
	sqlDelete := `
		DELETE FROM role_user WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, ro.Id); errDB != nil {
		return ae.DBError("RoleUser Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLRoleUser) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM role_user"); errDB != nil {
		return 0, ae.DBError("RoleUser count: unable to get count.", errDB)
	}
	return count + 1, nil
}
