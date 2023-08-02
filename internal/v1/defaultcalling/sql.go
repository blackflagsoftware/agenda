package defaultcalling

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLDefaultCalling struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLDefaultCalling {
	db := stor.SqliteInit()
	return &SQLDefaultCalling{DB: db}
}

func (d *SQLDefaultCalling) Read(def *DefaultCalling) error {
	sqlGet := `
		SELECT
			id,
			bishop,
			b_1st,
			b_2nd,
			s_pres,
			s_1st,
			s_2nd,
			organist,
			chorister,
			newsletter,
			stake
		FROM default_calling WHERE id = $1`
	if errDB := d.DB.Get(def, sqlGet, def.Id); errDB != nil {
		return ae.DBError("DefaultCalling Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLDefaultCalling) ReadAll(def *[]DefaultCalling, param DefaultCallingParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			bishop,
			b_1st,
			b_2nd,
			s_pres,
			s_1st,
			s_2nd,
			organist,
			chorister,
			newsletter,
			stake
		FROM default_calling
		%s
		ORDER BY %s %s`, searchStmt, param.Sort, param.Limit)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(def, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("DefaultCalling Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM default_calling
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("default_calling Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLDefaultCalling) Create(def *DefaultCalling) error {
	sqlPost := `
		INSERT INTO default_calling (
			id,
			bishop,
			b_1st,
			b_2nd,
			s_pres,
			s_1st,
			s_2nd,
			organist,
			chorister,
			newsletter,
			stake
		) VALUES (
			:id,
			:bishop,
			:b_1st,
			:b_2nd,
			:s_pres,
			:s_1st,
			:s_2nd,
			:organist,
			:chorister,
			:newsletter,
			:stake
		)`
	_, errDB := d.DB.NamedExec(sqlPost, def)
	if errDB != nil {
		return ae.DBError("DefaultCalling Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLDefaultCalling) Update(def DefaultCalling) error {
	sqlPatch := `
		UPDATE default_calling SET
			bishop = :bishop,
			b_1st = :b_1st,
			b_2nd = :b_2nd,
			s_pres = :s_pres,
			s_1st = :s_1st,
			s_2nd = :s_2nd,
			organist = :organist,
			chorister = :chorister,
			newsletter = :newsletter,
			stake = :stake
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, def); errDB != nil {
		return ae.DBError("DefaultCalling Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLDefaultCalling) Delete(def *DefaultCalling) error {
	sqlDelete := `
		DELETE FROM default_calling WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, def.Id); errDB != nil {
		return ae.DBError("DefaultCalling Delete: unable to delete record.", errDB)
	}
	return nil
}
