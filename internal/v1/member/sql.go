package member

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLMember struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLMember {
	db := stor.SqliteInit()
	return &SQLMember{DB: db}
}

func (d *SQLMember) Read(mem *Member) error {
	sqlGet := `
		SELECT
			id,
			first_name,
			last_name,
			gender,
			last_prayed,
			last_talked,
			active,
			no_prayer,
			no_talk
		FROM member WHERE id = $1`
	if errDB := d.DB.Get(mem, sqlGet, mem.Id); errDB != nil {
		return ae.DBError("Member Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLMember) ReadAll(mem *[]Member, param MemberParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			first_name,
			last_name,
			gender,
			last_prayed,
			last_talked,
			active,
			no_prayer,
			no_talk
		FROM member
		%s
		ORDER BY %s %s`, searchStmt, param.Sort, param.Limit)
	sqlSearch = d.DB.Rebind(sqlSearch)
	fmt.Println(sqlSearch)
	fmt.Printf("%+v\n", args)
	if errDB := d.DB.Select(mem, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Member Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM member
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("member Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLMember) Create(mem *Member) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	mem.Id = count
	sqlPost := `
		INSERT INTO member (
			id,
			first_name,
			last_name,
			gender,
			last_prayed,
			last_talked,
			active,
			no_prayer,
			no_talk
		) VALUES (
			:id,
			:first_name,
			:last_name,
			:gender,
			:last_prayed,
			:last_talked,
			:active,
			:no_prayer,
			:no_talk
		)`
	_, errDB := d.DB.NamedExec(sqlPost, mem)
	if errDB != nil {
		return ae.DBError("Member Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLMember) Update(mem Member) error {
	sqlPatch := `
		UPDATE member SET
			first_name = :first_name,
			last_name = :last_name,
			gender = :gender,
			last_prayed = :last_prayed,
			last_talked = :last_talked,
			active = :active,
			no_prayer = :no_prayer,
			no_talk = :no_talk
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, mem); errDB != nil {
		return ae.DBError("Member Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLMember) Delete(mem *Member) error {
	sqlDelete := `
		DELETE FROM member WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, mem.Id); errDB != nil {
		return ae.DBError("Member Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLMember) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM member"); errDB != nil {
		return 0, ae.DBError("Member count: unable to get count.", errDB)
	}
	return count + 1, nil
}
