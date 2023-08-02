package newmember

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	NewMemberGrpc struct {
		p.UnimplementedNewMemberServiceServer
		managerNewMember ManagerNewMemberAdapter
	}
)

func NewNewMemberGrpc(mnew ManagerNewMemberAdapter) *NewMemberGrpc {
	return &NewMemberGrpc{managerNewMember: mnew}
}

func (a *NewMemberGrpc) GetNewMember(ctx context.Context, in *p.NewMemberIDIn) (*p.NewMemberResponse, error) {
	result := &p.Result{Success: false}
	response := &p.NewMemberResponse{Result: result}
	new := &NewMember{Id: int(in.Id)}
	if err := a.managerNewMember.Get(new); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.NewMember, err = translateOut(new)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *NewMemberGrpc) SearchNewMember(ctx context.Context, in *p.NewMember) (*p.NewMemberRepeatResponse, error) {
	newMemberParam := NewMemberParam{}
	result := &p.Result{Success: false}
	response := &p.NewMemberRepeatResponse{Result: result}
	news := &[]NewMember{}
	if _, err := a.managerNewMember.Search(news, newMemberParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *news {
		protoNewMember, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.NewMember = append(response.NewMember, protoNewMember)
	}
	response.Result.Success = true
	return response, nil
}

func (a *NewMemberGrpc) PostNewMember(ctx context.Context, in *p.NewMember) (*p.NewMemberResponse, error) {
	result := &p.Result{Success: false}
	response := &p.NewMemberResponse{Result: result}
	new, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerNewMember.Post(new); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.NewMember, errTranslate = translateOut(new)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *NewMemberGrpc) PatchNewMember(ctx context.Context, in *p.NewMember) (*p.Result, error) {
	response := &p.Result{Success: false}
	new, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerNewMember.Patch(*new); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *NewMemberGrpc) DeleteNewMember(ctx context.Context, in *p.NewMemberIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	new := &NewMember{Id: int(in.Id)}
	if err := a.managerNewMember.Delete(new); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(new *NewMember) (*p.NewMember, error) {
	protoNewMember := p.NewMember{}
	protoNewMember.Id = int64(new.Id)
	protoNewMember.Date = new.Date.String
	protoNewMember.FamilyName = new.FamilyName.String
	protoNewMember.Names = new.Names.String
	return &protoNewMember, nil
}

func translateIn(in *p.NewMember) (*NewMember, error) {
	new := NewMember{}
	new.Id = int(in.Id)
	new.Date.Scan(in.Date)
	new.FamilyName.Scan(in.FamilyName)
	new.Names.Scan(in.Names)
	return &new, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(new *NewMember) (*p.NewMember, error) {
	protoNewMember := p.NewMember{}
	outBytes, err := json.Marshal(new)
	if err != nil {
		return &protoNewMember, ae.GeneralError("Unable to encode from NewMember", err)
	}
	err = json.Unmarshal(outBytes, &protoNewMember)
	if err != nil {
		return &protoNewMember, ae.GeneralError("Unable to decode to proto.NewMember", err)
	}
	return &protoNewMember, nil
}

func translateJsonIn(in *p.NewMember) (*NewMember, error) {
	new := NewMember{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &new, ae.GeneralError("Unable to encode from proto.NewMember", err)
	}
	err = json.Unmarshal(outBytes, &new)
	if err != nil {
		return &new, ae.GeneralError("Unable to decode to NewMember", err)
	}
	return &new, nil
}
