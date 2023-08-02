package speaker

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	SpeakerGrpc struct {
		p.UnimplementedSpeakerServiceServer
		managerSpeaker ManagerSpeakerAdapter
	}
)

func NewSpeakerGrpc(mspe ManagerSpeakerAdapter) *SpeakerGrpc {
	return &SpeakerGrpc{managerSpeaker: mspe}
}

func (a *SpeakerGrpc) GetSpeaker(ctx context.Context, in *p.SpeakerIDIn) (*p.SpeakerResponse, error) {
	result := &p.Result{Success: false}
	response := &p.SpeakerResponse{Result: result}
	spe := &Speaker{Id: int(in.Id)}
	if err := a.managerSpeaker.Get(spe); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Speaker, err = translateOut(spe)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *SpeakerGrpc) SearchSpeaker(ctx context.Context, in *p.Speaker) (*p.SpeakerRepeatResponse, error) {
	speakerParam := SpeakerParam{}
	result := &p.Result{Success: false}
	response := &p.SpeakerRepeatResponse{Result: result}
	spes := &[]Speaker{}
	if _, err := a.managerSpeaker.Search(spes, speakerParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *spes {
		protoSpeaker, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Speaker = append(response.Speaker, protoSpeaker)
	}
	response.Result.Success = true
	return response, nil
}

func (a *SpeakerGrpc) PostSpeaker(ctx context.Context, in *p.Speaker) (*p.SpeakerResponse, error) {
	result := &p.Result{Success: false}
	response := &p.SpeakerResponse{Result: result}
	spe, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerSpeaker.Post(spe); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Speaker, errTranslate = translateOut(spe)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *SpeakerGrpc) PatchSpeaker(ctx context.Context, in *p.Speaker) (*p.Result, error) {
	response := &p.Result{Success: false}
	spe, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerSpeaker.Patch(*spe); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *SpeakerGrpc) DeleteSpeaker(ctx context.Context, in *p.SpeakerIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	spe := &Speaker{Id: int(in.Id)}
	if err := a.managerSpeaker.Delete(spe); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(spe *Speaker) (*p.Speaker, error) {
	protoSpeaker := p.Speaker{}
	protoSpeaker.Id = int64(spe.Id)
	protoSpeaker.Date = spe.Date.String
	protoSpeaker.Position = spe.Position.String
	protoSpeaker.Name = spe.Name.String
	return &protoSpeaker, nil
}

func translateIn(in *p.Speaker) (*Speaker, error) {
	spe := Speaker{}
	spe.Id = int(in.Id)
	spe.Date.Scan(in.Date)
	spe.Position.Scan(in.Position)
	spe.Name.Scan(in.Name)
	return &spe, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(spe *Speaker) (*p.Speaker, error) {
	protoSpeaker := p.Speaker{}
	outBytes, err := json.Marshal(spe)
	if err != nil {
		return &protoSpeaker, ae.GeneralError("Unable to encode from Speaker", err)
	}
	err = json.Unmarshal(outBytes, &protoSpeaker)
	if err != nil {
		return &protoSpeaker, ae.GeneralError("Unable to decode to proto.Speaker", err)
	}
	return &protoSpeaker, nil
}

func translateJsonIn(in *p.Speaker) (*Speaker, error) {
	spe := Speaker{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &spe, ae.GeneralError("Unable to encode from proto.Speaker", err)
	}
	err = json.Unmarshal(outBytes, &spe)
	if err != nil {
		return &spe, ae.GeneralError("Unable to decode to Speaker", err)
	}
	return &spe, nil
}
