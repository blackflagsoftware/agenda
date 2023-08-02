package wardbusinesssus

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	WardBusinessSusGrpc struct {
		p.UnimplementedWardBusinessSusServiceServer
		managerWardBusinessSus ManagerWardBusinessSusAdapter
	}
)

func NewWardBusinessSusGrpc(mwa ManagerWardBusinessSusAdapter) *WardBusinessSusGrpc {
	return &WardBusinessSusGrpc{managerWardBusinessSus: mwa}
}

func (a *WardBusinessSusGrpc) GetWardBusinessSus(ctx context.Context, in *p.WardBusinessSusIDIn) (*p.WardBusinessSusResponse, error) {
	result := &p.Result{Success: false}
	response := &p.WardBusinessSusResponse{Result: result}
	wa := &WardBusinessSus{Id: int(in.Id)}
	if err := a.managerWardBusinessSus.Get(wa); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.WardBusinessSus, err = translateOut(wa)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *WardBusinessSusGrpc) SearchWardBusinessSus(ctx context.Context, in *p.WardBusinessSus) (*p.WardBusinessSusRepeatResponse, error) {
	wardBusinessSusParam := WardBusinessSusParam{}
	result := &p.Result{Success: false}
	response := &p.WardBusinessSusRepeatResponse{Result: result}
	was := &[]WardBusinessSus{}
	if _, err := a.managerWardBusinessSus.Search(was, wardBusinessSusParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *was {
		protoWardBusinessSus, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.WardBusinessSus = append(response.WardBusinessSus, protoWardBusinessSus)
	}
	response.Result.Success = true
	return response, nil
}

func (a *WardBusinessSusGrpc) PostWardBusinessSus(ctx context.Context, in *p.WardBusinessSus) (*p.WardBusinessSusResponse, error) {
	result := &p.Result{Success: false}
	response := &p.WardBusinessSusResponse{Result: result}
	wa, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerWardBusinessSus.Post(wa); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.WardBusinessSus, errTranslate = translateOut(wa)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *WardBusinessSusGrpc) PatchWardBusinessSus(ctx context.Context, in *p.WardBusinessSus) (*p.Result, error) {
	response := &p.Result{Success: false}
	wa, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerWardBusinessSus.Patch(*wa); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *WardBusinessSusGrpc) DeleteWardBusinessSus(ctx context.Context, in *p.WardBusinessSusIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	wa := &WardBusinessSus{Id: int(in.Id)}
	if err := a.managerWardBusinessSus.Delete(wa); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(wa *WardBusinessSus) (*p.WardBusinessSus, error) {
	protoWardBusinessSus := p.WardBusinessSus{}
	protoWardBusinessSus.Id = int64(wa.Id)
	protoWardBusinessSus.Date = wa.Date.String
	protoWardBusinessSus.Name = wa.Name.String
	protoWardBusinessSus.Calling = wa.Calling.String
	return &protoWardBusinessSus, nil
}

func translateIn(in *p.WardBusinessSus) (*WardBusinessSus, error) {
	wa := WardBusinessSus{}
	wa.Id = int(in.Id)
	wa.Date.Scan(in.Date)
	wa.Name.Scan(in.Name)
	wa.Calling.Scan(in.Calling)
	return &wa, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(wa *WardBusinessSus) (*p.WardBusinessSus, error) {
	protoWardBusinessSus := p.WardBusinessSus{}
	outBytes, err := json.Marshal(wa)
	if err != nil {
		return &protoWardBusinessSus, ae.GeneralError("Unable to encode from WardBusinessSus", err)
	}
	err = json.Unmarshal(outBytes, &protoWardBusinessSus)
	if err != nil {
		return &protoWardBusinessSus, ae.GeneralError("Unable to decode to proto.WardBusinessSus", err)
	}
	return &protoWardBusinessSus, nil
}

func translateJsonIn(in *p.WardBusinessSus) (*WardBusinessSus, error) {
	wa := WardBusinessSus{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &wa, ae.GeneralError("Unable to encode from proto.WardBusinessSus", err)
	}
	err = json.Unmarshal(outBytes, &wa)
	if err != nil {
		return &wa, ae.GeneralError("Unable to decode to WardBusinessSus", err)
	}
	return &wa, nil
}
