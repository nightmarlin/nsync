package calendar

import (
	"context"
	"image/color"
	"time"

	"github.com/google/uuid"
)

type Source interface {
	Events(ctx context.Context, from, until time.Time) ([]Event, error)
}

type FreeBusy int

const (
	Free = FreeBusy(iota)
	Busy
)

type ShowFreeBusy int

const (
	Default = ShowFreeBusy(iota)
	Public
	Private
)

type Attending int

const (
	No = Attending(iota)
	YesRemote
	YesInPerson
)

// An Event is a single period of time scheduled into a calendar. They do not
// recur.
type Event struct {
	UUID uuid.UUID

	Name        string      // The name given to this event.
	Description string      // The description given to this event.
	Color       color.Color // The colour assigned to this event.
	Location    string      // Where this event will take place.

	Attending    Attending    // Whether the user is attending this event, and how.
	FreeBusy     FreeBusy     // Whether to show this event as free or busy (if Attending).
	ShowFreeBusy ShowFreeBusy // Whether to share the free/busy status for this event.

	Timezone time.Location // The timezone the event is being held in.
	Start    time.Time     // When the event starts. Times are in UTC.
	End      time.Time     // When the event ends. Times are in UTC.
}
