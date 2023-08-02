package bishopbusiness

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	BishopBusinessGrpc struct {
		p.UnimplementedBishopBusinessServiceServer
		managerBishopBusiness ManagerBishopBusinessAdapter
	}
)

func NewBishopBusinessGrpc(mbis ManagerBishopBusinessAdapter) *BishopBusinessGrpc {
	return &BishopBusinessGrpc{managerBishopBusiness: mbis}
}

func (a *BishopBusinessGrpc) GetBishopBusiness(ctx context.Context, in *p.BishopBusinessIDIn) (*p.BishopBusinessResponse, error) {
	result := &p.Result{Success: false}
	response := &p.BishopBusinessResponse{Result: result}
	bis := &BishopBusiness{Id: int(in.Id)}
	if err := a.managerBishopBusiness.Get(bis); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.BishopBusiness, err = translateOut(bis)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *BishopBusinessGrpc) SearchBishopBusiness(ctx context.Context, in *p.BishopBusiness) (*p.BishopBusinessRepeatResponse, error) {
	bishopBusinessParam := BishopBusinessParam{}
	result := &p.Result{Success: false}
	response := &p.BishopBusinessRepeatResponse{Result: result}
	biss := &[]BishopBusiness{}
	if _, err := a.managerBishopBusiness.Search(biss, bishopBusinessParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *biss {
		protoBishopBusiness, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.BishopBusiness = append(response.BishopBusiness, protoBishopBusiness)
	}
	response.Result.Success = true
	return response, nil
}

func (a *BishopBusinessGrpc) PostBishopBusiness(ctx context.Context, in *p.BishopBusiness) (*p.BishopBusinessResponse, error) {
	result := &p.Result{Success: false}
	response := &p.BishopBusinessResponse{Result: result}
	bis, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerBishopBusiness.Post(bis); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.BishopBusiness, errTranslate = translateOut(bis)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *BishopBusinessGrpc) PatchBishopBusiness(ctx context.Context, in *p.BishopBusiness) (*p.Result, error) {
	response := &p.Result{Success: false}
	bis, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerBishopBusiness.Patch(*bis); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *BishopBusinessGrpc) DeleteBishopBusiness(ctx context.Context, in *p.BishopBusinessIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	bis := &BishopBusiness{Id: int(in.Id)}
	if err := a.managerBishopBusiness.Delete(bis); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(bis *BishopBusiness) (*p.BishopBusiness, error) {
	protoBishopBusiness := p.BishopBusiness{}
	protoBishopBusiness.Id = int64(bis.Id)
	protoBishopBusiness.Date = bis.Date.String
	protoBishopBusiness.Message = bis.Message.String
	return &protoBishopBusiness, nil
}

func translateIn(in *p.BishopBusiness) (*BishopBusiness, error) {
	bis := BishopBusiness{}
	bis.Id = int(in.Id)
	bis.Date.Scan(in.Date)
	bis.Message.Scan(in.Message)
	return &bis, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(bis *BishopBusiness) (*p.BishopBusiness, error) {
	protoBishopBusiness := p.BishopBusiness{}
	outBytes, err := json.Marshal(bis)
	if err != nil {
		return &protoBishopBusiness, ae.GeneralError("Unable to encode from BishopBusiness", err)
	}
	err = json.Unmarshal(outBytes, &protoBishopBusiness)
	if err != nil {
		return &protoBishopBusiness, ae.GeneralError("Unable to decode to proto.BishopBusiness", err)
	}
	return &protoBishopBusiness, nil
}

func translateJsonIn(in *p.BishopBusiness) (*BishopBusiness, error) {
	bis := BishopBusiness{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &bis, ae.GeneralError("Unable to encode from proto.BishopBusiness", err)
	}
	err = json.Unmarshal(outBytes, &bis)
	if err != nil {
		return &bis, ae.GeneralError("Unable to decode to BishopBusiness", err)
	}
	return &bis, nil
}
