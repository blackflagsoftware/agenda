package bishopbusiness

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerBishopBusiness_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataBishopBusiness := NewMockDataBishopBusinessAdapter(ctrl)

	tests := []struct {
		name    string
		bis     *BishopBusiness
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerBishopBusiness{dataBishopBusiness: mockDataBishopBusiness}
			err := m.Get(tt.bis)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerBishopBusiness.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerBishopBusiness.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerBishopBusiness_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataBishopBusiness := NewMockDataBishopBusinessAdapter(ctrl)

	tests := []struct {
		name    string
		bis     *BishopBusiness
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&BishopBusiness{Date: null.NewString("a", true), Message: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataBishopBusiness.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerBishopBusiness{dataBishopBusiness: mockDataBishopBusiness}
			err := m.Post(tt.bis)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerBishopBusiness.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerBishopBusiness.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerBishopBusiness_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataBishopBusiness := NewMockDataBishopBusinessAdapter(ctrl)

	tests := []struct {
		name    string
		body    BishopBusiness
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			BishopBusiness{},
			false,
			[]*gomock.Call{
				mockDataBishopBusiness.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataBishopBusiness.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			BishopBusiness{},
			true,
			[]*gomock.Call{
				mockDataBishopBusiness.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerBishopBusiness{dataBishopBusiness: mockDataBishopBusiness}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerBishopBusiness.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerBishopBusiness.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerBishopBusiness_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataBishopBusiness := NewMockDataBishopBusinessAdapter(ctrl)

	tests := []struct {
		name    string
		bis     *BishopBusiness
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerBishopBusiness{dataBishopBusiness: mockDataBishopBusiness}
			err := m.Delete(tt.bis)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerBishopBusiness.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerBishopBusiness.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
