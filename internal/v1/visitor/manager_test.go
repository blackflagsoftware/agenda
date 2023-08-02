package visitor

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerVisitor_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataVisitor := NewMockDataVisitorAdapter(ctrl)

	tests := []struct {
		name    string
		vis     *Visitor
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerVisitor{dataVisitor: mockDataVisitor}
			err := m.Get(tt.vis)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerVisitor.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerVisitor.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerVisitor_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataVisitor := NewMockDataVisitorAdapter(ctrl)

	tests := []struct {
		name    string
		vis     *Visitor
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Visitor{Date: null.NewString("a", true), Name: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataVisitor.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerVisitor{dataVisitor: mockDataVisitor}
			err := m.Post(tt.vis)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerVisitor.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerVisitor.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerVisitor_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataVisitor := NewMockDataVisitorAdapter(ctrl)

	tests := []struct {
		name    string
		body    Visitor
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Visitor{},
			false,
			[]*gomock.Call{
				mockDataVisitor.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataVisitor.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Visitor{},
			true,
			[]*gomock.Call{
				mockDataVisitor.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerVisitor{dataVisitor: mockDataVisitor}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerVisitor.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerVisitor.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerVisitor_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataVisitor := NewMockDataVisitorAdapter(ctrl)

	tests := []struct {
		name    string
		vis     *Visitor
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerVisitor{dataVisitor: mockDataVisitor}
			err := m.Delete(tt.vis)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerVisitor.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerVisitor.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
