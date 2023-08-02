package announcement

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	AnnouncementGrpc struct {
		p.UnimplementedAnnouncementServiceServer
		managerAnnouncement ManagerAnnouncementAdapter
	}
)

func NewAnnouncementGrpc(mann ManagerAnnouncementAdapter) *AnnouncementGrpc {
	return &AnnouncementGrpc{managerAnnouncement: mann}
}

func (a *AnnouncementGrpc) GetAnnouncement(ctx context.Context, in *p.AnnouncementIDIn) (*p.AnnouncementResponse, error) {
	result := &p.Result{Success: false}
	response := &p.AnnouncementResponse{Result: result}
	ann := &Announcement{Id: int(in.Id)}
	if err := a.managerAnnouncement.Get(ann); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Announcement, err = translateOut(ann)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *AnnouncementGrpc) SearchAnnouncement(ctx context.Context, in *p.Announcement) (*p.AnnouncementRepeatResponse, error) {
	announcementParam := AnnouncementParam{}
	result := &p.Result{Success: false}
	response := &p.AnnouncementRepeatResponse{Result: result}
	anns := &[]Announcement{}
	if _, err := a.managerAnnouncement.Search(anns, announcementParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *anns {
		protoAnnouncement, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Announcement = append(response.Announcement, protoAnnouncement)
	}
	response.Result.Success = true
	return response, nil
}

func (a *AnnouncementGrpc) PostAnnouncement(ctx context.Context, in *p.Announcement) (*p.AnnouncementResponse, error) {
	result := &p.Result{Success: false}
	response := &p.AnnouncementResponse{Result: result}
	ann, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerAnnouncement.Post(ann); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Announcement, errTranslate = translateOut(ann)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *AnnouncementGrpc) PatchAnnouncement(ctx context.Context, in *p.Announcement) (*p.Result, error) {
	response := &p.Result{Success: false}
	ann, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerAnnouncement.Patch(*ann); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *AnnouncementGrpc) DeleteAnnouncement(ctx context.Context, in *p.AnnouncementIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	ann := &Announcement{Id: int(in.Id)}
	if err := a.managerAnnouncement.Delete(ann); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(ann *Announcement) (*p.Announcement, error) {
	protoAnnouncement := p.Announcement{}
	protoAnnouncement.Id = int64(ann.Id)
	protoAnnouncement.Date = ann.Date.String
	protoAnnouncement.Message = ann.Message.String
	protoAnnouncement.Pulpit = ann.Pulpit.Bool
	return &protoAnnouncement, nil
}

func translateIn(in *p.Announcement) (*Announcement, error) {
	ann := Announcement{}
	ann.Id = int(in.Id)
	ann.Date.Scan(in.Date)
	ann.Message.Scan(in.Message)
	ann.Pulpit.Scan(in.Pulpit)
	return &ann, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(ann *Announcement) (*p.Announcement, error) {
	protoAnnouncement := p.Announcement{}
	outBytes, err := json.Marshal(ann)
	if err != nil {
		return &protoAnnouncement, ae.GeneralError("Unable to encode from Announcement", err)
	}
	err = json.Unmarshal(outBytes, &protoAnnouncement)
	if err != nil {
		return &protoAnnouncement, ae.GeneralError("Unable to decode to proto.Announcement", err)
	}
	return &protoAnnouncement, nil
}

func translateJsonIn(in *p.Announcement) (*Announcement, error) {
	ann := Announcement{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &ann, ae.GeneralError("Unable to encode from proto.Announcement", err)
	}
	err = json.Unmarshal(outBytes, &ann)
	if err != nil {
		return &ann, ae.GeneralError("Unable to decode to Announcement", err)
	}
	return &ann, nil
}
