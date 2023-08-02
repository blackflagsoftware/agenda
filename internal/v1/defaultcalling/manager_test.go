package defaultcalling

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerDefaultCalling_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataDefaultCalling := NewMockDataDefaultCallingAdapter(ctrl)

	tests := []struct {
		name    string
		def     *DefaultCalling
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerDefaultCalling{dataDefaultCalling: mockDataDefaultCalling}
			err := m.Get(tt.def)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerDefaultCalling.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerDefaultCalling.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerDefaultCalling_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataDefaultCalling := NewMockDataDefaultCallingAdapter(ctrl)

	tests := []struct {
		name    string
		def     *DefaultCalling
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&DefaultCalling{Organist: null.NewString("a", true), Chorister: null.NewString("a", true), Newsletter: null.NewString("a", true), Stake: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataDefaultCalling.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerDefaultCalling{dataDefaultCalling: mockDataDefaultCalling}
			err := m.Post(tt.def)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerDefaultCalling.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerDefaultCalling.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerDefaultCalling_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataDefaultCalling := NewMockDataDefaultCallingAdapter(ctrl)

	tests := []struct {
		name    string
		body    DefaultCalling
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			DefaultCalling{},
			false,
			[]*gomock.Call{
				mockDataDefaultCalling.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataDefaultCalling.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			DefaultCalling{},
			true,
			[]*gomock.Call{
				mockDataDefaultCalling.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerDefaultCalling{dataDefaultCalling: mockDataDefaultCalling}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerDefaultCalling.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerDefaultCalling.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerDefaultCalling_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataDefaultCalling := NewMockDataDefaultCallingAdapter(ctrl)

	tests := []struct {
		name    string
		def     *DefaultCalling
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerDefaultCalling{dataDefaultCalling: mockDataDefaultCalling}
			err := m.Delete(tt.def)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerDefaultCalling.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerDefaultCalling.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
