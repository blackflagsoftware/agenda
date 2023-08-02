package agenda

import (
	"fmt"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	stor "github.com/blackflagsoftware/agenda/internal/storage"
	"github.com/blackflagsoftware/agenda/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	SQLAgenda struct {
		DB *sqlx.DB
	}
)

func InitSQL() *SQLAgenda {
	db := stor.SqliteInit()
	return &SQLAgenda{DB: db}
}

func (d *SQLAgenda) Read(age *Agenda) error {
	sqlGet := `
		SELECT
			date,
			presiding,
			conducting,
			organist,
			chorister,
			newsletter,
			opening_hymn,
			sacrament_hymn, 
			intermediate_hymn,
			musical_number,
			closing_hymn,
			invocation,
			benediction,
			ward_business,
			bishop_business,
			letter_read,
			stake_business,
			stake,
			new_members,
			ordinance,
			fast_sunday,
			agenda_published,
			program_published
		FROM agenda WHERE date = $1`
	if errDB := d.DB.Get(age, sqlGet, age.Date); errDB != nil {
		return ae.DBError("Agenda Get: unable to get record.", errDB)
	}
	return nil
}

func (d *SQLAgenda) ReadAll(age *[]Agenda, param AgendaParam) (int, error) {
	searchStmt, args := util.BuildSearchString(param.Search)
	sqlSearch := fmt.Sprintf(`
		SELECT
			date,
			presiding,
			conducting,
			organist,
			chorister,
			newsletter,
			opening_hymn,
			sacrament_hymn, 
			intermediate_hymn,
			musical_number,
			closing_hymn,
			invocation,
			benediction,
			ward_business,
			bishop_business,
			letter_read,
			stake_business,
			stake,
			new_members,
			ordinance,
			fast_sunday,
			agenda_published,
			program_published
		FROM agenda
		%s
		ORDER BY %s %s`, searchStmt, param.Sort, param.Limit)
	sqlSearch = d.DB.Rebind(sqlSearch)
	if errDB := d.DB.Select(age, sqlSearch, args...); errDB != nil {
		return 0, ae.DBError("Agenda Search: unable to select records.", errDB)
	}
	sqlCount := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM agenda
		%s`, searchStmt)
	var count int
	sqlCount = d.DB.Rebind(sqlCount)
	if errDB := d.DB.Get(&count, sqlCount, args...); errDB != nil {
		return 0, ae.DBError("agenda Search: unable to select count.", errDB)
	}
	return count, nil
}

func (d *SQLAgenda) Create(age *Agenda) error {
	sqlPost := `
		INSERT INTO agenda (
			date,
			presiding,
			conducting,
			organist,
			chorister,
			newsletter,
			opening_hymn,
			sacrament_hymn, 
			intermediate_hymn,
			musical_number,
			closing_hymn,
			invocation,
			benediction,
			ward_business,
			bishop_business,
			letter_read,
			stake_business,
			stake,
			new_members,
			ordinance,
			fast_sunday,
			agenda_published,
			program_published
		) VALUES (
			:date,
			:presiding,
			:conducting,
			:organist,
			:chorister,
			:newsletter,
			:opening_hymn,
			:sacrament_hymn,
			:intermediate_hymn,
			:musical_number,
			:closing_hymn,
			:invocation,
			:benediction,
			:ward_business,
			:bishop_business,
			:letter_read,
			:stake_business,
			:stake,
			:new_members,
			:ordinance,
			:fast_sunday,
			:agenda_published,
			:program_published
		)`
	_, errDB := d.DB.NamedExec(sqlPost, age)
	if errDB != nil {
		return ae.DBError("Agenda Post: unable to insert record.", errDB)
	}

	return nil
}

func (d *SQLAgenda) Update(age Agenda) error {
	sqlPatch := `
		UPDATE agenda SET
			presiding = :presiding,
			conducting = :conducting,
			organist = :organist,
			chorister = :chorister,
			newsletter = :newsletter,
			opening_hymn = :opening_hymn,
			sacrament_hymn = :sacrament_hymn,
			intermediate_hymn = :intermediate_hymn,
			musical_number = :musical_number,
			closing_hymn = :closing_hymn,
			invocation = :invocation,
			benediction = :benediction,
			ward_business = :ward_business,
			bishop_business = :bishop_business,
			letter_read = :letter_read,
			stake_business = :stake_business,
			stake = :stake,
			new_members = :new_members,
			ordinance = :ordinance,
			fast_sunday = :fast_sunday,
			agenda_published = :agenda_published,
			program_published = :program_published
		WHERE date = :date`
	if _, errDB := d.DB.NamedExec(sqlPatch, age); errDB != nil {
		return ae.DBError("Agenda Patch: unable to update record.", errDB)
	}
	return nil
}

func (d *SQLAgenda) Delete(age *Agenda) error {
	sqlDelete := `
		DELETE FROM agenda WHERE date = $1`
	if _, errDB := d.DB.Exec(sqlDelete, age.Date); errDB != nil {
		return ae.DBError("Agenda Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *SQLAgenda) Check(age *Agenda) error {
	sqlGet := `SELECT EXISTS (SELECT date FROM agenda WHERE date = $1)`
	exists := false
	if errDB := d.DB.Get(&exists, sqlGet, age.Date); errDB != nil {
		return ae.DBError("Agenda Check: unable to get record.", errDB)
	}
	if exists {
		return fmt.Errorf("Record exists")
	}
	return nil
}
