package roles

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=roles
type (
	DataRolesAdapter interface {
		Read(*Roles) error
		ReadAll(*[]Roles, RolesParam) (int, error)
		Create(*Roles) error
		Update(Roles) error
		Delete(*Roles) error
	}

	ManagerRoles struct {
		dataRoles   DataRolesAdapter
		auditWriter a.AuditAdapter
	}
)

func NewManagerRoles(crol DataRolesAdapter) *ManagerRoles {
	aw := a.AuditInit()
	return &ManagerRoles{dataRoles: crol, auditWriter: aw}
}

func (m *ManagerRoles) Get(rol *Roles) error {

	return m.dataRoles.Read(rol)
}

func (m *ManagerRoles) Search(rol *[]Roles, param RolesParam) (int, error) {
	param.Param.CalculateParam("name", map[string]string{"name": "name"})

	return m.dataRoles.ReadAll(rol, param)
}

func (m *ManagerRoles) Post(rol *Roles) error {

	if err := m.dataRoles.Create(rol); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *rol, RolesConst, a.KeysToString("Id", rol.Id))
	return nil
}

func (m *ManagerRoles) Patch(rolIn Roles) error {
	rol := &Roles{Id: rolIn.Id}
	errGet := m.dataRoles.Read(rol)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Name
	if rolIn.Name.Valid {
		existingValues["name"] = rol.Name.String
		rol.Name = rolIn.Name
	}
	if err := m.dataRoles.Update(*rol); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *rol, RolesConst, a.KeysToString("Id", rol.Id), existingValues)
	return nil
}

func (m *ManagerRoles) Delete(rol *Roles) error {

	if err := m.dataRoles.Delete(rol); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *rol, RolesConst, a.KeysToString("Id", rol.Id))
	return nil
}
