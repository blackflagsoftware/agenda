package roles

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerRoles_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoles := NewMockDataRolesAdapter(ctrl)

	tests := []struct {
		name    string
		rol     *Roles
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoles{dataRoles: mockDataRoles}
			err := m.Get(tt.rol)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoles.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoles.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRoles_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoles := NewMockDataRolesAdapter(ctrl)

	tests := []struct {
		name    string
		rol     *Roles
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Roles{Name: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataRoles.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoles{dataRoles: mockDataRoles}
			err := m.Post(tt.rol)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoles.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoles.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRoles_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoles := NewMockDataRolesAdapter(ctrl)

	tests := []struct {
		name    string
		body    Roles
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Roles{},
			false,
			[]*gomock.Call{
				mockDataRoles.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataRoles.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Roles{},
			true,
			[]*gomock.Call{
				mockDataRoles.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoles{dataRoles: mockDataRoles}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoles.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoles.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRoles_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoles := NewMockDataRolesAdapter(ctrl)

	tests := []struct {
		name    string
		rol     *Roles
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoles{dataRoles: mockDataRoles}
			err := m.Delete(tt.rol)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoles.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoles.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
