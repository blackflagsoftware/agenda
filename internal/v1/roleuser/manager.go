package roleuser

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=roleuser
type (
	DataRoleUserAdapter interface {
		Login(RoleUser) (RoleLogin, error)
		Read(*RoleUser) error
		ReadAll(*[]RoleUser, RoleUserParam) (int, error)
		Create(*RoleUser) error
		Update(RoleUser) error
		Delete(*RoleUser) error
	}

	ManagerRoleUser struct {
		dataRoleUser DataRoleUserAdapter
		auditWriter  a.AuditAdapter
	}
)

func NewManagerRoleUser(cro DataRoleUserAdapter) *ManagerRoleUser {
	aw := a.AuditInit()
	return &ManagerRoleUser{dataRoleUser: cro, auditWriter: aw}
}

func (m *ManagerRoleUser) Login(ro RoleUser) (RoleLogin, error) {
	if ro.Name.String == "" {
		return RoleLogin{}, ae.MissingParamError("Name")
	}
	if ro.Pwd.String == "" {
		return RoleLogin{}, ae.MissingParamError("Password")
	}
	return m.dataRoleUser.Login(ro)
}

func (m *ManagerRoleUser) Get(ro *RoleUser) error {

	return m.dataRoleUser.Read(ro)
}

func (m *ManagerRoleUser) Search(ro *[]RoleUser, param RoleUserParam) (int, error) {
	param.Param.CalculateParam("role_id", map[string]string{"role_id": "role_id", "name": "name", "pwd": "pwd"})

	return m.dataRoleUser.ReadAll(ro, param)
}

func (m *ManagerRoleUser) Post(ro *RoleUser) error {

	if err := m.dataRoleUser.Create(ro); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *ro, RoleUserConst, a.KeysToString("Id", ro.Id))
	return nil
}

func (m *ManagerRoleUser) Patch(roIn RoleUser) error {
	ro := &RoleUser{Id: roIn.Id}
	errGet := m.dataRoleUser.Read(ro)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// RoleId
	if roIn.RoleId.Valid {
		existingValues["role_id"] = ro.RoleId.Int64
		ro.RoleId = roIn.RoleId
	}
	// Name
	if roIn.Name.Valid {
		existingValues["name"] = ro.Name.String
		ro.Name = roIn.Name
	}
	// Pwd
	if roIn.Pwd.Valid {
		existingValues["pwd"] = ro.Pwd.String
		ro.Pwd = roIn.Pwd
	}
	if err := m.dataRoleUser.Update(*ro); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *ro, RoleUserConst, a.KeysToString("Id", ro.Id), existingValues)
	return nil
}

func (m *ManagerRoleUser) Delete(ro *RoleUser) error {

	if err := m.dataRoleUser.Delete(ro); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *ro, RoleUserConst, a.KeysToString("Id", ro.Id))
	return nil
}
