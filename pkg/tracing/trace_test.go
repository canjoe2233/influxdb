package tracing_test

import (
	"fmt"
	"testing"

	"time"

	"github.com/influxdata/influxdb/pkg/tracing"
	"github.com/stretchr/testify/assert"
)

type testVisitor struct{}

func (v *testVisitor) Visit(node *tracing.TreeNode) tracing.Visitor {
	fmt.Println(node.Raw.Name)
	node.Raw.Labels.ForEach(func(k, v string) {
		fmt.Println("  ", k, ":", v)
	})
	return v
}

func TestNewTrace(t *testing.T) {
	tr, s := tracing.NewTrace("foo")

	ch1a := s.StartSpan("ch1a")
	ch1a.SetLabels("ch1a-k0", "v0")
	ch1a2a := ch1a.StartSpan("ch1a-2a")
	ch1a2a.Finish()
	ch1a.Finish()

	ch1b := s.StartSpan("ch1b")
	ch1b.SetLabels("ch1b-k0", "v0")
	ch1b.Finish()

	s.Finish()

	n := tr.Tree()
	tracing.Walk(&testVisitor{}, n)
}

func startTime(v int64) tracing.StartTime {
	return tracing.StartTime(time.Unix(v, 0))
}

func baseTrace() *tracing.Trace {
	tr, s := tracing.NewTrace("foo", startTime(0))
	ch1a := s.StartSpan("ch1a", startTime(1))
	ch1a.SetLabels("ch1a-k0", "v0")
	ch1a2a := ch1a.StartSpan("ch1a-2a", startTime(2))
	ch1a2a.SetLabels("ch1a-2a-k0", "v0")
	ch1a2a.Finish()
	ch1a.Finish()
	ch1b := s.StartSpan("ch1b", startTime(3))
	ch1b.SetLabels("ch1b-k0", "v0")
	ch1b.Finish()
	s.Finish()
	return tr
}

func TestTrace_MarshalUnmarshalBinary(t *testing.T) {
	tr, s := tracing.NewTrace("foo", startTime(0))
	ch1a := s.StartSpan("ch1a", startTime(1))
	ch1a.SetLabels("ch1a-k0", "v0")
	ch1a2a := ch1a.StartSpan("ch1a-2a", startTime(2))
	ch1a2a.SetLabels("ch1a-2a-k0", "v0")
	ch1a2a.Finish()
	ch1a.Finish()
	ch1b := s.StartSpan("ch1b", startTime(3))
	ch1b.SetLabels("ch1b-k0", "v0")
	ch1b.Finish()
	s.Finish()

	d, err := tr.MarshalBinary()
	assert.NoError(t, err)

	var tr2 tracing.Trace
	err = tr2.UnmarshalBinary(d)
	assert.NoError(t, err)

	assert.Equal(t, tr.Tree().String(), tr2.Tree().String())
}

func TestTrace_Merge(t *testing.T) {
	tr, s := tracing.NewTrace("foo", startTime(0))
	ch1a := s.StartSpan("ch1a", startTime(1))
	ch1a.SetLabels("ch1a-k0", "v0")

	tr2, ch1a2a := tracing.NewTraceFromSpan("ch1a-2a", ch1a.Context(), startTime(2))
	ch1a2a.SetLabels("ch1a-2a-k0", "v0")
	ch1a2a.Finish()

	tr.Merge(tr2)

	ch1a.Finish()
	ch1b := s.StartSpan("ch1b", startTime(3))
	ch1b.SetLabels("ch1b-k0", "v0")
	ch1b.Finish()
	s.Finish()

	assert.Equal(t, baseTrace().Tree().String(), tr.Tree().String())
}
