package roles

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	RolesGrpc struct {
		p.UnimplementedRolesServiceServer
		managerRoles ManagerRolesAdapter
	}
)

func NewRolesGrpc(mrol ManagerRolesAdapter) *RolesGrpc {
	return &RolesGrpc{managerRoles: mrol}
}

func (a *RolesGrpc) GetRoles(ctx context.Context, in *p.RolesIDIn) (*p.RolesResponse, error) {
	result := &p.Result{Success: false}
	response := &p.RolesResponse{Result: result}
	rol := &Roles{Id: int(in.Id)}
	if err := a.managerRoles.Get(rol); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Roles, err = translateOut(rol)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *RolesGrpc) SearchRoles(ctx context.Context, in *p.Roles) (*p.RolesRepeatResponse, error) {
	rolesParam := RolesParam{}
	result := &p.Result{Success: false}
	response := &p.RolesRepeatResponse{Result: result}
	rols := &[]Roles{}
	if _, err := a.managerRoles.Search(rols, rolesParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *rols {
		protoRoles, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Roles = append(response.Roles, protoRoles)
	}
	response.Result.Success = true
	return response, nil
}

func (a *RolesGrpc) PostRoles(ctx context.Context, in *p.Roles) (*p.RolesResponse, error) {
	result := &p.Result{Success: false}
	response := &p.RolesResponse{Result: result}
	rol, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerRoles.Post(rol); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Roles, errTranslate = translateOut(rol)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *RolesGrpc) PatchRoles(ctx context.Context, in *p.Roles) (*p.Result, error) {
	response := &p.Result{Success: false}
	rol, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerRoles.Patch(*rol); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *RolesGrpc) DeleteRoles(ctx context.Context, in *p.RolesIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	rol := &Roles{Id: int(in.Id)}
	if err := a.managerRoles.Delete(rol); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(rol *Roles) (*p.Roles, error) {
	protoRoles := p.Roles{}
	protoRoles.Id = int64(rol.Id)
	protoRoles.Name = rol.Name.String
	return &protoRoles, nil
}

func translateIn(in *p.Roles) (*Roles, error) {
	rol := Roles{}
	rol.Id = int(in.Id)
	rol.Name.Scan(in.Name)
	return &rol, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(rol *Roles) (*p.Roles, error) {
	protoRoles := p.Roles{}
	outBytes, err := json.Marshal(rol)
	if err != nil {
		return &protoRoles, ae.GeneralError("Unable to encode from Roles", err)
	}
	err = json.Unmarshal(outBytes, &protoRoles)
	if err != nil {
		return &protoRoles, ae.GeneralError("Unable to decode to proto.Roles", err)
	}
	return &protoRoles, nil
}

func translateJsonIn(in *p.Roles) (*Roles, error) {
	rol := Roles{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &rol, ae.GeneralError("Unable to encode from proto.Roles", err)
	}
	err = json.Unmarshal(outBytes, &rol)
	if err != nil {
		return &rol, ae.GeneralError("Unable to decode to Roles", err)
	}
	return &rol, nil
}
