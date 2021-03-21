package flags

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/pflag"
)

// UUIDFlag implements a flag.Value for uuid.UUID values.
type UUIDFlag struct {
	target *uuid.UUID
}

// ensure interface
var _ flag.Value = (*TimeFlag)(nil)
var _ pflag.Value = (*TimeFlag)(nil)

// NewUUID returns a flag.Value implementation for parsing flags as uuid.UUID
// values.
func NewUUID(target *uuid.UUID) *UUIDFlag {
	return &UUIDFlag{target: target}
}

// String implements flag.Value.
func (f *UUIDFlag) String() string {
	return f.target.String()
}

// Set implements flag.Value.
func (f *UUIDFlag) Set(value string) error {
	id, err := uuid.Parse(value)
	if err != nil {
		return fmt.Errorf("unable to parse uuid: %s", err.Error())
	}

	*f.target = id
	return nil
}

// Type implements pflag.Value.
func (f *UUIDFlag) Type() string {
	return "UUID"
}
