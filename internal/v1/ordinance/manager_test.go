package ordinance

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerOrdinance_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataOrdinance := NewMockDataOrdinanceAdapter(ctrl)

	tests := []struct {
		name    string
		ord     *Ordinance
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerOrdinance{dataOrdinance: mockDataOrdinance}
			err := m.Get(tt.ord)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerOrdinance.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerOrdinance.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerOrdinance_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataOrdinance := NewMockDataOrdinanceAdapter(ctrl)

	tests := []struct {
		name    string
		ord     *Ordinance
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Ordinance{Date: null.NewString("a", true), Confirmations: null.NewString("a", true), Blessings: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataOrdinance.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerOrdinance{dataOrdinance: mockDataOrdinance}
			err := m.Post(tt.ord)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerOrdinance.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerOrdinance.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerOrdinance_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataOrdinance := NewMockDataOrdinanceAdapter(ctrl)

	tests := []struct {
		name    string
		body    Ordinance
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Ordinance{},
			false,
			[]*gomock.Call{
				mockDataOrdinance.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataOrdinance.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Ordinance{},
			true,
			[]*gomock.Call{
				mockDataOrdinance.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerOrdinance{dataOrdinance: mockDataOrdinance}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerOrdinance.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerOrdinance.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerOrdinance_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataOrdinance := NewMockDataOrdinanceAdapter(ctrl)

	tests := []struct {
		name    string
		ord     *Ordinance
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerOrdinance{dataOrdinance: mockDataOrdinance}
			err := m.Delete(tt.ord)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerOrdinance.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerOrdinance.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
