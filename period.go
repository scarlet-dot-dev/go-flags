package flags

import (
	"fmt"

	"github.com/senseyeio/duration"
)

// PeriodFlag implements pflag.Value for ISO-8601 compatible duration strings.
type PeriodFlag struct {
	target *duration.Duration
}

// NewPeriod returns a flag.Value implementation for parsing flags as ISO-8601
// duration values.
func NewPeriod(target *duration.Duration) *PeriodFlag {
	return &PeriodFlag{target: target}
}

// String implements flag.Value.
func (f *PeriodFlag) String() string {
	return f.target.String()
}

// Set implements flag.Value.
func (f *PeriodFlag) Set(value string) error {
	p, err := duration.ParseISO8601(value)
	if err != nil {
		return fmt.Errorf(
			"unable to parse [%s] as ISO-8601 Duration: %s",
			value,
			err.Error(),
		)
	}

	*f.target = p
	return nil
}

// Type implements pflag.Value.
func (f *PeriodFlag) Type() string {
	return "period"
}
