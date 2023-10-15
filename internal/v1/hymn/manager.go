package hymn

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=hymn
type (
	DataHymnAdapter interface {
		Read(*Hymn) error
		ReadAll(*[]Hymn, HymnParam) (int, error)
		Create(*Hymn) error
		Update(Hymn) error
		Delete(*Hymn) error
	}

	ManagerHymn struct {
		dataHymn    DataHymnAdapter
		auditWriter a.AuditAdapter
	}
)

func NewManagerHymn(chym DataHymnAdapter) *ManagerHymn {
	aw := a.AuditInit()
	return &ManagerHymn{dataHymn: chym, auditWriter: aw}
}

func (m *ManagerHymn) Get(hym *Hymn) error {

	return m.dataHymn.Read(hym)
}

func (m *ManagerHymn) Search(hym *[]Hymn, param HymnParam) (int, error) {
	param.Param.CalculateParam("name", map[string]string{"name": "name"})

	return m.dataHymn.ReadAll(hym, param)
}

func (m *ManagerHymn) Post(hym *Hymn) error {

	if err := m.dataHymn.Create(hym); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *hym, HymnConst, a.KeysToString("Id", hym.Id))
	return nil
}

func (m *ManagerHymn) Patch(hymIn Hymn) error {
	hym := &Hymn{Id: hymIn.Id}
	errGet := m.dataHymn.Read(hym)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Name
	if hymIn.Name.Valid {
		existingValues["name"] = hym.Name.String
		hym.Name = hymIn.Name
	}
	// Name
	if hymIn.PdfLink.Valid {
		existingValues["pdf_name"] = hym.PdfLink.String
		hym.Name = hymIn.Name
	}
	if err := m.dataHymn.Update(*hym); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *hym, HymnConst, a.KeysToString("Id", hym.Id), existingValues)
	return nil
}

func (m *ManagerHymn) Delete(hym *Hymn) error {

	if err := m.dataHymn.Delete(hym); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *hym, HymnConst, a.KeysToString("Id", hym.Id))
	return nil
}
