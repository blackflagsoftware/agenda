package announcement

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLAnnouncement struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLAnnouncement {
	db := stor.SqliteInit()
	return &SQLAnnouncement{DB: db}
}

func (d *SQLAnnouncement) Read(ann *Announcement) error {
	sqlGet := `
		SELECT
			id,
			date,
			message,
			pulpit
		FROM announcement where id = $1`
	if errDB := d.DB.Get(ann, sqlGet, ann.Id); errDB != nil {
		return ae.DBError("Announcement Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLAnnouncement) ReadAll(ann *[]Announcement, param AnnouncementParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			id,
			date,
			message,
			pulpit
		FROM announcement`)
	// %s`, searchStmt)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(ann, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Announcement Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM announcement
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("announcement Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLAnnouncement) Create(ann *Announcement) error {
	count, errCount := d.count()
	if errCount != nil {
		return errCount
	}
	ann.Id = count
	sqlPost := `
		INSERT INTO announcement (
			id,
			date,
			message,
			pulpit
		) VALUES (
			:id,
			:date,
			:message,
			:pulpit
		)`
	_, errDB := d.DB.NamedExec(sqlPost, ann)
	if errDB != nil {
		return ae.DBError("Announcement Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLAnnouncement) Update(ann Announcement) error {
	sqlPatch := `
		UPDATE announcement SET
			date = :date,
			message = :message,
			pulpit = :pulpit
		WHERE id = :id`
	if _, errDB := d.DB.NamedExec(sqlPatch, ann); errDB != nil {
		return ae.DBError("Announcement Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLAnnouncement) Delete(ann *Announcement) error {
	sqlDelete := `
		DELETE FROM announcement WHERE id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, ann.Id); errDB != nil {
		return ae.DBError("Announcement Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLAnnouncement) count() (int, error) {
	count := 0
	if errDB := d.DB.Get(&count, "SELECT COALESCE(MAX(id), 0) FROM announcement"); errDB != nil {
		return 0, ae.DBError("Announcement count: unable to get count.", errDB)
	}
	return count + 1, nil
}
