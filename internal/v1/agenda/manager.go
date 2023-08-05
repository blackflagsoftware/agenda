package agenda

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/blackflagsoftware/agenda/config"
	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	a "github.com/blackflagsoftware/agenda/internal/audit"
	"github.com/blackflagsoftware/agenda/internal/util"
	ann "github.com/blackflagsoftware/agenda/internal/v1/announcement"
	bis "github.com/blackflagsoftware/agenda/internal/v1/bishopbusiness"
	hym "github.com/blackflagsoftware/agenda/internal/v1/hymn"
	new "github.com/blackflagsoftware/agenda/internal/v1/newmember"
	ord "github.com/blackflagsoftware/agenda/internal/v1/ordinance"
	spe "github.com/blackflagsoftware/agenda/internal/v1/speaker"
	vis "github.com/blackflagsoftware/agenda/internal/v1/visitor"
	rel "github.com/blackflagsoftware/agenda/internal/v1/wardbusinessrel"
	sus "github.com/blackflagsoftware/agenda/internal/v1/wardbusinesssus"
	"github.com/jung-kurt/gofpdf"
)

const FONT = "helvetica"

//go:generate mockgen -source=manager.go -destination=mock.go -package=agenda
type (
	DataAgendaAdapter interface {
		Read(*Agenda) error
		ReadAll(*[]Agenda, AgendaParam) (int, error)
		Create(*Agenda) error
		Update(Agenda) error
		Delete(*Agenda) error
		Check(*Agenda) error
	}

	ManagerAgenda struct {
		dataAgenda  DataAgendaAdapter
		auditWriter a.AuditAdapter
	}
)

func NewManagerAgenda(cage DataAgendaAdapter) *ManagerAgenda {
	aw := a.AuditInit()
	return &ManagerAgenda{dataAgenda: cage, auditWriter: aw}
}

func (m *ManagerAgenda) Get(age *Agenda) error {

	return m.dataAgenda.Read(age)
}

func (m *ManagerAgenda) Search(age *[]Agenda, param AgendaParam) (int, error) {
	param.Param.CalculateParam("presiding", map[string]string{"presiding": "presiding", "conducting": "conducting", "organist": "organist", "chorister": "chorister", "newsletter": "newsletter", "intermediate_hymn": "intermediate_hymn", "musical_number": "musical_number", "closing_hymn": "closing_hymn", "invocation": "invocation", "benediction": "benediction", "ward_business": "ward_business", "bishop_business": "bishop_business", "letter_read": "letter_read", "stake_business": "stake_business", "stake": "stake", "new_members": "new_members", "ordinance": "ordinance", "fast_sunday": "fast_sunday", "agenda_published": "agenda_published", "program_published": "program_published"})

	return m.dataAgenda.ReadAll(age, param)
}

func (m *ManagerAgenda) Post(age *Agenda) error {
	ageCheck := &Agenda{Date: age.Date}
	if err := m.dataAgenda.Check(ageCheck); err != nil {
		if err.Error() == "Record exists" {
			return m.dataAgenda.Read(age)
		}
		return err
	}
	if err := m.dataAgenda.Create(age); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *age, AgendaConst, a.KeysToString("Date", age.Date))
	return nil
}

func (m *ManagerAgenda) Patch(ageIn Agenda) error {
	age := &Agenda{Date: ageIn.Date}
	errGet := m.dataAgenda.Read(age)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Presiding
	if ageIn.Presiding.Valid {
		existingValues["presiding"] = age.Presiding.String
		age.Presiding = ageIn.Presiding
	}
	// Conducting
	if ageIn.Conducting.Valid {
		existingValues["conducting"] = age.Conducting.String
		age.Conducting = ageIn.Conducting
	}
	// Organist
	if ageIn.Organist.Valid {
		existingValues["organist"] = age.Organist.String
		age.Organist = ageIn.Organist
	}
	// Chorister
	if ageIn.Chorister.Valid {
		existingValues["chorister"] = age.Chorister.String
		age.Chorister = ageIn.Chorister
	}
	// Newsletter
	if ageIn.Newsletter.Valid {
		existingValues["newsletter"] = age.Newsletter.String
		age.Newsletter = ageIn.Newsletter
	}
	// OpeningHymn
	if ageIn.OpeningHymn.Valid {
		existingValues["opening_hymn"] = age.OpeningHymn.Int64
		age.OpeningHymn = ageIn.OpeningHymn
	}
	// SacramentHymn
	if ageIn.SacramentHymn.Valid {
		existingValues["sacrament_hymn"] = age.SacramentHymn.Int64
		age.SacramentHymn = ageIn.SacramentHymn
	}
	// IntermediateHymn
	if ageIn.IntermediateHymn.Valid {
		existingValues["intermediate_hymn"] = age.IntermediateHymn.Int64
		age.IntermediateHymn = ageIn.IntermediateHymn
	}
	// MusicalNumber
	if ageIn.MusicalNumber.Valid {
		existingValues["musical_number"] = age.MusicalNumber.String
		age.MusicalNumber = ageIn.MusicalNumber
	}
	// ClosingHymn
	if ageIn.ClosingHymn.Valid {
		existingValues["closing_hymn"] = age.ClosingHymn.Int64
		age.ClosingHymn = ageIn.ClosingHymn
	}
	// Invocation
	if ageIn.Invocation.Valid {
		existingValues["invocation"] = age.Invocation.String
		age.Invocation = ageIn.Invocation
	}
	// Benediction
	if ageIn.Benediction.Valid {
		existingValues["benediction"] = age.Benediction.String
		age.Benediction = ageIn.Benediction
	}
	// WardBusiness
	if ageIn.WardBusiness.Valid {
		existingValues["ward_business"] = age.WardBusiness.Bool
		age.WardBusiness = ageIn.WardBusiness
	}
	// BishopBusiness
	if ageIn.BishopBusiness.Valid {
		existingValues["bishop_business"] = age.BishopBusiness.Bool
		age.BishopBusiness = ageIn.BishopBusiness
	}
	// LetterRead
	if ageIn.LetterRead.Valid {
		existingValues["letter_read"] = age.LetterRead.Bool
		age.LetterRead = ageIn.LetterRead
	}
	// StakeBusiness
	if ageIn.StakeBusiness.Valid {
		existingValues["stake_business"] = age.StakeBusiness.Bool
		age.StakeBusiness = ageIn.StakeBusiness
	}
	// Stake
	if ageIn.Stake.Valid {
		existingValues["stake"] = age.Stake.String
		age.Stake = ageIn.Stake
	}
	// NewMembers
	if ageIn.NewMembers.Valid {
		existingValues["new_members"] = age.NewMembers.Bool
		age.NewMembers = ageIn.NewMembers
	}
	// Ordinance
	if ageIn.Ordinance.Valid {
		existingValues["ordinance"] = age.Ordinance.Bool
		age.Ordinance = ageIn.Ordinance
	}
	// Fastsunday
	if ageIn.Fastsunday.Valid {
		existingValues["fast_sunday"] = age.Fastsunday.Bool
		age.Fastsunday = ageIn.Fastsunday
	}
	// AgendaPublished
	if ageIn.AgendaPublished.Valid {
		existingValues["agenda_published"] = age.AgendaPublished.Bool
		age.AgendaPublished = ageIn.AgendaPublished
	}
	// ProgramPublished
	if ageIn.ProgramPublished.Valid {
		existingValues["program_published"] = age.ProgramPublished.Bool
		age.ProgramPublished = ageIn.ProgramPublished
	}
	if err := m.dataAgenda.Update(*age); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *age, AgendaConst, a.KeysToString("Date", age.Date), existingValues)
	return nil
}

func (m *ManagerAgenda) Delete(age *Agenda) error {

	if err := m.dataAgenda.Delete(age); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *age, AgendaConst, a.KeysToString("Date", age.Date))
	return nil
}

func (m *ManagerAgenda) Print(date string) error {
	// get the main agenda record
	// create the pdf struct
	// call each section as needed
	dateTime, errParse := time.Parse("2006-01-02", date)
	if errParse != nil {
		return ae.NewApiError(http.StatusBadRequest, "Date Format", "Date format was not able to parse: "+date, false, nil)
	}
	dateStr := dateTime.Format("01/02/2006")
	agenda := &Agenda{Date: date}
	if err := m.dataAgenda.Read(agenda); err != nil {
		return err
	}
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.AddPage()
	pdf.SetFont(FONT, "", 12)
	pdf.CellFormat(0, 5, "Date: "+dateStr, "", 2, "TC", false, 0, "")
	pdf.SetFont(FONT, "B", 20)
	pdf.CellFormat(0, 14, "Sacrament Meeting Agenda", "", 1, "MC", false, 0, "")
	m.printPersons(pdf, agenda)
	m.printAnnouncements(pdf, agenda)
	hymMgr := m.printOpening(pdf, agenda)
	m.printWardBusiness(pdf, agenda)
	m.printBishopBusiness(pdf, agenda)
	m.printNewMembers(pdf, agenda)
	m.printOrdinance(pdf, agenda)
	m.printSacrament(pdf, agenda, hymMgr)
	m.printProgram(pdf, agenda, hymMgr)
	m.printClosing(pdf, agenda, hymMgr)
	m.printPrayers(pdf)

	path := config.DocumentDir + "/documents/" + date + "-agenda.pdf"
	fmt.Println(path)
	pdf.OutputFileAndClose(path)
	return nil
}

func (m *ManagerAgenda) printPersons(pdf *gofpdf.Fpdf, agenda *Agenda) {
	pdf.SetFont(FONT, "", 12)
	pdf.Cell(20, 5, "Presiding:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(78, 5, agenda.Presiding.String)

	pdf.SetFont(FONT, "", 12)
	pdf.Cell(24, 5, "Conducting:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(68, 5, agenda.Conducting.String)
	pdf.Ln(5)

	pdf.SetFont(FONT, "", 12)
	pdf.Cell(20, 5, "Organist:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(78, 5, agenda.Organist.String)

	pdf.SetFont(FONT, "", 12)
	pdf.Cell(24, 5, "Chorister:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(68, 5, agenda.Chorister.String)
	pdf.Ln(5)

	pdf.SetFont(FONT, "", 12)
	pdf.Cell(38, 5, "Visiting Authorities:")

	visStor := vis.InitStorage()
	visMgr := vis.NewManagerVisitor(visStor)
	visitors := []vis.Visitor{}
	if _, err := visMgr.Search(&visitors, vis.VisitorParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
		fmt.Println("printPerson: getting visitors")
		return
	}
	pdf.SetFont(FONT, "B", 12)
	names := []string{}
	for _, v := range visitors {
		names = append(names, v.Name.String)
	}
	pdf.Cell(0, 5, strings.Join(names, "; "))
	pdf.Ln(8)
	pdf.SetFont(FONT, "", 12)
	pdf.Cell(34, 5, "Welcome Visitors")
	pdf.Ln(8)
}

func (m *ManagerAgenda) printAnnouncements(pdf *gofpdf.Fpdf, agenda *Agenda) {
	pdf.SetFont(FONT, "U", 12)
	pdf.Cell(34, 5, "Announcements:")
	pdf.SetFont(FONT, "I", 11)
	pdf.Cell(116, 5, "\"Please remember to read your programs and weekly email sent to you by ")
	pdf.Ln(5)
	pdf.Cell(34, 5, "")
	pdf.SetFont(FONT, "B", 11)
	pdf.Cell(48, 5, agenda.Newsletter.String)
	pdf.SetFont(FONT, "I", 11)
	pdf.Cell(0, 5, "for ward, stake, and church announcements.\"")
	pdf.Ln(8)

	annStor := ann.InitStorage()
	annMgr := ann.NewManagerAnnouncement(annStor)
	announcements := []ann.Announcement{}
	if _, err := annMgr.Search(&announcements, ann.AnnouncementParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
		fmt.Println("printAnnouncements: getting announcements")
		return
	}
	printOnce := true
	for _, a := range announcements {
		if a.Pulpit.Bool {
			if printOnce {
				pdf.SetFont(FONT, "U", 12)
				pdf.Cell(0, 5, "Special Announcements:")
				pdf.SetFont(FONT, "B", 11)
				printOnce = false
			}
			pdf.Ln(5)
			pdf.Cell(4, 5, "")
			pdf.MultiCell(0, 5, "- "+a.Message.String, "", "", false)
		}
	}
	pdf.Ln(8)
}

func (m *ManagerAgenda) printOpening(pdf *gofpdf.Fpdf, agenda *Agenda) (hymMgr *hym.ManagerHymn) {
	hymStor := hym.InitStorage()
	hymMgr = hym.NewManagerHymn(hymStor)
	hymn := hym.Hymn{Id: int(agenda.OpeningHymn.Int64)}
	if err := hymMgr.Get(&hymn); err != nil {
		fmt.Println("printOpening: getting opening hymn")
		return
	}
	pdf.SetFont(FONT, "", 12)
	pdf.Cell(30, 5, "Opening Hymn:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(78, 5, fmt.Sprintf("#%d - %s", hymn.Id, hymn.Name.String))

	pdf.SetFont(FONT, "", 12)
	pdf.Cell(22, 5, "Invocation:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(0, 5, agenda.Invocation.String)
	pdf.Ln(8)
	return
}

func (m *ManagerAgenda) printWardBusiness(pdf *gofpdf.Fpdf, agenda *Agenda) {
	if agenda.WardBusiness.Bool {
		pdf.SetFont(FONT, "U", 12)
		pdf.Cell(30, 5, "Ward Business:")

		relStor := rel.InitStorage()
		relMgr := rel.NewManagerWardBusinessRel(relStor)
		releases := []rel.WardBusinessRel{}
		if _, err := relMgr.Search(&releases, rel.WardBusinessRelParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
			fmt.Println("printWardBusiness: getting releases")
			return
		}
		if len(releases) > 0 {
			pdf.Ln(5)
			pdf.SetFont(FONT, "B", 12)
			pdf.Cell(4, 5, "")
			pdf.Cell(22, 5, "Releases:")
			pdf.SetFont(FONT, "I", 11)
			pdf.Cell(0, 5, "\"We have released the following individuals.\"")
			pdf.SetFont(FONT, "", 12)
			for _, r := range releases {
				pdf.Ln(5)
				pdf.Cell(4, 5, "")
				pdf.SetFont(FONT, "B", 12)
				pdf.Cell(54, 5, r.Name.String)
				pdf.SetFont(FONT, "", 12)
				pdf.Cell(44, 5, "has been released as")
				pdf.SetFont(FONT, "B", 12)
				pdf.Cell(0, 5, r.Calling.String)
			}
			pdf.Ln(8)
			pdf.Cell(4, 5, "")
			pdf.SetFont(FONT, "I", 11)
			pdf.MultiCell(0, 5, "\"We propose they be given thanks for their service.  Those who wish to express appreciation may do so by the uplifted hand\"", "", "", false)
			pdf.Ln(6)
		}
		susStor := sus.InitStorage()
		susMgr := sus.NewManagerWardBusinessSus(susStor)
		sustainings := []sus.WardBusinessSus{}
		if _, err := susMgr.Search(&sustainings, sus.WardBusinessSusParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
			fmt.Println("printWardBusiness: getting sustainings")
			return
		}
		if len(sustainings) > 0 {
			pdf.SetFont(FONT, "B", 12)
			pdf.Cell(4, 5, "")
			pdf.Cell(26, 5, "Sustainings:")
			pdf.SetFont(FONT, "I", 11)
			pdf.Cell(0, 5, "\"The following people have been called.  If they are present, would they please stand.\"")
			pdf.SetFont(FONT, "", 12)
			for _, s := range sustainings {
				pdf.Ln(5)
				pdf.Cell(4, 5, "")
				pdf.SetFont(FONT, "B", 12)
				pdf.Cell(54, 5, s.Name.String)
				pdf.SetFont(FONT, "", 12)
				pdf.Cell(40, 5, "has been called as")
				pdf.SetFont(FONT, "B", 12)
				pdf.Cell(0, 5, s.Calling.String)
			}
			pdf.Ln(8)
			pdf.Cell(4, 5, "")
			pdf.SetFont(FONT, "I", 11)
			pdf.Cell(0, 5, "\"We propose that (He, She, They) be sustained.\"")
			pdf.Ln(5)
			pdf.Cell(4, 5, "")
			pdf.Cell(0, 5, "\"Those in favor may manifest it by the uplifted hand.\" (Pause) \"Those opposed, if any, may manifest it.\"")
			pdf.Ln(8)
			pdf.Cell(4, 5, "")
			pdf.MultiCell(0, 5, "\"We invite all those who have been sustained today to come to the Bishop's office immediately after the block to be set apart\"", "", "", false)
			pdf.Ln(6)
		}
	}
}

func (m *ManagerAgenda) printBishopBusiness(pdf *gofpdf.Fpdf, agenda *Agenda) {
	if agenda.BishopBusiness.Bool {
		pdf.SetFont(FONT, "U", 12)
		pdf.Cell(30, 5, "Bishop Business:")

		bisStor := bis.InitStorage()
		bisMgr := bis.NewManagerBishopBusiness(bisStor)
		bishop := []bis.BishopBusiness{}
		if _, err := bisMgr.Search(&bishop, bis.BishopBusinessParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
			fmt.Println("printBishopBusiness: getting messages")
			return
		}
		pdf.SetFont(FONT, "B", 12)
		for _, b := range bishop {
			pdf.Ln(5)
			pdf.Cell(4, 5, "")
			pdf.Cell(0, 5, "- "+b.Message.String)
		}
		pdf.Ln(8)
	}
	if agenda.LetterRead.Bool {
		pdf.SetFont(FONT, "U", 12)
		pdf.Cell(32, 5, "Letter To Read:")
		pdf.SetFont(FONT, "B", 12)
		pdf.Cell(0, 5, "Yes")
		pdf.Ln(8)
	}
	if agenda.StakeBusiness.Bool {
		pdf.SetFont(FONT, "U", 12)
		pdf.Cell(34, 5, "Stake Business:")
		pdf.SetFont(FONT, "B", 12)
		pdf.Cell(0, 5, agenda.Stake.String)
		pdf.Ln(8)
	}
}

func (m *ManagerAgenda) printNewMembers(pdf *gofpdf.Fpdf, agenda *Agenda) {
	if agenda.NewMembers.Bool {
		pdf.SetFont(FONT, "U", 12)
		pdf.Cell(0, 5, "Membership Records Received:")
		newStor := new.InitStorage()
		newMgr := new.NewManagerNewMember(newStor)
		newMembers := []new.NewMember{}
		if _, err := newMgr.Search(&newMembers, new.NewMemberParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
			fmt.Println("printNewMembers: getting newmembers")
			return
		}
		pdf.Ln(5)
		pdf.Cell(4, 5, "")
		pdf.SetFont(FONT, "I", 11)
		pdf.Cell(0, 5, "\"We have received the membership records for the following individuals. Please stand as your name is read.\"")
		pdf.Ln(3)
		for _, n := range newMembers {
			pdf.Ln(5)
			pdf.Cell(4, 5, "")
			pdf.SetFont(FONT, "B", 12)
			pdf.Cellf(0, 5, "%s; %s", n.FamilyName.String, n.Names.String)
		}
		pdf.Ln(8)
		pdf.Cell(4, 5, "")
		pdf.SetFont(FONT, "I", 11)
		pdf.Cell(0, 5, "\"All who can join us in welcoming these new ward members, please do so by the uplifted hand.\"")
		pdf.Ln(8)
	}
}

func (m *ManagerAgenda) printOrdinance(pdf *gofpdf.Fpdf, agenda *Agenda) {
	if agenda.Ordinance.Bool {
		pdf.SetFont(FONT, "U", 12)
		pdf.Cell(0, 5, "Baptism Confirmations, Blessings")
		ordStor := ord.InitStorage()
		ordMgr := ord.NewManagerOrdinance(ordStor)
		ordinances := []ord.Ordinance{}
		if _, err := ordMgr.Search(&ordinances, ord.OrdinanceParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
			fmt.Println("printOrdinance: getting ordinances")
			return
		}
		pdf.Ln(5)
		for _, n := range ordinances {
			if n.Confirmations.String != "" {
				pdf.Cell(4, 5, "")
				pdf.SetFont(FONT, "", 12)
				pdf.Cell(20, 5, "Baptisms:")
				pdf.SetFont(FONT, "B", 12)
				pdf.Cell(0, 5, n.Confirmations.String)
			}
			if n.Blessings.String != "" {
				pdf.Ln(5)
				pdf.Cell(4, 5, "")
				pdf.SetFont(FONT, "", 12)
				pdf.Cell(20, 5, "Blessings:")
				pdf.SetFont(FONT, "B", 12)
				pdf.Cell(0, 5, n.Blessings.String)
			}
		}
		pdf.Ln(8)
	}
}

func (m *ManagerAgenda) printSacrament(pdf *gofpdf.Fpdf, agenda *Agenda, hymMgr *hym.ManagerHymn) {
	hymn := hym.Hymn{Id: int(agenda.SacramentHymn.Int64)}
	if err := hymMgr.Get(&hymn); err != nil {
		fmt.Println("printSacrament: getting sacrament hymn")
		return
	}
	pdf.SetFont(FONT, "", 12)
	pdf.Cell(34, 5, "Sacrament Hymn:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(0, 5, fmt.Sprintf("#%d - %s", hymn.Id, hymn.Name.String))
	pdf.Ln(5)

	pdf.SetFont(FONT, "", 12)
	pdf.Cell(0, 5, "Then the Administration of the Sacrament")
	pdf.Ln(5)
	pdf.SetFont(FONT, "I", 11)
	pdf.Cell(0, 5, "\"We thank you for your reverence during the sacrament.\"")
	pdf.Ln(8)
	return
}

func (m *ManagerAgenda) printProgram(pdf *gofpdf.Fpdf, agenda *Agenda, hymMgr *hym.ManagerHymn) {
	pdf.SetFont(FONT, "U", 12)
	pdf.Cell(0, 5, "Program:")
	pdf.Ln(5)
	if agenda.Fastsunday.Bool {
		pdf.SetFont(FONT, "", 12)
		pdf.Cell(4, 5, "")
		pdf.Cell(0, 5, "Fast and Testimony Meeting Agenda")
		pdf.Ln(5)
		pdf.Cell(4, 5, "")
		pdf.Cell(0, 5, "1) Bare your testimony")
		pdf.Ln(5)
		pdf.Cell(4, 5, "")
		pdf.Cell(0, 5, "2) Turn the time over to the congregation \"We will end baring testimonies 5 minutes to the hour.\"")
		pdf.Ln(5)
		pdf.Cell(4, 5, "")
		pdf.Cell(0, 5, "3) Thank all who shared their testimones")
		pdf.Ln(5)
		pdf.Cell(4, 5, "")
		pdf.Cell(0, 5, "4) Announce closing hymn and benediction")
		pdf.Ln(8)
		return
	}
	speakerPosition := 1
	positionMapping := map[int]string{1: "1st", 2: "2nd", 3: "3rd", 4: "4th", 5: "5th"}

	speStor := spe.InitStorage()
	speMgr := spe.NewManagerSpeaker(speStor)
	speakers := []spe.Speaker{}
	if _, err := speMgr.Search(&speakers, spe.SpeakerParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
		fmt.Println("printProgram: getting speakers")
		return
	}
	positionStr := ""
	speaker := ""
	for _, s := range speakers {
		foundOther := false
		if s.Name.String == "Intermediate Hymn" {
			hymn := hym.Hymn{Id: int(agenda.IntermediateHymn.Int64)}
			if err := hymMgr.Get(&hymn); err != nil {
				fmt.Println("printProgram: getting intermediate hymn")
				return
			}
			positionStr = "Int. Hymn"
			speaker = fmt.Sprintf("#%d - %s", hymn.Id, hymn.Name.String)
			foundOther = true
		}
		if s.Name.String == "Musical Number" {
			positionStr = "Mus. Number"
			speaker = agenda.MusicalNumber.String
			foundOther = true
		}
		if !foundOther {
			positionStr = positionMapping[speakerPosition]
			speakerPosition++
			speaker = s.Name.String
		}
		pdf.SetFont(FONT, "", 12)
		pdf.Cell(4, 5, "")
		pdf.Cell(20, 5, positionStr)
		pdf.Cell(0, 5, speaker)
		pdf.Ln(5)
	}
	pdf.Ln(5)
}

func (m *ManagerAgenda) printClosing(pdf *gofpdf.Fpdf, agenda *Agenda, hymMgr *hym.ManagerHymn) {
	hymStor := hym.InitStorage()
	hymMgr = hym.NewManagerHymn(hymStor)
	hymn := hym.Hymn{Id: int(agenda.ClosingHymn.Int64)}
	if err := hymMgr.Get(&hymn); err != nil {
		fmt.Println("printClosing: getting closing hymn")
		return
	}
	pdf.SetFont(FONT, "", 12)
	pdf.Cell(30, 5, "Closing Hymn:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(78, 5, fmt.Sprintf("#%d - %s", hymn.Id, hymn.Name.String))

	pdf.SetFont(FONT, "", 12)
	pdf.Cell(24, 5, "Benediction:")
	pdf.SetFont(FONT, "B", 12)
	pdf.Cell(0, 5, agenda.Benediction.String)
	pdf.Ln(12)
	return
}

func (m *ManagerAgenda) printPrayers(pdf *gofpdf.Fpdf) {
	pdf.SetFont(FONT, "BU", 14)
	pdf.Cell(0, 5, "Blessing of the Bread")
	pdf.Ln(6)
	pdf.SetFont(FONT, "", 14)
	pdf.MultiCell(0, 5, "\"O God, the Eternal Father, we ask thee in the name of thy Son, Jesus Christ, to bless and sanctify this bread to the souls of all those who partake of it, that they may eat in remembrance of the body of thy Son, and witness unto thee, O God, the Eternal Father, that they are willing to take upon them the name of thy Son, and always remember him and keep his commandments which he has given them; that they may always have his Spirit to be with them. Amen.\"", "", "", false)
	pdf.Ln(8)
	pdf.SetFont(FONT, "BU", 14)
	pdf.Cell(0, 5, "Blessing of the Water")
	pdf.Ln(6)
	pdf.SetFont(FONT, "", 14)
	pdf.MultiCell(0, 5, "\"O God, the Eternal Father, we ask thee in the name of thy Son, Jesus Christ, to bless and sanctify this water to the souls of all those who drink of it, that they may do it in remembrance of the blood of thy Son, which was shed for them; that they may witness unto thee, O God, the Eternal Father, that they do always remember him, that they may have his Spirit to be with them. Amen.\"", "", "", false)
}

func (m *ManagerAgenda) Publish(date string) error {
	dateTime, errParse := time.Parse("2006-01-02", date)
	if errParse != nil {
		return ae.NewApiError(http.StatusBadRequest, "Date Format", "Date format was not able to parse: "+date, false, nil)
	}
	dateStr := dateTime.Format("01/02/2006")
	agenda := &Agenda{Date: date}
	if err := m.dataAgenda.Read(agenda); err != nil {
		return err
	}
	pdfP := gofpdf.New("P", "mm", "Letter", "")
	pdfP.AddPage()
	pdfL := gofpdf.New("L", "mm", "Letter", "")
	pdfL.AddPage()
	m.printProgramHeader(dateStr, pdfP, pdfL)
	m.printProgramPersons(pdfP, pdfL, agenda)
	m.printProgramProgram(pdfP, pdfL, agenda)
	m.printProgramAnnouncements(pdfP, pdfL, agenda)

	pdfP.OutputFileAndClose(config.DocumentDir + "/documents/" + date + "-qr.pdf")
	pdfL.OutputFileAndClose(config.DocumentDir + "/documents/" + date + "-program.pdf")
	return nil
}

func (m *ManagerAgenda) printProgramHeader(dateStr string, pdfP *gofpdf.Fpdf, pdfL *gofpdf.Fpdf) {
	pdfP.SetFont(FONT, "B", 16)
	pdfL.SetFont(FONT, "B", 16)
	pdfP.CellFormat(0, 14, "River Ridge 11th Ward Sacrament Meeting", "", 1, "MC", false, 0, "")
	pdfL.CellFormat(119, 14, "River Ridge 11th Ward Sacrament Meeting", "", 0, "MC", false, 0, "")
	pdfL.Cell(20, 5, "")
	pdfL.CellFormat(0, 14, "River Ridge 11th Ward Sacrament Meeting", "", 1, "MC", false, 0, "")
	pdfP.SetFont(FONT, "", 12)
	pdfL.SetFont(FONT, "", 12)
	pdfP.CellFormat(0, 5, dateStr, "", 2, "TC", false, 0, "")
	pdfL.CellFormat(119, 5, dateStr, "", 0, "TC", false, 0, "")
	pdfL.Cell(10, 5, "")
	pdfL.CellFormat(0, 5, dateStr, "", 2, "TC", false, 0, "")
}

func (m *ManagerAgenda) printProgramPersons(pdfP *gofpdf.Fpdf, pdfL *gofpdf.Fpdf, agenda *Agenda) {
	pdfP.SetFont(FONT, "", 12)
	pdfL.SetFont(FONT, "", 12)
	pdfP.Ln(10)
	pdfL.Ln(10)
	pdfP.Cell(24, 5, "Presiding:")
	pdfP.Cell(0, 5, agenda.Presiding.String)
	pdfP.Ln(5)
	pdfP.Cell(24, 5, "Conducting:")
	pdfP.Cell(0, 5, agenda.Conducting.String)
	pdfP.Ln(5)
	pdfP.Cell(24, 5, "Organist")
	pdfP.Cell(0, 5, agenda.Organist.String)
	pdfP.Ln(5)
	pdfP.Cell(24, 5, "Chorister")
	pdfP.Cell(0, 5, agenda.Chorister.String)
	pdfP.Ln(5)
	pdfL.Cell(24, 5, "Presiding:")
	pdfL.Cell(95, 5, agenda.Presiding.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(24, 5, "Presiding:")
	pdfL.Cell(0, 5, agenda.Presiding.String)
	pdfL.Ln(5)
	pdfL.Cell(24, 5, "Conducting:")
	pdfL.Cell(95, 5, agenda.Conducting.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(24, 5, "Conducting:")
	pdfL.Cell(0, 5, agenda.Conducting.String)
	pdfL.Ln(5)
	pdfL.Cell(24, 5, "Organist")
	pdfL.Cell(95, 5, agenda.Organist.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(24, 5, "Organist")
	pdfL.Cell(0, 5, agenda.Organist.String)
	pdfL.Ln(5)
	pdfL.Cell(24, 5, "Chorister")
	pdfL.Cell(95, 5, agenda.Chorister.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(24, 5, "Chorister")
	pdfL.Cell(0, 5, agenda.Chorister.String)
	pdfP.Ln(8)
	pdfL.Ln(8)
}

func (m *ManagerAgenda) printProgramAnnouncements(pdfP *gofpdf.Fpdf, pdfL *gofpdf.Fpdf, agenda *Agenda) {
	annStor := ann.InitStorage()
	annMgr := ann.NewManagerAnnouncement(annStor)
	announcements := []ann.Announcement{}
	if _, err := annMgr.Search(&announcements, ann.AnnouncementParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
		fmt.Println("printProgramAnnouncements: getting announcements")
		return
	}
	if len(announcements) > 0 {
		pdfP.SetFont(FONT, "U", 12)
		pdfL.SetFont(FONT, "U", 12)
		pdfP.Cell(34, 5, "Announcements:")
		pdfL.Cell(119, 5, "Announcements:")
		pdfL.Cell(20, 5, "")
		pdfL.Cell(0, 5, "Announcements:")
		pdfP.SetFont(FONT, "", 12)
		pdfL.SetFont(FONT, "", 12)

		for _, a := range announcements {
			pdfP.Ln(6)
			pdfL.Ln(6)
			pdfP.Cell(4, 5, "")
			pdfP.MultiCell(0, 5, a.Message.String, "", "", false)
			resetY := pdfL.GetY()
			pdfL.Cell(4, 5, "")
			pdfL.MultiCell(114, 5, a.Message.String, "", "", false)
			pdfL.SetXY(153, resetY)
			pdfL.MultiCell(114, 5, a.Message.String, "", "", false)
		}
	}
	pdfP.Ln(10)
	pdfL.Ln(10)
}

func (m *ManagerAgenda) printProgramProgram(pdfP *gofpdf.Fpdf, pdfL *gofpdf.Fpdf, agenda *Agenda) {
	hymStor := hym.InitStorage()
	hymMgr := hym.NewManagerHymn(hymStor)
	hymnOpening := hym.Hymn{Id: int(agenda.OpeningHymn.Int64)}
	if err := hymMgr.Get(&hymnOpening); err != nil {
		fmt.Println("printOpening: getting opening hymn")
		return
	}
	pdfP.SetFont(FONT, "U", 12)
	pdfL.SetFont(FONT, "U", 12)
	pdfP.Cell(0, 5, "Program:")
	pdfL.Cell(119, 5, "Program:")
	pdfL.Cell(20, 5, "")
	pdfL.Cell(0, 5, "Program:")
	pdfP.Ln(6)
	pdfL.Ln(6)
	pdfP.SetFont(FONT, "", 12)
	pdfL.SetFont(FONT, "", 12)
	pdfP.Cell(4, 5, "")
	pdfP.Cell(38, 5, "Opening Hymn:")
	pdfP.Cellf(81, 5, "%d - %s", hymnOpening.Id, hymnOpening.Name.String)
	pdfL.Cell(4, 5, "")
	pdfL.Cell(38, 5, "Opening Hymn:")
	pdfL.Cellf(81, 5, "%d - %s", hymnOpening.Id, hymnOpening.Name.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(38, 5, "Opening Hymn:")
	pdfL.Cellf(0, 5, "%d - %s", hymnOpening.Id, hymnOpening.Name.String)
	pdfP.Ln(5)
	pdfL.Ln(5)
	pdfP.Cell(4, 5, "")
	pdfP.Cell(38, 5, "Invocation:")
	pdfP.Cell(0, 5, agenda.Invocation.String)
	pdfL.Cell(4, 5, "")
	pdfL.Cell(38, 5, "Invocation:")
	pdfL.Cell(81, 5, agenda.Invocation.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(38, 5, "Invocation:")
	pdfL.Cell(0, 5, agenda.Invocation.String)
	pdfP.Ln(5)
	pdfL.Ln(5)
	if agenda.WardBusiness.Bool {
		pdfP.Ln(2)
		pdfL.Ln(2)
		pdfP.Cell(4, 5, "")
		pdfP.Cell(0, 5, "Ward Business")
		pdfL.Cell(4, 5, "")
		pdfL.Cell(115, 5, "Ward Business")
		pdfL.Cell(24, 5, "")
		pdfL.Cell(0, 5, "Ward Business")
		pdfP.Ln(6)
		pdfL.Ln(7)
	}
	hymnSacrament := hym.Hymn{Id: int(agenda.SacramentHymn.Int64)}
	if err := hymMgr.Get(&hymnSacrament); err != nil {
		fmt.Println("printProgramProgram: getting sacrament hymn")
		return
	}
	pdfP.Cell(4, 5, "")
	pdfP.Cell(38, 5, "Sacrament Hymn:")
	pdfP.Cellf(81, 5, "%d - %s", hymnSacrament.Id, hymnSacrament.Name.String)
	pdfL.Cell(4, 5, "")
	pdfL.Cell(38, 5, "Sacrament Hymn:")
	pdfL.Cellf(81, 5, "%d - %s", hymnSacrament.Id, hymnSacrament.Name.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(38, 5, "Sacrament Hymn:")
	pdfL.Cellf(0, 5, "%d - %s", hymnSacrament.Id, hymnSacrament.Name.String)
	pdfP.Ln(5)
	pdfL.Ln(5)
	pdfP.Cell(4, 5, "")
	pdfP.Cell(0, 5, "Administration of the Sacrament")
	pdfL.Cell(4, 5, "")
	pdfL.Cell(115, 5, "Administration of the Sacrament")
	pdfL.Cell(24, 5, "")
	pdfL.Cell(0, 5, "Administration of the Sacrament")
	pdfP.Ln(7)
	pdfL.Ln(7)
	if agenda.Fastsunday.Bool {
		pdfP.SetFont(FONT, "", 12)
		pdfL.SetFont(FONT, "", 12)
		pdfP.Ln(3)
		pdfP.Cell(4, 5, "")
		pdfP.Cell(0, 5, "Testimonies")
		pdfP.Ln(8)
		pdfL.Ln(3)
		pdfL.Cell(4, 5, "")
		pdfL.Cell(115, 5, "Testimonies")
		pdfL.Cell(24, 5, "")
		pdfL.Cell(0, 5, "Testimonies")
		pdfL.Ln(7)
	} else {
		speakerPosition := 1
		positionMapping := map[int]string{1: "1st", 2: "2nd", 3: "3rd", 4: "4th", 5: "5th"}

		speStor := spe.InitStorage()
		speMgr := spe.NewManagerSpeaker(speStor)
		speakers := []spe.Speaker{}
		if _, err := speMgr.Search(&speakers, spe.SpeakerParam{Param: util.Param{Search: []util.ParamSearch{{Column: "date", Value: agenda.Date, Compare: "="}}}}); err != nil {
			fmt.Println("printProgramProgram: getting speakers")
			return
		}
		positionStr := ""
		speaker := ""
		for _, s := range speakers {
			foundOther := false
			if s.Name.String == "Intermediate Hymn" {
				hymn := hym.Hymn{Id: int(agenda.IntermediateHymn.Int64)}
				if err := hymMgr.Get(&hymn); err != nil {
					fmt.Println("printProgramProgram: getting intermediate hymn")
					return
				}
				positionStr = "Intermediate Hymn:"
				speaker = fmt.Sprintf("%d - %s", hymn.Id, hymn.Name.String)
				foundOther = true
			}
			if s.Name.String == "Musical Number" {
				positionStr = "Musical Number:"
				speaker = agenda.MusicalNumber.String
				foundOther = true
			}
			if !foundOther {
				positionStr = positionMapping[speakerPosition] + " Speaker:"
				speakerPosition++
				speaker = s.Name.String
			}
			pdfP.SetFont(FONT, "", 12)
			pdfL.SetFont(FONT, "", 12)
			pdfP.Cell(4, 5, "")
			pdfP.Cell(38, 5, positionStr)
			pdfP.Cell(0, 5, speaker)
			pdfP.Ln(5)
			pdfL.Cell(4, 5, "")
			pdfL.Cell(38, 5, positionStr)
			pdfL.Cell(81, 5, speaker)
			pdfL.Cell(20, 5, "")
			pdfL.Cell(38, 5, positionStr)
			pdfL.Cell(0, 5, speaker)
			pdfL.Ln(5)
		}
	}
	hymnClosing := hym.Hymn{Id: int(agenda.ClosingHymn.Int64)}
	if err := hymMgr.Get(&hymnClosing); err != nil {
		fmt.Println("printProgramProgram: getting closing hymn")
		return
	}
	pdfP.Ln(2)
	pdfL.Ln(2)
	pdfP.Cell(4, 5, "")
	pdfP.Cell(38, 5, "Closing Hymn:")
	pdfP.Cellf(81, 5, "%d - %s", hymnClosing.Id, hymnClosing.Name.String)
	pdfL.Cell(4, 5, "")
	pdfL.Cell(38, 5, "Closing Hymn:")
	pdfL.Cellf(81, 5, "%d - %s", hymnClosing.Id, hymnClosing.Name.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(38, 5, "Closing Hymn:")
	pdfL.Cellf(0, 5, "%d - %s", hymnClosing.Id, hymnClosing.Name.String)
	pdfP.Ln(5)
	pdfL.Ln(5)
	pdfP.Cell(4, 5, "")
	pdfP.Cell(38, 5, "Benediction:")
	pdfP.Cell(0, 5, agenda.Benediction.String)
	pdfL.Cell(4, 5, "")
	pdfL.Cell(38, 5, "Benediction:")
	pdfL.Cell(81, 5, agenda.Benediction.String)
	pdfL.Cell(20, 5, "")
	pdfL.Cell(38, 5, "Benediction:")
	pdfL.Cell(0, 5, agenda.Benediction.String)
	pdfP.Ln(8)
	pdfL.Ln(8)
}
