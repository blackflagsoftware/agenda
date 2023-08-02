package visitor

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=visitor
type (
	DataVisitorAdapter interface {
		Read(*Visitor) error
		ReadAll(*[]Visitor, VisitorParam) (int, error)
		Create(*Visitor) error
		Update(Visitor) error
		Delete(*Visitor) error
	}

	ManagerVisitor struct {
		dataVisitor DataVisitorAdapter
		auditWriter a.AuditAdapter
	}
)

func NewManagerVisitor(cvis DataVisitorAdapter) *ManagerVisitor {
	aw := a.AuditInit()
	return &ManagerVisitor{dataVisitor: cvis, auditWriter: aw}
}

func (m *ManagerVisitor) Get(vis *Visitor) error {

	return m.dataVisitor.Read(vis)
}

func (m *ManagerVisitor) Search(vis *[]Visitor, param VisitorParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "name": "name"})

	return m.dataVisitor.ReadAll(vis, param)
}

func (m *ManagerVisitor) Post(vis *Visitor) error {

	if err := m.dataVisitor.Create(vis); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *vis, VisitorConst, a.KeysToString("Id", vis.Id))
	return nil
}

func (m *ManagerVisitor) Patch(visIn Visitor) error {
	vis := &Visitor{Id: visIn.Id}
	errGet := m.dataVisitor.Read(vis)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if visIn.Date.Valid {
		existingValues["date"] = vis.Date.String
		vis.Date = visIn.Date
	}
	// Name
	if visIn.Name.Valid {
		existingValues["name"] = vis.Name.String
		vis.Name = visIn.Name
	}
	if err := m.dataVisitor.Update(*vis); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *vis, VisitorConst, a.KeysToString("Id", vis.Id), existingValues)
	return nil
}

func (m *ManagerVisitor) Delete(vis *Visitor) error {

	if err := m.dataVisitor.Delete(vis); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *vis, VisitorConst, a.KeysToString("Id", vis.Id))
	return nil
}
