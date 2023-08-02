package wardbusinesssus

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerWardBusinessSus_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessSus := NewMockDataWardBusinessSusAdapter(ctrl)

	tests := []struct {
		name    string
		wa      *WardBusinessSus
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessSus{dataWardBusinessSus: mockDataWardBusinessSus}
			err := m.Get(tt.wa)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessSus.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessSus.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerWardBusinessSus_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessSus := NewMockDataWardBusinessSusAdapter(ctrl)

	tests := []struct {
		name    string
		wa      *WardBusinessSus
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&WardBusinessSus{Calling: null.NewString("a", true), Date: null.NewString("a", true), Name: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataWardBusinessSus.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessSus{dataWardBusinessSus: mockDataWardBusinessSus}
			err := m.Post(tt.wa)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessSus.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessSus.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerWardBusinessSus_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessSus := NewMockDataWardBusinessSusAdapter(ctrl)

	tests := []struct {
		name    string
		body    WardBusinessSus
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			WardBusinessSus{},
			false,
			[]*gomock.Call{
				mockDataWardBusinessSus.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataWardBusinessSus.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			WardBusinessSus{},
			true,
			[]*gomock.Call{
				mockDataWardBusinessSus.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessSus{dataWardBusinessSus: mockDataWardBusinessSus}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessSus.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessSus.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerWardBusinessSus_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessSus := NewMockDataWardBusinessSusAdapter(ctrl)

	tests := []struct {
		name    string
		wa      *WardBusinessSus
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessSus{dataWardBusinessSus: mockDataWardBusinessSus}
			err := m.Delete(tt.wa)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessSus.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessSus.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
