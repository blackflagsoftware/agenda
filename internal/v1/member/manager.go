package member

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=member
type (
	DataMemberAdapter interface {
		Read(*Member) error
		ReadAll(*[]Member, MemberParam) (int, error)
		Create(*Member) error
		Update(Member) error
		Delete(*Member) error
		Splice() error
	}

	ManagerMember struct {
		dataMember  DataMemberAdapter
		auditWriter a.AuditAdapter
	}
)

func NewManagerMember(cmem DataMemberAdapter) *ManagerMember {
	aw := a.AuditInit()
	return &ManagerMember{dataMember: cmem, auditWriter: aw}
}

func (m *ManagerMember) Get(mem *Member) error {

	return m.dataMember.Read(mem)
}

func (m *ManagerMember) Search(mem *[]Member, param MemberParam) (int, error) {
	param.Param.CalculateParam("first_name", map[string]string{"first_name": "first_name", "last_name": "last_name", "gender": "gender", "last_prayed": "last_prayed", "last_talked": "last_talked", "active": "active", "no_prayer": "no_prayer", "no_talk": "no_talk"})

	return m.dataMember.ReadAll(mem, param)
}

func (m *ManagerMember) Post(mem *Member) error {

	if err := m.dataMember.Create(mem); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *mem, MemberConst, a.KeysToString("Id", mem.Id))
	return nil
}

func (m *ManagerMember) Patch(memIn Member) error {
	mem := &Member{Id: memIn.Id}
	errGet := m.dataMember.Read(mem)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// FirstName
	if memIn.FirstName.Valid {
		existingValues["first_name"] = mem.FirstName.String
		mem.FirstName = memIn.FirstName
	}
	// LastName
	if memIn.LastName.Valid {
		existingValues["last_name"] = mem.LastName.String
		mem.LastName = memIn.LastName
	}
	// Gender
	if memIn.Gender.Valid {
		existingValues["gender"] = mem.Gender.String
		mem.Gender = memIn.Gender
	}
	// LastPrayed
	if memIn.LastPrayed.Valid {
		existingValues["last_prayed"] = mem.LastPrayed.String
		mem.LastPrayed = memIn.LastPrayed
	}
	// LastTalked
	if memIn.LastTalked.Valid {
		existingValues["last_talked"] = mem.LastTalked.String
		mem.LastTalked = memIn.LastTalked
	}
	// Active
	if memIn.Active.Valid {
		existingValues["active"] = mem.Active.Bool
		mem.Active = memIn.Active
	}
	// NoPrayer
	if memIn.NoPrayer.Valid {
		existingValues["no_prayer"] = mem.NoPrayer.Bool
		mem.NoPrayer = memIn.NoPrayer
	}
	// NoTalk
	if memIn.NoTalk.Valid {
		existingValues["no_talk"] = mem.NoTalk.Bool
		mem.NoTalk = memIn.NoTalk
	}
	if err := m.dataMember.Update(*mem); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *mem, MemberConst, a.KeysToString("Id", mem.Id), existingValues)
	return nil
}

func (m *ManagerMember) Delete(mem *Member) error {

	if err := m.dataMember.Delete(mem); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *mem, MemberConst, a.KeysToString("Id", mem.Id))
	return nil
}

func (m *ManagerMember) Splice() error {
	return m.dataMember.Splice()
}
