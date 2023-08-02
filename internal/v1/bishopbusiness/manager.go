package bishopbusiness

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=bishopbusiness
type (
	DataBishopBusinessAdapter interface {
		Read(*BishopBusiness) error
		ReadAll(*[]BishopBusiness, BishopBusinessParam) (int, error)
		Create(*BishopBusiness) error
		Update(BishopBusiness) error
		Delete(*BishopBusiness) error
	}

	ManagerBishopBusiness struct {
		dataBishopBusiness DataBishopBusinessAdapter
		auditWriter        a.AuditAdapter
	}
)

func NewManagerBishopBusiness(cbis DataBishopBusinessAdapter) *ManagerBishopBusiness {
	aw := a.AuditInit()
	return &ManagerBishopBusiness{dataBishopBusiness: cbis, auditWriter: aw}
}

func (m *ManagerBishopBusiness) Get(bis *BishopBusiness) error {

	return m.dataBishopBusiness.Read(bis)
}

func (m *ManagerBishopBusiness) Search(bis *[]BishopBusiness, param BishopBusinessParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "message": "message"})

	return m.dataBishopBusiness.ReadAll(bis, param)
}

func (m *ManagerBishopBusiness) Post(bis *BishopBusiness) error {

	if err := m.dataBishopBusiness.Create(bis); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *bis, BishopBusinessConst, a.KeysToString("Id", bis.Id))
	return nil
}

func (m *ManagerBishopBusiness) Patch(bisIn BishopBusiness) error {
	bis := &BishopBusiness{Id: bisIn.Id}
	errGet := m.dataBishopBusiness.Read(bis)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if bisIn.Date.Valid {
		existingValues["date"] = bis.Date.String
		bis.Date = bisIn.Date
	}
	// Message
	if bisIn.Message.Valid {
		existingValues["message"] = bis.Message.String
		bis.Message = bisIn.Message
	}
	if err := m.dataBishopBusiness.Update(*bis); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *bis, BishopBusinessConst, a.KeysToString("Id", bis.Id), existingValues)
	return nil
}

func (m *ManagerBishopBusiness) Delete(bis *BishopBusiness) error {

	if err := m.dataBishopBusiness.Delete(bis); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *bis, BishopBusinessConst, a.KeysToString("Id", bis.Id))
	return nil
}
