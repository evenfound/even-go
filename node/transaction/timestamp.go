package transaction

import (
	"encoding/json"
	"fmt"
	"time"
)

const timeLayout = time.RubyDate

type timestamp time.Time

// String satisfies interface Stringer.
func (t *timestamp) String() string {
	return time.Time(*t).String()
}

// UnixNanoStr returns string representation of UnixNano.
func (t *timestamp) UnixNanoStr() string {
	return fmt.Sprintf("%d", uint64(time.Time(*t).UnixNano()))
}

// MarshalJSON satisfies interface Marshaler.
func (t *timestamp) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, time.Time(*t).Format(timeLayout))
	return []byte(s), nil
}

// UnmarshalJSON satisfies interface Unmarshaler.
func (t *timestamp) UnmarshalJSON(b []byte) error {
	var str string // get rid of quotes
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	r, err := time.Parse(timeLayout, str)
	if err != nil {
		return err
	}

	*t = timestamp(r)
	return nil
}
