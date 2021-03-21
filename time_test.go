package flags

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	cases := []struct {
		value  string
		format string
		accept bool
	}{
		{
			"2020-12-12T12:12:12Z",
			time.RFC3339,
			true,
		},
	}

	for i := range cases {
		c := cases[i]

		var out time.Time
		flag := NewTime(&out, c.format)

		if !c.accept {
			require.Error(t, flag.Set(c.value))
			continue // reject, go to next test
		}

		require.NoError(t, flag.Set(c.value))
	}
}
