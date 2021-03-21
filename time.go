package flags

import (
	"flag"
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

// TimeFlag implements a flag.Value for time.Time values.
type TimeFlag struct {
	target *time.Time
	format string
}

// ensure interface
var _ flag.Value = (*TimeFlag)(nil)
var _ pflag.Value = (*TimeFlag)(nil)

// NewTime returns a flag.Value implementation for parsing flags as time.Time
// values. The format value is used as described in time.Parse.
func NewTime(target *time.Time, format string) *TimeFlag {
	return &TimeFlag{target: target, format: format}
}

// String implements flag.Value.
func (f *TimeFlag) String() string {
	return f.target.String()
}

// Set implements flag.Value.
func (f *TimeFlag) Set(value string) error {
	t, err := time.Parse(f.format, value)
	if err != nil {
		return fmt.Errorf("unable to parse time: %s", err.Error())
	}

	*f.target = t
	return nil
}

// Type implements pflag.Value.
func (f *TimeFlag) Type() string {
	return "time"
}
