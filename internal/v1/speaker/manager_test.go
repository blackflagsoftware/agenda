package speaker

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerSpeaker_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataSpeaker := NewMockDataSpeakerAdapter(ctrl)

	tests := []struct {
		name    string
		spe     *Speaker
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerSpeaker{dataSpeaker: mockDataSpeaker}
			err := m.Get(tt.spe)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerSpeaker.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerSpeaker.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerSpeaker_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataSpeaker := NewMockDataSpeakerAdapter(ctrl)

	tests := []struct {
		name    string
		spe     *Speaker
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Speaker{Name: null.NewString("a", true), Date: null.NewString("a", true), Position: null.NewString("a", true)},
			false,
			[]*gomock.Call{mockDataSpeaker.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerSpeaker{dataSpeaker: mockDataSpeaker}
			err := m.Post(tt.spe)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerSpeaker.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerSpeaker.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerSpeaker_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataSpeaker := NewMockDataSpeakerAdapter(ctrl)

	tests := []struct {
		name    string
		body    Speaker
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Speaker{},
			false,
			[]*gomock.Call{
				mockDataSpeaker.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataSpeaker.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Speaker{},
			true,
			[]*gomock.Call{
				mockDataSpeaker.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerSpeaker{dataSpeaker: mockDataSpeaker}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerSpeaker.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerSpeaker.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerSpeaker_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataSpeaker := NewMockDataSpeakerAdapter(ctrl)

	tests := []struct {
		name    string
		spe     *Speaker
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerSpeaker{dataSpeaker: mockDataSpeaker}
			err := m.Delete(tt.spe)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerSpeaker.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerSpeaker.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
