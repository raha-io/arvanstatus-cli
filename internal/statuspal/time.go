package statuspal

import (
	"strings"
	"time"
)

// tehranLoc is the status page's configured timezone. The StatusPal API emits
// naive datetimes (no Z, no offset) in Asia/Tehran local time. time.Parse
// would silently interpret them as UTC.
var tehranLoc = func() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		return time.FixedZone("IRST", int((3*time.Hour + 30*time.Minute).Seconds()))
	}
	return loc
}()

const statuspalLayout = "2006-01-02T15:04:05"

// LocalTime wraps time.Time with a JSON unmarshaller that interprets the
// statuspal naive format in Asia/Tehran local time.
type LocalTime struct {
	time.Time
}

func (lt *LocalTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		return nil
	}
	t, err := time.ParseInLocation(statuspalLayout, s, tehranLoc)
	if err != nil {
		return err
	}
	lt.Time = t
	return nil
}
