package role

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=role
type (
	DataRoleAdapter interface {
		Read(*Role) error
		ReadAll(*[]Role, RoleParam) (int, error)
		Create(*Role) error
		Update(Role) error
		Delete(*Role) error
	}

	ManagerRole struct {
		dataRole    DataRoleAdapter
		auditWriter a.AuditAdapter
	}
)

func NewManagerRole(crol DataRoleAdapter) *ManagerRole {
	aw := a.AuditInit()
	return &ManagerRole{dataRole: crol, auditWriter: aw}
}

func (m *ManagerRole) Get(rol *Role) error {

	return m.dataRole.Read(rol)
}

func (m *ManagerRole) Search(rol *[]Role, param RoleParam) (int, error) {
	param.Param.CalculateParam("name", map[string]string{"name": "name"})

	return m.dataRole.ReadAll(rol, param)
}

func (m *ManagerRole) Post(rol *Role) error {

	if err := m.dataRole.Create(rol); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *rol, RoleConst, a.KeysToString("Id", rol.Id))
	return nil
}

func (m *ManagerRole) Patch(rolIn Role) error {
	rol := &Role{Id: rolIn.Id}
	errGet := m.dataRole.Read(rol)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Name
	if rolIn.Name.Valid {
		existingValues["name"] = rol.Name.String
		rol.Name = rolIn.Name
	}
	if err := m.dataRole.Update(*rol); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *rol, RoleConst, a.KeysToString("Id", rol.Id), existingValues)
	return nil
}

func (m *ManagerRole) Delete(rol *Role) error {

	if err := m.dataRole.Delete(rol); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *rol, RoleConst, a.KeysToString("Id", rol.Id))
	return nil
}
