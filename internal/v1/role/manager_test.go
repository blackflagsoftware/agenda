package role

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerRole_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRole := NewMockDataRoleAdapter(ctrl)

	tests := []struct {
		name    string
		rol     *Role
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRole{dataRole: mockDataRole}
			err := m.Get(tt.rol)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRole.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRole.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRole_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRole := NewMockDataRoleAdapter(ctrl)

	tests := []struct {
		name    string
		rol     *Role
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Role{Name: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataRole.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRole{dataRole: mockDataRole}
			err := m.Post(tt.rol)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRole.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRole.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRole_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRole := NewMockDataRoleAdapter(ctrl)

	tests := []struct {
		name    string
		body    Role
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Role{},
			false,
			[]*gomock.Call{
				mockDataRole.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataRole.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Role{},
			true,
			[]*gomock.Call{
				mockDataRole.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRole{dataRole: mockDataRole}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRole.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRole.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRole_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRole := NewMockDataRoleAdapter(ctrl)

	tests := []struct {
		name    string
		rol     *Role
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRole{dataRole: mockDataRole}
			err := m.Delete(tt.rol)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRole.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRole.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
