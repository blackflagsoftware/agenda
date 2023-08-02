package roleuser

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	RoleUserGrpc struct {
		p.UnimplementedRoleUserServiceServer
		managerRoleUser ManagerRoleUserAdapter
	}
)

func NewRoleUserGrpc(mro ManagerRoleUserAdapter) *RoleUserGrpc {
	return &RoleUserGrpc{managerRoleUser: mro}
}

func (a *RoleUserGrpc) GetRoleUser(ctx context.Context, in *p.RoleUserIDIn) (*p.RoleUserResponse, error) {
	result := &p.Result{Success: false}
	response := &p.RoleUserResponse{Result: result}
	ro := &RoleUser{Id: int(in.Id)}
	if err := a.managerRoleUser.Get(ro); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.RoleUser, err = translateOut(ro)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *RoleUserGrpc) SearchRoleUser(ctx context.Context, in *p.RoleUser) (*p.RoleUserRepeatResponse, error) {
	roleUserParam := RoleUserParam{}
	result := &p.Result{Success: false}
	response := &p.RoleUserRepeatResponse{Result: result}
	ros := &[]RoleUser{}
	if _, err := a.managerRoleUser.Search(ros, roleUserParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *ros {
		protoRoleUser, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.RoleUser = append(response.RoleUser, protoRoleUser)
	}
	response.Result.Success = true
	return response, nil
}

func (a *RoleUserGrpc) PostRoleUser(ctx context.Context, in *p.RoleUser) (*p.RoleUserResponse, error) {
	result := &p.Result{Success: false}
	response := &p.RoleUserResponse{Result: result}
	ro, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerRoleUser.Post(ro); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.RoleUser, errTranslate = translateOut(ro)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *RoleUserGrpc) PatchRoleUser(ctx context.Context, in *p.RoleUser) (*p.Result, error) {
	response := &p.Result{Success: false}
	ro, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerRoleUser.Patch(*ro); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *RoleUserGrpc) DeleteRoleUser(ctx context.Context, in *p.RoleUserIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	ro := &RoleUser{Id: int(in.Id)}
	if err := a.managerRoleUser.Delete(ro); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(ro *RoleUser) (*p.RoleUser, error) {
	protoRoleUser := p.RoleUser{}
	protoRoleUser.Id = int64(ro.Id)
	protoRoleUser.RoleId = ro.RoleId.Int64
	protoRoleUser.Name = ro.Name.String
	protoRoleUser.Pwd = ro.Pwd.String
	return &protoRoleUser, nil
}

func translateIn(in *p.RoleUser) (*RoleUser, error) {
	ro := RoleUser{}
	ro.Id = int(in.Id)
	ro.RoleId.Scan(in.RoleId)
	ro.Name.Scan(in.Name)
	ro.Pwd.Scan(in.Pwd)
	return &ro, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(ro *RoleUser) (*p.RoleUser, error) {
	protoRoleUser := p.RoleUser{}
	outBytes, err := json.Marshal(ro)
	if err != nil {
		return &protoRoleUser, ae.GeneralError("Unable to encode from RoleUser", err)
	}
	err = json.Unmarshal(outBytes, &protoRoleUser)
	if err != nil {
		return &protoRoleUser, ae.GeneralError("Unable to decode to proto.RoleUser", err)
	}
	return &protoRoleUser, nil
}

func translateJsonIn(in *p.RoleUser) (*RoleUser, error) {
	ro := RoleUser{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &ro, ae.GeneralError("Unable to encode from proto.RoleUser", err)
	}
	err = json.Unmarshal(outBytes, &ro)
	if err != nil {
		return &ro, ae.GeneralError("Unable to decode to RoleUser", err)
	}
	return &ro, nil
}
