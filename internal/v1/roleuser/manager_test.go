package roleuser

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerRoleUser_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoleUser := NewMockDataRoleUserAdapter(ctrl)

	tests := []struct {
		name    string
		ro      *RoleUser
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoleUser{dataRoleUser: mockDataRoleUser}
			err := m.Get(tt.ro)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoleUser.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoleUser.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRoleUser_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoleUser := NewMockDataRoleUserAdapter(ctrl)

	tests := []struct {
		name    string
		ro      *RoleUser
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&RoleUser{Name: null.NewString("a", true), Pwd: null.NewString("a", true), RoleId: null.NewInt(1, true)},
			false,
			[]*gomock.Call{mockDataRoleUser.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoleUser{dataRoleUser: mockDataRoleUser}
			err := m.Post(tt.ro)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoleUser.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoleUser.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRoleUser_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoleUser := NewMockDataRoleUserAdapter(ctrl)

	tests := []struct {
		name    string
		body    RoleUser
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			RoleUser{},
			false,
			[]*gomock.Call{
				mockDataRoleUser.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataRoleUser.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			RoleUser{},
			true,
			[]*gomock.Call{
				mockDataRoleUser.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoleUser{dataRoleUser: mockDataRoleUser}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoleUser.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoleUser.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerRoleUser_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataRoleUser := NewMockDataRoleUserAdapter(ctrl)

	tests := []struct {
		name    string
		ro      *RoleUser
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerRoleUser{dataRoleUser: mockDataRoleUser}
			err := m.Delete(tt.ro)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerRoleUser.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerRoleUser.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
