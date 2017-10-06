package label

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelSet_Merge(t *testing.T) {
	cases := []struct {
		n    string
		l, r LabelSet
		exp  LabelSet
	}{
		{
			n:   "no matching keys",
			l:   Labels("k05", "v05", "k03", "v03", "k01", "v01"),
			r:   Labels("k02", "v02", "k04", "v04", "k00", "v00"),
			exp: Labels("k05", "v05", "k03", "v03", "k01", "v01", "k02", "v02", "k04", "v04", "k00", "v00"),
		},
		{
			n:   "multiple matching keys",
			l:   Labels("k05", "v05", "k03", "v03", "k01", "v01"),
			r:   Labels("k02", "v02", "k03", "v03a", "k05", "v05a"),
			exp: Labels("k05", "v05a", "k03", "v03a", "k01", "v01", "k02", "v02"),
		},
		{
			n:   "source empty",
			l:   Labels(),
			r:   Labels("k02", "v02", "k04", "v04", "k00", "v00"),
			exp: Labels("k02", "v02", "k04", "v04", "k00", "v00"),
		},
		{
			n:   "other empty",
			l:   Labels("k02", "v02", "k04", "v04", "k00", "v00"),
			r:   Labels(),
			exp: Labels("k02", "v02", "k04", "v04", "k00", "v00"),
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
