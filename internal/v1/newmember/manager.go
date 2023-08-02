package newmember

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=newmember
type (
	DataNewMemberAdapter interface {
		Read(*NewMember) error
		ReadAll(*[]NewMember, NewMemberParam) (int, error)
		Create(*NewMember) error
		Update(NewMember) error
		Delete(*NewMember) error
	}

	ManagerNewMember struct {
		dataNewMember DataNewMemberAdapter
		auditWriter   a.AuditAdapter
	}
)

func NewManagerNewMember(cnew DataNewMemberAdapter) *ManagerNewMember {
	aw := a.AuditInit()
	return &ManagerNewMember{dataNewMember: cnew, auditWriter: aw}
}

func (m *ManagerNewMember) Get(new *NewMember) error {

	return m.dataNewMember.Read(new)
}

func (m *ManagerNewMember) Search(new *[]NewMember, param NewMemberParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "family_name": "family_name", "names": "names"})

	return m.dataNewMember.ReadAll(new, param)
}

func (m *ManagerNewMember) Post(new *NewMember) error {

	if err := m.dataNewMember.Create(new); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *new, NewMemberConst, a.KeysToString("Id", new.Id))
	return nil
}

func (m *ManagerNewMember) Patch(newIn NewMember) error {
	new := &NewMember{Id: newIn.Id}
	errGet := m.dataNewMember.Read(new)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if newIn.Date.Valid {
		existingValues["date"] = new.Date.String
		new.Date = newIn.Date
	}
	// FamilyName
	if newIn.FamilyName.Valid {
		existingValues["family_name"] = new.FamilyName.String
		new.FamilyName = newIn.FamilyName
	}
	// Names
	if newIn.Names.Valid {
		existingValues["names"] = new.Names.String
		new.Names = newIn.Names
	}
	if err := m.dataNewMember.Update(*new); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *new, NewMemberConst, a.KeysToString("Id", new.Id), existingValues)
	return nil
}

func (m *ManagerNewMember) Delete(new *NewMember) error {

	if err := m.dataNewMember.Delete(new); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *new, NewMemberConst, a.KeysToString("Id", new.Id))
	return nil
}
