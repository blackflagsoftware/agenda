package hymn

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerHymn_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataHymn := NewMockDataHymnAdapter(ctrl)

	tests := []struct {
		name    string
		hym     *Hymn
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerHymn{dataHymn: mockDataHymn}
			err := m.Get(tt.hym)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerHymn.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerHymn.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerHymn_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataHymn := NewMockDataHymnAdapter(ctrl)

	tests := []struct {
		name    string
		hym     *Hymn
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Hymn{Name: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataHymn.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerHymn{dataHymn: mockDataHymn}
			err := m.Post(tt.hym)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerHymn.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerHymn.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerHymn_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataHymn := NewMockDataHymnAdapter(ctrl)

	tests := []struct {
		name    string
		body    Hymn
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Hymn{},
			false,
			[]*gomock.Call{
				mockDataHymn.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataHymn.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Hymn{},
			true,
			[]*gomock.Call{
				mockDataHymn.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerHymn{dataHymn: mockDataHymn}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerHymn.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerHymn.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerHymn_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataHymn := NewMockDataHymnAdapter(ctrl)

	tests := []struct {
		name    string
		hym     *Hymn
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerHymn{dataHymn: mockDataHymn}
			err := m.Delete(tt.hym)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerHymn.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerHymn.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
