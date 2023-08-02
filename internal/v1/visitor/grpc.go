package visitor

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	VisitorGrpc struct {
		p.UnimplementedVisitorServiceServer
		managerVisitor ManagerVisitorAdapter
	}
)

func NewVisitorGrpc(mvis ManagerVisitorAdapter) *VisitorGrpc {
	return &VisitorGrpc{managerVisitor: mvis}
}

func (a *VisitorGrpc) GetVisitor(ctx context.Context, in *p.VisitorIDIn) (*p.VisitorResponse, error) {
	result := &p.Result{Success: false}
	response := &p.VisitorResponse{Result: result}
	vis := &Visitor{Id: int(in.Id)}
	if err := a.managerVisitor.Get(vis); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Visitor, err = translateOut(vis)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *VisitorGrpc) SearchVisitor(ctx context.Context, in *p.Visitor) (*p.VisitorRepeatResponse, error) {
	vistorsParam := VisitorParam{}
	result := &p.Result{Success: false}
	response := &p.VisitorRepeatResponse{Result: result}
	viss := &[]Visitor{}
	if _, err := a.managerVisitor.Search(viss, vistorsParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *viss {
		protoVisitor, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Visitor = append(response.Visitor, protoVisitor)
	}
	response.Result.Success = true
	return response, nil
}

func (a *VisitorGrpc) PostVisitor(ctx context.Context, in *p.Visitor) (*p.VisitorResponse, error) {
	result := &p.Result{Success: false}
	response := &p.VisitorResponse{Result: result}
	vis, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerVisitor.Post(vis); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Visitor, errTranslate = translateOut(vis)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *VisitorGrpc) PatchVisitor(ctx context.Context, in *p.Visitor) (*p.Result, error) {
	response := &p.Result{Success: false}
	vis, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerVisitor.Patch(*vis); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *VisitorGrpc) DeleteVisitor(ctx context.Context, in *p.VisitorIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	vis := &Visitor{Id: int(in.Id)}
	if err := a.managerVisitor.Delete(vis); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(vis *Visitor) (*p.Visitor, error) {
	protoVisitor := p.Visitor{}
	protoVisitor.Id = int64(vis.Id)
	protoVisitor.Date = vis.Date.String
	protoVisitor.Name = vis.Name.String
	return &protoVisitor, nil
}

func translateIn(in *p.Visitor) (*Visitor, error) {
	vis := Visitor{}
	vis.Id = int(in.Id)
	vis.Date.Scan(in.Date)
	vis.Name.Scan(in.Name)
	return &vis, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(vis *Visitor) (*p.Visitor, error) {
	protoVisitor := p.Visitor{}
	outBytes, err := json.Marshal(vis)
	if err != nil {
		return &protoVisitor, ae.GeneralError("Unable to encode from Visitor", err)
	}
	err = json.Unmarshal(outBytes, &protoVisitor)
	if err != nil {
		return &protoVisitor, ae.GeneralError("Unable to decode to proto.Visitor", err)
	}
	return &protoVisitor, nil
}

func translateJsonIn(in *p.Visitor) (*Visitor, error) {
	vis := Visitor{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &vis, ae.GeneralError("Unable to encode from proto.Visitor", err)
	}
	err = json.Unmarshal(outBytes, &vis)
	if err != nil {
		return &vis, ae.GeneralError("Unable to decode to Visitor", err)
	}
	return &vis, nil
}
