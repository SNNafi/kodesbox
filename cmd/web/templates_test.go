package main

import (
	"kodesbox.snnafi.dev/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			"Local",
			time.Date(2024, 7, 1, 10, 15, 0, 0, time.Local),
			"01 Jul 2024 at 10:15",
		},
		{
			"Empty",
			time.Time{},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, humanDate(tt.tm), tt.want)
		})
	}

}
