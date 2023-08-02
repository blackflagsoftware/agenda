package member

import (
	"fmt"
	"os"
	"strings"

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

func (d *SQLMember) Splice() error {
	prayers := []Prayers{}
	speakers := []SpeakerTalk{}
	qryPrayer := "select * from prayer"
	qrySpeaker := "select * from speaker_talk"

	errPray := d.DB.Select(&prayers, qryPrayer)
	if errPray != nil {
		return errPray
	}
	errSpe := d.DB.Select(&speakers, qrySpeaker)
	if errSpe != nil {
		return errSpe
	}
	lines := []string{"insert into member (id, first_name, last_name, gender, last_prayed, last_talked, active, no_prayer, no_talk) values"}
	for _, p := range prayers {
		matchIdx := -1
		for i, s := range speakers {
			if p.Name.String == s.Name.String {
				matchIdx = i
				break
			}
			if strings.Contains(p.Name.String, s.Name.String) {
				matchIdx = i
				break
			}
		}
		// create table member (
		// 	id integer,
		// 	first_name text,
		// 	last_name text,
		// 	gender text,
		// 	last_prayed text,
		// 	last_talked text,
		// 	active boolean,
		// 	no_prayer boolean,
		// 	no_talk boolean,
		// 	primary key(id)
		// );

		split := strings.Split(p.Name.String, ",")
		fmt.Println("%+v\n", split)
		last := strings.ReplaceAll(split[0], "'", "''")
		first := strings.TrimSpace(strings.ReplaceAll(split[1], "'", "''"))
		gender := "Female"
		lastPrayed := "0001-01-01"
		if p.LastPrayedDate.String != "2001-01-01" {
			lastPrayed = p.LastPrayedDate.String
		}
		lastTalked := "0001-01-01"
		active := false
		noTalk := false
		if matchIdx > -1 {
			if speakers[matchIdx].Gender.String == "M" {
				gender = "Male"
			}
			if speakers[matchIdx].LastTalked.String != "200-01-01" {
				lastTalked = speakers[matchIdx].LastTalked.String
			}
			active = speakers[matchIdx].Active.Bool
			noTalk = speakers[matchIdx].Rntt.Bool
		}
		line := fmt.Sprintf("(%d, '%s', '%s', '%s', '%s', '%s', %t, %t, %t),", p.Id, first, last, gender, lastPrayed, lastTalked, active, p.Rntp.Bool, noTalk)
		lines = append(lines, line)
	}
	f, err := os.OpenFile("../../scripts/member-insert.sql", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	f.WriteString(strings.Join(lines, "\n"))
	return nil
}
