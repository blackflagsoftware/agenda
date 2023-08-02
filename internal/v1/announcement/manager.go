package announcement

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=announcement
type (
	DataAnnouncementAdapter interface {
		Read(*Announcement) error
		ReadAll(*[]Announcement, AnnouncementParam) (int, error)
		Create(*Announcement) error
		Update(Announcement) error
		Delete(*Announcement) error
	}

	ManagerAnnouncement struct {
		dataAnnouncement DataAnnouncementAdapter
		auditWriter      a.AuditAdapter
	}
)

func NewManagerAnnouncement(cann DataAnnouncementAdapter) *ManagerAnnouncement {
	aw := a.AuditInit()
	return &ManagerAnnouncement{dataAnnouncement: cann, auditWriter: aw}
}

func (m *ManagerAnnouncement) Get(ann *Announcement) error {

	return m.dataAnnouncement.Read(ann)
}

func (m *ManagerAnnouncement) Search(ann *[]Announcement, param AnnouncementParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "message": "message", "pulpit": "pulpit"})

	return m.dataAnnouncement.ReadAll(ann, param)
}

func (m *ManagerAnnouncement) Post(ann *Announcement) error {

	if err := m.dataAnnouncement.Create(ann); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *ann, AnnouncementConst, a.KeysToString("Id", ann.Id))
	return nil
}

func (m *ManagerAnnouncement) Patch(annIn Announcement) error {
	ann := &Announcement{Id: annIn.Id}
	errGet := m.dataAnnouncement.Read(ann)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if annIn.Date.Valid {
		existingValues["date"] = ann.Date.String
		ann.Date = annIn.Date
	}
	// Message
	if annIn.Message.Valid {
		existingValues["message"] = ann.Message.String
		ann.Message = annIn.Message
	}
	// Pulpit
	if annIn.Pulpit.Valid {
		existingValues["pulpit"] = ann.Pulpit.Bool
		ann.Pulpit = annIn.Pulpit
	}
	if err := m.dataAnnouncement.Update(*ann); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *ann, AnnouncementConst, a.KeysToString("Id", ann.Id), existingValues)
	return nil
}

func (m *ManagerAnnouncement) Delete(ann *Announcement) error {

	if err := m.dataAnnouncement.Delete(ann); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *ann, AnnouncementConst, a.KeysToString("Id", ann.Id))
	return nil
}
