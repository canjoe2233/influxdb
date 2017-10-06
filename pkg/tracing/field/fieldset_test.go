package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldSet_Merge(t *testing.T) {
	cases := []struct {
		n    string
		l, r FieldSet
		exp  FieldSet
	}{
		{
			n:   "no matching keys",
			l:   Fields(String("k05", "v05"), String("k03", "v03"), String("k01", "v01")),
			r:   Fields(String("k02", "v02"), String("k04", "v04"), String("k00", "v00")),
			exp: Fields(String("k05", "v05"), String("k03", "v03"), String("k01", "v01"), String("k02", "v02"), String("k04", "v04"), String("k00", "v00")),
		},
		{
			n:   "multiple matching keys",
			l:   Fields(String("k05", "v05"), String("k03", "v03"), String("k01", "v01")),
			r:   Fields(String("k02", "v02"), String("k03", "v03a"), String("k05", "v05a")),
			exp: Fields(String("k05", "v05a"), String("k03", "v03a"), String("k01", "v01"), String("k02", "v02")),
		},
		{
			n:   "source empty",
			l:   Fields(),
			r:   Fields(String("k02", "v02"), String("k04", "v04"), String("k00", "v00")),
			exp: Fields(String("k02", "v02"), String("k04", "v04"), String("k00", "v00")),
		},
		{
			n:   "other empty",
			l:   Fields(String("k02", "v02"), String("k04", "v04"), String("k00", "v00")),
			r:   Fields(),
			exp: Fields(String("k02", "v02"), String("k04", "v04"), String("k00", "v00")),
		},
	}

	for _, tc := range cases {
		t.Run(tc.n, func(t *testing.T) {
			l := tc.l
			l.Merge(tc.r)
			assert.Equal(t, tc.exp.list, l.list)
		})
	}
}
