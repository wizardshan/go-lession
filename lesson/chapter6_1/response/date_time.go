package response

import (
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte("\"\""), nil
		//return []byte("null"), nil
	}
	formatted := fmt.Sprintf(`"%s"`, dt.Time.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (dt *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*dt = DateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
