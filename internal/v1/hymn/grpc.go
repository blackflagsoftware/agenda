package hymn

import (
	"context"
	"encoding/json"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	p "github.com/blackflagsoftware/agenda/pkg/proto"
)

type (
	HymnGrpc struct {
		p.UnimplementedHymnServiceServer
		managerHymn ManagerHymnAdapter
	}
)

func NewHymnGrpc(mhym ManagerHymnAdapter) *HymnGrpc {
	return &HymnGrpc{managerHymn: mhym}
}

func (a *HymnGrpc) GetHymn(ctx context.Context, in *p.HymnIDIn) (*p.HymnResponse, error) {
	result := &p.Result{Success: false}
	response := &p.HymnResponse{Result: result}
	hym := &Hymn{Id: int(in.Id)}
	if err := a.managerHymn.Get(hym); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var err error
	response.Hymn, err = translateOut(hym)
	if err != nil {
		return response, err
	}
	response.Result.Success = true
	return response, nil
}

func (a *HymnGrpc) SearchHymn(ctx context.Context, in *p.Hymn) (*p.HymnRepeatResponse, error) {
	hymnParam := HymnParam{}
	result := &p.Result{Success: false}
	response := &p.HymnRepeatResponse{Result: result}
	hyms := &[]Hymn{}
	if _, err := a.managerHymn.Search(hyms, hymnParam); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	for _, a := range *hyms {
		protoHymn, err := translateOut(&a)
		if err != nil {
			return response, err
		}
		response.Hymn = append(response.Hymn, protoHymn)
	}
	response.Result.Success = true
	return response, nil
}

func (a *HymnGrpc) PostHymn(ctx context.Context, in *p.Hymn) (*p.HymnResponse, error) {
	result := &p.Result{Success: false}
	response := &p.HymnResponse{Result: result}
	hym, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerHymn.Post(hym); err != nil {
		response.Result.Error = err.Error()
		return response, err
	}
	var errTranslate error
	response.Hymn, errTranslate = translateOut(hym)
	if err != nil {
		return response, errTranslate
	}
	response.Result.Success = true
	return response, nil
}

func (a *HymnGrpc) PatchHymn(ctx context.Context, in *p.Hymn) (*p.Result, error) {
	response := &p.Result{Success: false}
	hym, err := translateIn(in)
	if err != nil {
		return response, err
	}
	if err := a.managerHymn.Patch(*hym); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func (a *HymnGrpc) DeleteHymn(ctx context.Context, in *p.HymnIDIn) (*p.Result, error) {
	response := &p.Result{Success: false}
	hym := &Hymn{Id: int(in.Id)}
	if err := a.managerHymn.Delete(hym); err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Success = true
	return response, nil
}

func translateOut(hym *Hymn) (*p.Hymn, error) {
	protoHymn := p.Hymn{}
	protoHymn.Id = int64(hym.Id)
	protoHymn.Name = hym.Name.String
	return &protoHymn, nil
}

func translateIn(in *p.Hymn) (*Hymn, error) {
	hym := Hymn{}
	hym.Id = int(in.Id)
	hym.Name.Scan(in.Name)
	return &hym, nil
}

// found these are slower; deprecated; keep them, just in case
func translateJsonOut(hym *Hymn) (*p.Hymn, error) {
	protoHymn := p.Hymn{}
	outBytes, err := json.Marshal(hym)
	if err != nil {
		return &protoHymn, ae.GeneralError("Unable to encode from Hymn", err)
	}
	err = json.Unmarshal(outBytes, &protoHymn)
	if err != nil {
		return &protoHymn, ae.GeneralError("Unable to decode to proto.Hymn", err)
	}
	return &protoHymn, nil
}

func translateJsonIn(in *p.Hymn) (*Hymn, error) {
	hym := Hymn{}
	outBytes, err := json.Marshal(in)
	if err != nil {
		return &hym, ae.GeneralError("Unable to encode from proto.Hymn", err)
	}
	err = json.Unmarshal(outBytes, &hym)
	if err != nil {
		return &hym, ae.GeneralError("Unable to decode to Hymn", err)
	}
	return &hym, nil
}
