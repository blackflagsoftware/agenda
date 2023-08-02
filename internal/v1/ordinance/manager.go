package ordinance

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=ordinance
type (
	DataOrdinanceAdapter interface {
		Read(*Ordinance) error
		ReadAll(*[]Ordinance, OrdinanceParam) (int, error)
		Create(*Ordinance) error
		Update(Ordinance) error
		Delete(*Ordinance) error
	}

	ManagerOrdinance struct {
		dataOrdinance DataOrdinanceAdapter
		auditWriter   a.AuditAdapter
	}
)

func NewManagerOrdinance(cord DataOrdinanceAdapter) *ManagerOrdinance {
	aw := a.AuditInit()
	return &ManagerOrdinance{dataOrdinance: cord, auditWriter: aw}
}

func (m *ManagerOrdinance) Get(ord *Ordinance) error {
	return m.dataOrdinance.Read(ord)
}

func (m *ManagerOrdinance) Search(ord *[]Ordinance, param OrdinanceParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "confirmations": "confirmations", "blessings": "blessings"})

	return m.dataOrdinance.ReadAll(ord, param)
}

func (m *ManagerOrdinance) Post(ord *Ordinance) error {
	if ord.Id == 0 {
		if err := m.dataOrdinance.Create(ord); err != nil {
			return nil
		}
		go a.AuditCreate(m.auditWriter, *ord, OrdinanceConst, a.KeysToString("Id", ord.Id))
		return nil
	}
	// upsert
	return m.Patch(*ord)
}

func (m *ManagerOrdinance) Patch(ordIn Ordinance) error {
	ord := &Ordinance{Id: ordIn.Id}
	errGet := m.dataOrdinance.Read(ord)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if ordIn.Date.Valid {
		existingValues["date"] = ord.Date.String
		ord.Date = ordIn.Date
	}
	// Confirmations
	if ordIn.Confirmations.Valid {
		existingValues["confirmations"] = ord.Confirmations.String
		ord.Confirmations = ordIn.Confirmations
	}
	// Blessings
	if ordIn.Blessings.Valid {
		existingValues["blessings"] = ord.Blessings.String
		ord.Blessings = ordIn.Blessings
	}
	if err := m.dataOrdinance.Update(*ord); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *ord, OrdinanceConst, a.KeysToString("Id", ord.Id), existingValues)
	return nil
}

func (m *ManagerOrdinance) Delete(ord *Ordinance) error {

	if err := m.dataOrdinance.Delete(ord); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *ord, OrdinanceConst, a.KeysToString("Id", ord.Id))
	return nil
}
