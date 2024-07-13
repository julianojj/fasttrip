package domain

import (
	"time"

	"github.com/julianojj/fastrip/internal/core/exceptions"
)

type Period struct {
	Start time.Time
	End   time.Time
}

func (p *Period) Validate() error {
	today := time.Now().UTC()
	if p.Start.Before(today) || p.End.Before(today) {
		return exceptions.ErrInvalidPeriod
	}
	if p.DurationInDays() <= 0 {
		return exceptions.ErrInsufficientPeriod
	}
	return nil
}

func (p *Period) DurationInDays() int {
	return int(p.End.Sub(p.Start).Hours() / 24)
}
