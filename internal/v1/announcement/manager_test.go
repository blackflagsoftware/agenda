package announcement

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestManagerAnnouncement_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAnnouncement := NewMockDataAnnouncementAdapter(ctrl)

	tests := []struct {
		name    string
		ann     *Announcement
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAnnouncement{dataAnnouncement: mockDataAnnouncement}
			err := m.Get(tt.ann)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAnnouncement.Get().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAnnouncement.Get().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerAnnouncement_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAnnouncement := NewMockDataAnnouncementAdapter(ctrl)

	tests := []struct {
		name    string
		ann     *Announcement
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			&Announcement{Date: null.NewString("a", true), Message: null.NewString("a", true), Pulpit: null.NewBool(true, true)},
			false,
			[]*gomock.Call{mockDataAnnouncement.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAnnouncement{dataAnnouncement: mockDataAnnouncement}
			err := m.Post(tt.ann)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAnnouncement.Create().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAnnouncement.Create().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerAnnouncement_Patch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAnnouncement := NewMockDataAnnouncementAdapter(ctrl)

	tests := []struct {
		name    string
		body    Announcement
		wantErr bool
		calls   []*gomock.Call
	}{
		{
			"successful",
			Announcement{},
			false,
			[]*gomock.Call{
				mockDataAnnouncement.EXPECT().Read(gomock.Any()).Return(nil),
				mockDataAnnouncement.EXPECT().Update(gomock.Any()).Return(nil),
			},
		},
		{
			"invalid id",
			Announcement{},
			true,
			[]*gomock.Call{
				mockDataAnnouncement.EXPECT().Read(gomock.Any()).Return(fmt.Errorf("missing record")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAnnouncement{dataAnnouncement: mockDataAnnouncement}
			err := m.Patch(tt.body)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAnnouncement.Update().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAnnouncement.Update().%s => expected error: got nil", tt.name)
			}
		})
	}
}

func TestManagerAnnouncement_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDataAnnouncement := NewMockDataAnnouncementAdapter(ctrl)

	tests := []struct {
		name    string
		ann     *Announcement
		wantErr bool
		calls   []*gomock.Call
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagerAnnouncement{dataAnnouncement: mockDataAnnouncement}
			err := m.Delete(tt.ann)
			if !tt.wantErr {
				assert.Nil(t, err, "ManagerAnnouncement.Delete().%s => expected not error; got: %s", tt.name, err)
			}
			if tt.wantErr {
				assert.NotNil(t, err, "ManagerAnnouncement.Delete().%s => expected error: got nil", tt.name)
			}
		})
	}
}
