package wardbusinessrel

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	WardBusinessRelGrpc struct {
		p.UnimplementedWardBusinessRelServiceServer
		managerWardBusinessRel ManagerWardBusinessRelAdapter
	}
)

func NewWardBusinessRelGrpc(mwar ManagerWardBusinessRelAdapter) *WardBusinessRelGrpc {
	return &WardBusinessRelGrpc{managerWardBusinessRel: mwar}
}

func (a *WardBusinessRelGrpc) GetWardBusinessRel(ctx context.Context, in *p.WardBusinessRelIDIn) (*p.WardBusinessRelResponse, error) {
	result := &p.Result{Success: false}
	response := &p.WardBusinessRelResponse{Result: result}
	war := &WardBusinessRel{Id: int(in.Id)}
	if err := a.managerWardBusinessRel.Get(war); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.WardBusinessRel, err = translateOut(war)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *WardBusinessRelGrpc) SearchWardBusinessRel(ctx context.Context, in *p.WardBusinessRel) (*p.WardBusinessRelRepeatResponse, error) {
	wardBusinessRelParam := WardBusinessRelParam{}
	result := &p.Result{Success: false}
	response := &p.WardBusinessRelRepeatResponse{Result: result}
	wars := &[]WardBusinessRel{}
	if _, err := a.managerWardBusinessRel.Search(wars, wardBusinessRelParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *wars {
		protoWardBusinessRel, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.WardBusinessRel = append(response.WardBusinessRel, protoWardBusinessRel)
	}
	response.Result.Success = true
	return response, nil
}

func (a *WardBusinessRelGrpc) PostWardBusinessRel(ctx context.Context, in *p.WardBusinessRel) (*p.WardBusinessRelResponse, error) {
	result := &p.Result{Success: false}
	response := &p.WardBusinessRelResponse{Result: result}
	war, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerWardBusinessRel.Post(war); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.WardBusinessRel, errTranslate = translateOut(war)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *WardBusinessRelGrpc) PatchWardBusinessRel(ctx context.Context, in *p.WardBusinessRel) (*p.Result, error) {
	response := &p.Result{Success: false}
	war, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerWardBusinessRel.Patch(*war); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *WardBusinessRelGrpc) DeleteWardBusinessRel(ctx context.Context, in *p.WardBusinessRelIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	war := &WardBusinessRel{Id: int(in.Id)}
	if err := a.managerWardBusinessRel.Delete(war); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(war *WardBusinessRel) (*p.WardBusinessRel, error) {
	protoWardBusinessRel := p.WardBusinessRel{}
	protoWardBusinessRel.Id = int64(war.Id)
	protoWardBusinessRel.Date = war.Date.String
	protoWardBusinessRel.Name = war.Name.String
	protoWardBusinessRel.Calling = war.Calling.String
	return &protoWardBusinessRel, nil
}

func translateIn(in *p.WardBusinessRel) (*WardBusinessRel, error) {
	war := WardBusinessRel{}
	war.Id = int(in.Id)
	war.Date.Scan(in.Date)
	war.Name.Scan(in.Name)
	war.Calling.Scan(in.Calling)
	return &war, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(war *WardBusinessRel) (*p.WardBusinessRel, error) {
	protoWardBusinessRel := p.WardBusinessRel{}
	outBytes, err := json.Marshal(war)
	if err != nil {
		return &protoWardBusinessRel, ae.GeneralError("Unable to encode from WardBusinessRel", err)
	}
	err = json.Unmarshal(outBytes, &protoWardBusinessRel)
	if err != nil {
		return &protoWardBusinessRel, ae.GeneralError("Unable to decode to proto.WardBusinessRel", err)
	}
	return &protoWardBusinessRel, nil
}

func translateJsonIn(in *p.WardBusinessRel) (*WardBusinessRel, error) {
	war := WardBusinessRel{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &war, ae.GeneralError("Unable to encode from proto.WardBusinessRel", err)
	}
	err = json.Unmarshal(outBytes, &war)
	if err != nil {
		return &war, ae.GeneralError("Unable to decode to WardBusinessRel", err)
	}
	return &war, nil
}
