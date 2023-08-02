package agenda

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	AgendaGrpc struct {
		p.UnimplementedAgendaServiceServer
		managerAgenda ManagerAgendaAdapter
	}
)

func NewAgendaGrpc(mage ManagerAgendaAdapter) *AgendaGrpc {
	return &AgendaGrpc{managerAgenda: mage}
}

func (a *AgendaGrpc) GetAgenda(ctx context.Context, in *p.AgendaIDIn) (*p.AgendaResponse, error) {
	result := &p.Result{Success: false}
	response := &p.AgendaResponse{Result: result}
	age := &Agenda{Date: in.Date}
	if err := a.managerAgenda.Get(age); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Agenda, err = translateOut(age)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *AgendaGrpc) SearchAgenda(ctx context.Context, in *p.Agenda) (*p.AgendaRepeatResponse, error) {
	agendaParam := AgendaParam{}
	result := &p.Result{Success: false}
	response := &p.AgendaRepeatResponse{Result: result}
	ages := &[]Agenda{}
	if _, err := a.managerAgenda.Search(ages, agendaParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *ages {
		protoAgenda, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Agenda = append(response.Agenda, protoAgenda)
	}
	response.Result.Success = true
	return response, nil
}

func (a *AgendaGrpc) PostAgenda(ctx context.Context, in *p.Agenda) (*p.AgendaResponse, error) {
	result := &p.Result{Success: false}
	response := &p.AgendaResponse{Result: result}
	age, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerAgenda.Post(age); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Agenda, errTranslate = translateOut(age)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *AgendaGrpc) PatchAgenda(ctx context.Context, in *p.Agenda) (*p.Result, error) {
	response := &p.Result{Success: false}
	age, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerAgenda.Patch(*age); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *AgendaGrpc) DeleteAgenda(ctx context.Context, in *p.AgendaIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	age := &Agenda{Date: in.Date}
	if err := a.managerAgenda.Delete(age); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(age *Agenda) (*p.Agenda, error) {
	protoAgenda := p.Agenda{}
	protoAgenda.Date = age.Date
	protoAgenda.Presiding = age.Presiding.String
	protoAgenda.Conducting = age.Conducting.String
	protoAgenda.Organist = age.Organist.String
	protoAgenda.Chorister = age.Chorister.String
	protoAgenda.Newsletter = age.Newsletter.String
	protoAgenda.IntermediateHymn = age.IntermediateHymn.Int64
	protoAgenda.MusicalNumber = age.MusicalNumber.String
	protoAgenda.ClosingHymn = age.ClosingHymn.Int64
	protoAgenda.Invocation = age.Invocation.String
	protoAgenda.Benediction = age.Benediction.String
	protoAgenda.WardBusiness = age.WardBusiness.Bool
	protoAgenda.BishopBusiness = age.BishopBusiness.Bool
	protoAgenda.LetterRead = age.LetterRead.Bool
	protoAgenda.StakeBusiness = age.StakeBusiness.Bool
	protoAgenda.Stake = age.Stake.String
	protoAgenda.NewMembers = age.NewMembers.Bool
	protoAgenda.Ordinance = age.Ordinance.Bool
	protoAgenda.Fastsunday = age.Fastsunday.Bool
	protoAgenda.AgendaPublished = age.AgendaPublished.Bool
	protoAgenda.ProgramPublished = age.ProgramPublished.Bool
	return &protoAgenda, nil
}

func translateIn(in *p.Agenda) (*Agenda, error) {
	age := Agenda{}
	age.Date = in.Date
	age.Presiding.Scan(in.Presiding)
	age.Conducting.Scan(in.Conducting)
	age.Organist.Scan(in.Organist)
	age.Chorister.Scan(in.Chorister)
	age.Newsletter.Scan(in.Newsletter)
	age.IntermediateHymn.Scan(in.IntermediateHymn)
	age.MusicalNumber.Scan(in.MusicalNumber)
	age.ClosingHymn.Scan(in.ClosingHymn)
	age.Invocation.Scan(in.Invocation)
	age.Benediction.Scan(in.Benediction)
	age.WardBusiness.Scan(in.WardBusiness)
	age.BishopBusiness.Scan(in.BishopBusiness)
	age.LetterRead.Scan(in.LetterRead)
	age.StakeBusiness.Scan(in.StakeBusiness)
	age.Stake.Scan(in.Stake)
	age.NewMembers.Scan(in.NewMembers)
	age.Ordinance.Scan(in.Ordinance)
	age.Fastsunday.Scan(in.Fastsunday)
	age.AgendaPublished.Scan(in.AgendaPublished)
	age.ProgramPublished.Scan(in.ProgramPublished)
	return &age, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(age *Agenda) (*p.Agenda, error) {
	protoAgenda := p.Agenda{}
	outBytes, err := json.Marshal(age)
	if err != nil {
		return &protoAgenda, ae.GeneralError("Unable to encode from Agenda", err)
	}
	err = json.Unmarshal(outBytes, &protoAgenda)
	if err != nil {
		return &protoAgenda, ae.GeneralError("Unable to decode to proto.Agenda", err)
	}
	return &protoAgenda, nil
}

func translateJsonIn(in *p.Agenda) (*Agenda, error) {
	age := Agenda{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &age, ae.GeneralError("Unable to encode from proto.Agenda", err)
	}
	err = json.Unmarshal(outBytes, &age)
	if err != nil {
		return &age, ae.GeneralError("Unable to decode to Agenda", err)
	}
	return &age, nil
}
