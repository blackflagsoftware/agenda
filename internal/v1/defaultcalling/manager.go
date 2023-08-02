package defaultcalling

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=defaultcalling
type (
	DataDefaultCallingAdapter interface {
		Read(*DefaultCalling) error
		ReadAll(*[]DefaultCalling, DefaultCallingParam) (int, error)
		Create(*DefaultCalling) error
		Update(DefaultCalling) error
		Delete(*DefaultCalling) error
	}

	ManagerDefaultCalling struct {
		dataDefaultCalling DataDefaultCallingAdapter
		auditWriter        a.AuditAdapter
	}
)

func NewManagerDefaultCalling(cdef DataDefaultCallingAdapter) *ManagerDefaultCalling {
	aw := a.AuditInit()
	return &ManagerDefaultCalling{dataDefaultCalling: cdef, auditWriter: aw}
}

func (m *ManagerDefaultCalling) Get(def *DefaultCalling) error {

	return m.dataDefaultCalling.Read(def)
}

func (m *ManagerDefaultCalling) Search(def *[]DefaultCalling, param DefaultCallingParam) (int, error) {
	param.Param.CalculateParam("organist", map[string]string{"organist": "organist", "chorister": "chorister", "newsletter": "newsletter", "stake": "stake"})

	return m.dataDefaultCalling.ReadAll(def, param)
}

func (m *ManagerDefaultCalling) Post(def *DefaultCalling) error {

	if err := m.dataDefaultCalling.Create(def); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *def, DefaultCallingConst, a.KeysToString("Id", def.Id))
	return nil
}

func (m *ManagerDefaultCalling) Patch(defIn DefaultCalling) error {
	def := &DefaultCalling{Id: defIn.Id}
	errGet := m.dataDefaultCalling.Read(def)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Bishop
	if defIn.Bishop.Valid {
		existingValues["bishop"] = def.Bishop.String
		def.Bishop = defIn.Bishop
	}
	// B1st
	if defIn.B1st.Valid {
		existingValues["b_1st"] = def.B1st.String
		def.B1st = defIn.B1st
	}
	// B2nd
	if defIn.B2nd.Valid {
		existingValues["b_2nd"] = def.B2nd.String
		def.B2nd = defIn.B2nd
	}
	// SPres
	if defIn.SPres.Valid {
		existingValues["s_pres"] = def.SPres.String
		def.SPres = defIn.SPres
	}
	// S1st
	if defIn.S1st.Valid {
		existingValues["s_1st"] = def.S1st.String
		def.S1st = defIn.S1st
	}
	// S2nd
	if defIn.S2nd.Valid {
		existingValues["s_2nd"] = def.S2nd.String
		def.S2nd = defIn.S2nd
	}
	// Organist
	if defIn.Organist.Valid {
		existingValues["organist"] = def.Organist.String
		def.Organist = defIn.Organist
	}
	// Chorister
	if defIn.Chorister.Valid {
		existingValues["chorister"] = def.Chorister.String
		def.Chorister = defIn.Chorister
	}
	// Newsletter
	if defIn.Newsletter.Valid {
		existingValues["newsletter"] = def.Newsletter.String
		def.Newsletter = defIn.Newsletter
	}
	// Stake
	if defIn.Stake.Valid {
		existingValues["stake"] = def.Stake.String
		def.Stake = defIn.Stake
	}
	if err := m.dataDefaultCalling.Update(*def); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *def, DefaultCallingConst, a.KeysToString("Id", def.Id), existingValues)
	return nil
}

func (m *ManagerDefaultCalling) Delete(def *DefaultCalling) error {

	if err := m.dataDefaultCalling.Delete(def); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *def, DefaultCallingConst, a.KeysToString("Id", def.Id))
	return nil
}
