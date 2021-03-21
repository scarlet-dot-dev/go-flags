package flags

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// EnumFlag implements a flag.Value with a fixed set of valid values.
//
// Based on the enumflag implementation in HashiCorp's Packer.
type EnumFlag struct {
	target  *string
	options []string
}

// ensure interface
var _ flag.Value = (*EnumFlag)(nil)
var _ pflag.Value = (*EnumFlag)(nil)

// NewEnum returns a flag.Value implementation for parsing flags with a
// one-of-a-set value.
func NewEnum(target *string, options ...string) *EnumFlag {
	return &EnumFlag{target: target, options: options}
}

// WithDefault lets the user set a default value.
func (f *EnumFlag) WithDefault(val string) *EnumFlag {
	if err := f.Set(val); err != nil {
		panic(err.Error())
	}
	return f
}

// String implements flag.Value.
func (f *EnumFlag) String() string {
	return *f.target
}

// Set implements flag.Value.
func (f *EnumFlag) Set(value string) error {
	for _, v := range f.options {
		if v == value {
			*f.target = value
			return nil
		}
	}

	return fmt.Errorf(
		"expected one of %q, got [%s] instead",
		f.options,
		value,
	)
}

// Type implements pflag.Value.
func (f *EnumFlag) Type() string {
	return fmt.Sprintf("enum[%s]", strings.Join(f.options, ","))
}
