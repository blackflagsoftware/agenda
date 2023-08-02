package ordinance

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	OrdinanceGrpc struct {
		p.UnimplementedOrdinanceServiceServer
		managerOrdinance ManagerOrdinanceAdapter
	}
)

func NewOrdinanceGrpc(mord ManagerOrdinanceAdapter) *OrdinanceGrpc {
	return &OrdinanceGrpc{managerOrdinance: mord}
}

func (a *OrdinanceGrpc) GetOrdinance(ctx context.Context, in *p.OrdinanceIDIn) (*p.OrdinanceResponse, error) {
	result := &p.Result{Success: false}
	response := &p.OrdinanceResponse{Result: result}
	ord := &Ordinance{Id: int(in.Id)}
	if err := a.managerOrdinance.Get(ord); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Ordinance, err = translateOut(ord)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *OrdinanceGrpc) SearchOrdinance(ctx context.Context, in *p.Ordinance) (*p.OrdinanceRepeatResponse, error) {
	ordinanceParam := OrdinanceParam{}
	result := &p.Result{Success: false}
	response := &p.OrdinanceRepeatResponse{Result: result}
	ords := &[]Ordinance{}
	if _, err := a.managerOrdinance.Search(ords, ordinanceParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *ords {
		protoOrdinance, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Ordinance = append(response.Ordinance, protoOrdinance)
	}
	response.Result.Success = true
	return response, nil
}

func (a *OrdinanceGrpc) PostOrdinance(ctx context.Context, in *p.Ordinance) (*p.OrdinanceResponse, error) {
	result := &p.Result{Success: false}
	response := &p.OrdinanceResponse{Result: result}
	ord, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerOrdinance.Post(ord); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Ordinance, errTranslate = translateOut(ord)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *OrdinanceGrpc) PatchOrdinance(ctx context.Context, in *p.Ordinance) (*p.Result, error) {
	response := &p.Result{Success: false}
	ord, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerOrdinance.Patch(*ord); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *OrdinanceGrpc) DeleteOrdinance(ctx context.Context, in *p.OrdinanceIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	ord := &Ordinance{Id: int(in.Id)}
	if err := a.managerOrdinance.Delete(ord); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(ord *Ordinance) (*p.Ordinance, error) {
	protoOrdinance := p.Ordinance{}
	protoOrdinance.Id = int64(ord.Id)
	protoOrdinance.Date = ord.Date.String
	protoOrdinance.Confirmations = ord.Confirmations.String
	protoOrdinance.Blessings = ord.Blessings.String
	return &protoOrdinance, nil
}

func translateIn(in *p.Ordinance) (*Ordinance, error) {
	ord := Ordinance{}
	ord.Id = int(in.Id)
	ord.Date.Scan(in.Date)
	ord.Confirmations.Scan(in.Confirmations)
	ord.Blessings.Scan(in.Blessings)
	return &ord, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(ord *Ordinance) (*p.Ordinance, error) {
	protoOrdinance := p.Ordinance{}
	outBytes, err := json.Marshal(ord)
	if err != nil {
		return &protoOrdinance, ae.GeneralError("Unable to encode from Ordinance", err)
	}
	err = json.Unmarshal(outBytes, &protoOrdinance)
	if err != nil {
		return &protoOrdinance, ae.GeneralError("Unable to decode to proto.Ordinance", err)
	}
	return &protoOrdinance, nil
}

func translateJsonIn(in *p.Ordinance) (*Ordinance, error) {
	ord := Ordinance{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &ord, ae.GeneralError("Unable to encode from proto.Ordinance", err)
	}
	err = json.Unmarshal(outBytes, &ord)
	if err != nil {
		return &ord, ae.GeneralError("Unable to decode to Ordinance", err)
	}
	return &ord, nil
}
