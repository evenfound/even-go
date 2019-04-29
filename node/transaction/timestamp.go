package transaction

import (
	"encoding/json"
	"fmt"
	"time"
)

const timeLayout = time.RubyDate

type timestamp time.Time

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
