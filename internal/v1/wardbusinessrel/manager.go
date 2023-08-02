package wardbusinessrel

import (
	"fmt"

	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=wardbusinessrel
type (
	DataWardBusinessRelAdapter interface {
		Read(*WardBusinessRel) error
		ReadAll(*[]WardBusinessRel, WardBusinessRelParam) (int, error)
		Create(*WardBusinessRel) error
		Update(WardBusinessRel) error
		Delete(*WardBusinessRel) error
	}

	ManagerWardBusinessRel struct {
		dataWardBusinessRel DataWardBusinessRelAdapter
		auditWriter         a.AuditAdapter
	}
)

func NewManagerWardBusinessRel(cwar DataWardBusinessRelAdapter) *ManagerWardBusinessRel {
	aw := a.AuditInit()
	return &ManagerWardBusinessRel{dataWardBusinessRel: cwar, auditWriter: aw}
}

func (m *ManagerWardBusinessRel) Get(war *WardBusinessRel) error {

	return m.dataWardBusinessRel.Read(war)
}

func (m *ManagerWardBusinessRel) Search(war *[]WardBusinessRel, param WardBusinessRelParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "name": "name", "calling": "calling"})

	return m.dataWardBusinessRel.ReadAll(war, param)
}

func (m *ManagerWardBusinessRel) Post(war *WardBusinessRel) error {
	fmt.Printf("body: %+v\n", *war)
	if err := m.dataWardBusinessRel.Create(war); err != nil {
		return nil
	}
	id := a.KeysToString("Id", war.Id)
	fmt.Println("id", id)
	go a.AuditCreate(m.auditWriter, *war, WardBusinessRelConst, id)
	return nil
}

func (m *ManagerWardBusinessRel) Patch(warIn WardBusinessRel) error {
	war := &WardBusinessRel{Id: warIn.Id}
	errGet := m.dataWardBusinessRel.Read(war)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if warIn.Date.Valid {
		existingValues["date"] = war.Date.String
		war.Date = warIn.Date
	}
	// Name
	if warIn.Name.Valid {
		existingValues["name"] = war.Name.String
		war.Name = warIn.Name
	}
	// Calling
	if warIn.Calling.Valid {
		existingValues["calling"] = war.Calling.String
		war.Calling = warIn.Calling
	}
	if err := m.dataWardBusinessRel.Update(*war); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *war, WardBusinessRelConst, a.KeysToString("Id", war.Id), existingValues)
	return nil
}

func (m *ManagerWardBusinessRel) Delete(war *WardBusinessRel) error {

	if err := m.dataWardBusinessRel.Delete(war); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *war, WardBusinessRelConst, a.KeysToString("Id", war.Id))
	return nil
}
