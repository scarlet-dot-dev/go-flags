package flags

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUUID(t *testing.T) {
	cases := []struct {
		value  string
		accept bool
	}{
		{"7f7529fe-8a48-11eb-8dcd-0242ac130003", true}, // v1
		{"99910afb-d49c-4d0d-a42d-f3436c2792d8", true}, // v4
		{"anything else", false},
		{"", false},
	}

	for i := range cases {
		c := cases[i]

		var out uuid.UUID
		flag := NewUUID(&out)

		if !c.accept {
			require.Error(t, flag.Set(c.value))
			continue // reject, go to next test
		}

		require.NoError(t, flag.Set(c.value))
		require.Equal(t, out.String(), c.value)
	}
}
