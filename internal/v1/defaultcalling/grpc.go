package defaultcalling

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	DefaultCallingGrpc struct {
		p.UnimplementedDefaultCallingServiceServer
		managerDefaultCalling ManagerDefaultCallingAdapter
	}
)

func NewDefaultCallingGrpc(mdef ManagerDefaultCallingAdapter) *DefaultCallingGrpc {
	return &DefaultCallingGrpc{managerDefaultCalling: mdef}
}

func (a *DefaultCallingGrpc) GetDefaultCalling(ctx context.Context, in *p.DefaultCallingIDIn) (*p.DefaultCallingResponse, error) {
	result := &p.Result{Success: false}
	response := &p.DefaultCallingResponse{Result: result}
	def := &DefaultCalling{Id: int(in.Id)}
	if err := a.managerDefaultCalling.Get(def); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.DefaultCalling, err = translateOut(def)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *DefaultCallingGrpc) SearchDefaultCalling(ctx context.Context, in *p.DefaultCalling) (*p.DefaultCallingRepeatResponse, error) {
	defaultCallingParam := DefaultCallingParam{}
	result := &p.Result{Success: false}
	response := &p.DefaultCallingRepeatResponse{Result: result}
	defs := &[]DefaultCalling{}
	if _, err := a.managerDefaultCalling.Search(defs, defaultCallingParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *defs {
		protoDefaultCalling, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.DefaultCalling = append(response.DefaultCalling, protoDefaultCalling)
	}
	response.Result.Success = true
	return response, nil
}

func (a *DefaultCallingGrpc) PostDefaultCalling(ctx context.Context, in *p.DefaultCalling) (*p.DefaultCallingResponse, error) {
	result := &p.Result{Success: false}
	response := &p.DefaultCallingResponse{Result: result}
	def, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerDefaultCalling.Post(def); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.DefaultCalling, errTranslate = translateOut(def)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *DefaultCallingGrpc) PatchDefaultCalling(ctx context.Context, in *p.DefaultCalling) (*p.Result, error) {
	response := &p.Result{Success: false}
	def, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerDefaultCalling.Patch(*def); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *DefaultCallingGrpc) DeleteDefaultCalling(ctx context.Context, in *p.DefaultCallingIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	def := &DefaultCalling{Id: int(in.Id)}
	if err := a.managerDefaultCalling.Delete(def); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(def *DefaultCalling) (*p.DefaultCalling, error) {
	protoDefaultCalling := p.DefaultCalling{}
	protoDefaultCalling.Id = int64(def.Id)
	protoDefaultCalling.Organist = def.Organist.String
	protoDefaultCalling.Chorister = def.Chorister.String
	protoDefaultCalling.Newsletter = def.Newsletter.String
	protoDefaultCalling.Stake = def.Stake.String
	return &protoDefaultCalling, nil
}

func translateIn(in *p.DefaultCalling) (*DefaultCalling, error) {
	def := DefaultCalling{}
	def.Id = int(in.Id)
	def.Organist.Scan(in.Organist)
	def.Chorister.Scan(in.Chorister)
	def.Newsletter.Scan(in.Newsletter)
	def.Stake.Scan(in.Stake)
	return &def, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(def *DefaultCalling) (*p.DefaultCalling, error) {
	protoDefaultCalling := p.DefaultCalling{}
	outBytes, err := json.Marshal(def)
	if err != nil {
		return &protoDefaultCalling, ae.GeneralError("Unable to encode from DefaultCalling", err)
	}
	err = json.Unmarshal(outBytes, &protoDefaultCalling)
	if err != nil {
		return &protoDefaultCalling, ae.GeneralError("Unable to decode to proto.DefaultCalling", err)
	}
	return &protoDefaultCalling, nil
}

func translateJsonIn(in *p.DefaultCalling) (*DefaultCalling, error) {
	def := DefaultCalling{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &def, ae.GeneralError("Unable to encode from proto.DefaultCalling", err)
	}
	err = json.Unmarshal(outBytes, &def)
	if err != nil {
		return &def, ae.GeneralError("Unable to decode to DefaultCalling", err)
	}
	return &def, nil
}
