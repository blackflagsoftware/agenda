package wardbusinessrel

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerWardBusinessRel_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessRel := NewMockDataWardBusinessRelAdapter(ctrl)

	tests := []struct {
		name    string
		war     *WardBusinessRel
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessRel{dataWardBusinessRel: mockDataWardBusinessRel}
			err := m.Get(tt.war)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessRel.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessRel.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerWardBusinessRel_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessRel := NewMockDataWardBusinessRelAdapter(ctrl)

	tests := []struct {
		name    string
		war     *WardBusinessRel
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&WardBusinessRel{Date: null.NewString("a", true), Name: null.NewString("a", true), Calling: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataWardBusinessRel.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessRel{dataWardBusinessRel: mockDataWardBusinessRel}
			err := m.Post(tt.war)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessRel.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessRel.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerWardBusinessRel_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessRel := NewMockDataWardBusinessRelAdapter(ctrl)

	tests := []struct {
		name    string
		body    WardBusinessRel
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			WardBusinessRel{},
			false,
			[]*gomock.Call{
				mockDataWardBusinessRel.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataWardBusinessRel.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			WardBusinessRel{},
			true,
			[]*gomock.Call{
				mockDataWardBusinessRel.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessRel{dataWardBusinessRel: mockDataWardBusinessRel}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessRel.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessRel.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerWardBusinessRel_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataWardBusinessRel := NewMockDataWardBusinessRelAdapter(ctrl)

	tests := []struct {
		name    string
		war     *WardBusinessRel
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerWardBusinessRel{dataWardBusinessRel: mockDataWardBusinessRel}
			err := m.Delete(tt.war)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerWardBusinessRel.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerWardBusinessRel.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
