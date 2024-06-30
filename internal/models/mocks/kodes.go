package mocks

import (
	"kodesbox.snnafi.dev/internal/models"
	"time"
)

var mockKode = &models.Kode{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expired: time.Now(),
}

type KodesBox struct{}

func (box *KodesBox) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (box *KodesBox) Get(id int) (*models.Kode, error) {

	switch id {
	case 1:
		return mockKode, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (box *KodesBox) Latest() ([]*models.Kode, error) {
	return []*models.Kode{mockKode}, nil
}
