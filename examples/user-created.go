package examples

import (
	"time"

	"github.com/google/uuid"
)

type UserCreated struct {
	id        uuid.UUID
	name      string
	ocurredOn time.Time
	userEmail string
}

var USER_CREATED = "USER.CREATED"

func InitUserCreated(userEmail string) UserCreated {
	return UserCreated{
		id:        uuid.New(),
		name:      USER_CREATED,
		ocurredOn: time.Now(),
		userEmail: userEmail,
	}
}

func (uc UserCreated) Id() uuid.UUID {
	return uc.id
}

func (uc UserCreated) Name() string {
	return uc.name
}

func (uc UserCreated) OcurredOn() time.Time {
	return uc.ocurredOn
}

func (uc UserCreated) UserEmail() string {
	return uc.userEmail
}
