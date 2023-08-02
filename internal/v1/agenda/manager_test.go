package agenda

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerAgenda_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAgenda := NewMockDataAgendaAdapter(ctrl)

	tests := []struct {
		name    string
		age     *Agenda
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAgenda{dataAgenda: mockDataAgenda}
			err := m.Get(tt.age)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAgenda.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAgenda.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerAgenda_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAgenda := NewMockDataAgendaAdapter(ctrl)

	tests := []struct {
		name    string
		age     *Agenda
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Agenda{BishopBusiness: null.NewBool(true, true), Fastsunday: null.NewBool(true, true), AgendaPublished: null.NewBool(true, true), Presiding: null.NewString("a", true), MusicalNumber: null.NewString("a", true), ClosingHymn: null.NewInt(1, true), Benediction: null.NewString("a", true), ProgramPublished: null.NewBool(true, true), Organist: null.NewString("a", true), LetterRead: null.NewBool(true, true), Stake: null.NewString("a", true), NewMembers: null.NewBool(true, true), Newsletter: null.NewString("a", true), IntermediateHymn: null.NewInt(1, true), StakeBusiness: null.NewBool(true, true), Ordinance: null.NewBool(true, true), Conducting: null.NewString("a", true), Chorister: null.NewString("a", true), Invocation: null.NewString("a", true), WardBusiness: null.NewBool(true, true)},
			false,
			[]*gomock.Call{mockDataAgenda.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAgenda{dataAgenda: mockDataAgenda}
			err := m.Post(tt.age)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAgenda.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAgenda.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerAgenda_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAgenda := NewMockDataAgendaAdapter(ctrl)

	tests := []struct {
		name    string
		body    Agenda
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Agenda{},
			false,
			[]*gomock.Call{
				mockDataAgenda.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataAgenda.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Agenda{},
			true,
			[]*gomock.Call{
				mockDataAgenda.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAgenda{dataAgenda: mockDataAgenda}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAgenda.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAgenda.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerAgenda_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAgenda := NewMockDataAgendaAdapter(ctrl)

	tests := []struct {
		name    string
		age     *Agenda
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAgenda{dataAgenda: mockDataAgenda}
			err := m.Delete(tt.age)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAgenda.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAgenda.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
