package newmember

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerNewMember_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataNewMember := NewMockDataNewMemberAdapter(ctrl)

	tests := []struct {
		name    string
		new     *NewMember
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerNewMember{dataNewMember: mockDataNewMember}
			err := m.Get(tt.new)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerNewMember.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerNewMember.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerNewMember_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataNewMember := NewMockDataNewMemberAdapter(ctrl)

	tests := []struct {
		name    string
		new     *NewMember
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&NewMember{Date: null.NewString("a", true), FamilyName: null.NewString("a", true), Names: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataNewMember.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerNewMember{dataNewMember: mockDataNewMember}
			err := m.Post(tt.new)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerNewMember.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerNewMember.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerNewMember_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataNewMember := NewMockDataNewMemberAdapter(ctrl)

	tests := []struct {
		name    string
		body    NewMember
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			NewMember{},
			false,
			[]*gomock.Call{
				mockDataNewMember.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataNewMember.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			NewMember{},
			true,
			[]*gomock.Call{
				mockDataNewMember.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerNewMember{dataNewMember: mockDataNewMember}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerNewMember.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerNewMember.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerNewMember_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataNewMember := NewMockDataNewMemberAdapter(ctrl)

	tests := []struct {
		name    string
		new     *NewMember
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerNewMember{dataNewMember: mockDataNewMember}
			err := m.Delete(tt.new)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerNewMember.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerNewMember.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
