package event

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	Id() uuid.UUID
	Name() string
	OcurredOn() time.Time
}
