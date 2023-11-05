package speaker

import (
	a "github.com/blackflagsoftware/agenda/internal/audit"
)

//go:generate mockgen -source=manager.go -destination=mock.go -package=speaker
type (
	DataSpeakerAdapter interface {
		Read(*Speaker) error
		ReadAll(*[]Speaker, SpeakerParam) (int, error)
		Create(*Speaker) error
		Update(Speaker) error
		Delete(*Speaker) error
	}

	ManagerSpeaker struct {
		dataSpeaker DataSpeakerAdapter
		auditWriter a.AuditAdapter
	}
)

func NewManagerSpeaker(cspe DataSpeakerAdapter) *ManagerSpeaker {
	aw := a.AuditInit()
	return &ManagerSpeaker{dataSpeaker: cspe, auditWriter: aw}
}

func (m *ManagerSpeaker) Get(spe *Speaker) error {

	return m.dataSpeaker.Read(spe)
}

func (m *ManagerSpeaker) Search(spe *[]Speaker, param SpeakerParam) (int, error) {
	param.Param.CalculateParam("date", map[string]string{"date": "date", "position": "position", "name": "name"})

	return m.dataSpeaker.ReadAll(spe, param)
}

func (m *ManagerSpeaker) Post(spe *Speaker) error {

	if err := m.dataSpeaker.Create(spe); err != nil {
		return nil
	}
	go a.AuditCreate(m.auditWriter, *spe, SpeakerConst, a.KeysToString("Id", spe.Id))
	return nil
}

func (m *ManagerSpeaker) Patch(speIn Speaker) error {
	spe := &Speaker{Id: speIn.Id}
	errGet := m.dataSpeaker.Read(spe)
	if errGet != nil {
		return errGet
	}
	existingValues := make(map[string]interface{})
	// Date
	if speIn.Date.Valid {
		existingValues["date"] = spe.Date.String
		spe.Date = speIn.Date
	}
	// Position
	if speIn.Position.Valid {
		existingValues["position"] = spe.Position.String
		spe.Position = speIn.Position
	}
	// SpeakerType
	if speIn.SpeakerType.Valid {
		existingValues["speaker_type"] = spe.SpeakerType.String
		spe.SpeakerType = speIn.SpeakerType
	}
	// Name
	if speIn.Name.Valid {
		existingValues["name"] = spe.Name.String
		spe.Name = speIn.Name
	}
	if err := m.dataSpeaker.Update(*spe); err != nil {
		return err
	}
	go a.AuditPatch(m.auditWriter, *spe, SpeakerConst, a.KeysToString("Id", spe.Id), existingValues)
	return nil
}

func (m *ManagerSpeaker) Delete(spe *Speaker) error {

	if err := m.dataSpeaker.Delete(spe); err != nil {
		return err
	}
	go a.AuditDelete(m.auditWriter, *spe, SpeakerConst, a.KeysToString("Id", spe.Id))
	return nil
}
