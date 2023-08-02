package wardbusinesssus

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=wardbusinesssus
type (
	DataWardBusinessSusAdapter interface {
		Read(*WardBusinessSus) error
		ReadAll(*[]WardBusinessSus, WardBusinessSusParam) (int, error)
		Create(*WardBusinessSus) error
		Update(WardBusinessSus) error
		Delete(*WardBusinessSus) error
	}

	ManagerWardBusinessSus struct {
		dataWardBusinessSus DataWardBusinessSusAdapter
		auditWriter         a.AuditAdapter
	}
)

func NewManagerWardBusinessSus(cwa DataWardBusinessSusAdapter) *ManagerWardBusinessSus {
	aw := a.AuditInit()
	return &ManagerWardBusinessSus{dataWardBusinessSus: cwa, auditWriter: aw}
}

func (m *ManagerWardBusinessSus) Get(wa *WardBusinessSus) error {

	return m.dataWardBusinessSus.Read(wa)
}

func (m *ManagerWardBusinessSus) Search(wa *[]WardBusinessSus, param WardBusinessSusParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "name": "name", "calling": "calling"})

	return m.dataWardBusinessSus.ReadAll(wa, param)
}

func (m *ManagerWardBusinessSus) Post(wa *WardBusinessSus) error {

	if err := m.dataWardBusinessSus.Create(wa); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *wa, WardBusinessSusConst, a.KeysToString("Id", wa.Id))
	return nil
}

func (m *ManagerWardBusinessSus) Patch(waIn WardBusinessSus) error {
	wa := &WardBusinessSus{Id: waIn.Id}
	errGet := m.dataWardBusinessSus.Read(wa)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if waIn.Date.Valid {
		existingValues["date"] = wa.Date.String
		wa.Date = waIn.Date
	}
	// Name
	if waIn.Name.Valid {
		existingValues["name"] = wa.Name.String
		wa.Name = waIn.Name
	}
	// Calling
	if waIn.Calling.Valid {
		existingValues["calling"] = wa.Calling.String
		wa.Calling = waIn.Calling
	}
	if err := m.dataWardBusinessSus.Update(*wa); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *wa, WardBusinessSusConst, a.KeysToString("Id", wa.Id), existingValues)
	return nil
}

func (m *ManagerWardBusinessSus) Delete(wa *WardBusinessSus) error {

	if err := m.dataWardBusinessSus.Delete(wa); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *wa, WardBusinessSusConst, a.KeysToString("Id", wa.Id))
	return nil
}
