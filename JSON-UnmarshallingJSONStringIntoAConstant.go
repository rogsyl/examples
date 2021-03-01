package examples

import "strings"

type Something int

const (
	UNKNOWN Something = iota
	ONE
	TWO
	THREE
)

type MyData struct {
	Name      string    `json:"name"`
	Something Something `json:"something"`
}

func (s *Something) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	switch {
	case str == "ONE":
		*s = ONE

	case str == "TWO":
		*s = TWO

	case str == "THREE":
		*s = THREE

	default:
		*s = UNKNOWN
	}
	return nil
}
