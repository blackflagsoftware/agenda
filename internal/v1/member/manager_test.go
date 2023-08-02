package member

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerMember_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataMember := NewMockDataMemberAdapter(ctrl)

	tests := []struct {
		name    string
		mem     *Member
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerMember{dataMember: mockDataMember}
			err := m.Get(tt.mem)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerMember.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerMember.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerMember_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataMember := NewMockDataMemberAdapter(ctrl)

	tests := []struct {
		name    string
		mem     *Member
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Member{LastPrayed: null.NewString("a", true), LastTalked: null.NewString("a", true), Active: null.NewBool(true, true), NoPrayer: null.NewBool(true, true), NoTalk: null.NewBool(true, true), FirstName: null.NewString("a", true), LastName: null.NewString("a", true), Gender: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataMember.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerMember{dataMember: mockDataMember}
			err := m.Post(tt.mem)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerMember.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerMember.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerMember_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataMember := NewMockDataMemberAdapter(ctrl)

	tests := []struct {
		name    string
		body    Member
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Member{},
			false,
			[]*gomock.Call{
				mockDataMember.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataMember.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Member{},
			true,
			[]*gomock.Call{
				mockDataMember.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerMember{dataMember: mockDataMember}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerMember.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerMember.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerMember_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataMember := NewMockDataMemberAdapter(ctrl)

	tests := []struct {
		name    string
		mem     *Member
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerMember{dataMember: mockDataMember}
			err := m.Delete(tt.mem)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerMember.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerMember.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
