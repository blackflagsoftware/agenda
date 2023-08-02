package member

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	MemberGrpc struct {
		p.UnimplementedMembersServiceServer
		managerMember ManagerMemberAdapter
	}
)

func NewMemberGrpc(mmem ManagerMemberAdapter) *MemberGrpc {
	return &MemberGrpc{managerMember: mmem}
}

func (a *MemberGrpc) GetMember(ctx context.Context, in *p.MemberIDIn) (*p.MemberResponse, error) {
	result := &p.Result{Success: false}
	response := &p.MemberResponse{Result: result}
	mem := &Member{Id: int(in.Id)}
	if err := a.managerMember.Get(mem); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Member, err = translateOut(mem)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *MemberGrpc) SearchMember(ctx context.Context, in *p.Member) (*p.MemberRepeatResponse, error) {
	memberParam := MemberParam{}
	result := &p.Result{Success: false}
	response := &p.MemberRepeatResponse{Result: result}
	mems := &[]Member{}
	if _, err := a.managerMember.Search(mems, memberParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *mems {
		protoMember, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Member = append(response.Member, protoMember)
	}
	response.Result.Success = true
	return response, nil
}

func (a *MemberGrpc) PostMember(ctx context.Context, in *p.Member) (*p.MemberResponse, error) {
	result := &p.Result{Success: false}
	response := &p.MemberResponse{Result: result}
	mem, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerMember.Post(mem); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Member, errTranslate = translateOut(mem)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *MemberGrpc) PatchMember(ctx context.Context, in *p.Member) (*p.Result, error) {
	response := &p.Result{Success: false}
	mem, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerMember.Patch(*mem); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *MemberGrpc) DeleteMember(ctx context.Context, in *p.MemberIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	mem := &Member{Id: int(in.Id)}
	if err := a.managerMember.Delete(mem); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(mem *Member) (*p.Member, error) {
	protoMember := p.Member{}
	protoMember.Id = int64(mem.Id)
	protoMember.FirstName = mem.FirstName.String
	protoMember.LastName = mem.LastName.String
	protoMember.Gender = mem.Gender.String
	protoMember.LastPrayed = mem.LastPrayed.String
	protoMember.LastTalked = mem.LastTalked.String
	protoMember.Active = mem.Active.Bool
	protoMember.NoPrayer = mem.NoPrayer.Bool
	protoMember.NoTalk = mem.NoTalk.Bool
	return &protoMember, nil
}

func translateIn(in *p.Member) (*Member, error) {
	mem := Member{}
	mem.Id = int(in.Id)
	mem.FirstName.Scan(in.FirstName)
	mem.LastName.Scan(in.LastName)
	mem.Gender.Scan(in.Gender)
	mem.LastPrayed.Scan(in.LastPrayed)
	mem.LastTalked.Scan(in.LastTalked)
	mem.Active.Scan(in.Active)
	mem.NoPrayer.Scan(in.NoPrayer)
	mem.NoTalk.Scan(in.NoTalk)
	return &mem, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(mem *Member) (*p.Member, error) {
	protoMember := p.Member{}
	outBytes, err := json.Marshal(mem)
	if err != nil {
		return &protoMember, ae.GeneralError("Unable to encode from Member", err)
	}
	err = json.Unmarshal(outBytes, &protoMember)
	if err != nil {
		return &protoMember, ae.GeneralError("Unable to decode to proto.Member", err)
	}
	return &protoMember, nil
}

func translateJsonIn(in *p.Member) (*Member, error) {
	mem := Member{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &mem, ae.GeneralError("Unable to encode from proto.Member", err)
	}
	err = json.Unmarshal(outBytes, &mem)
	if err != nil {
		return &mem, ae.GeneralError("Unable to decode to Member", err)
	}
	return &mem, nil
}
