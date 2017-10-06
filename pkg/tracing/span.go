package tracing

import (
	"sync"

	"time"

	"github.com/influxdata/influxdb/pkg/tracing/field"
	"github.com/influxdata/influxdb/pkg/tracing/label"
)

type Span struct {
	tracer *Trace
	mu     sync.Mutex
	raw    RawSpan
}

type StartSpanOption interface {
	Apply(*Span)
}

type StartTime time.Time

func (t StartTime) Apply(s *Span) {
	s.raw.Start = time.Time(t)
}

// StartSpan creates a new child span.
func (s *Span) StartSpan(name string, opt ...StartSpanOption) *Span {
	return s.tracer.startSpan(name, s.raw.Context, opt)
}

func (s *Span) Context() SpanContext {
	return s.raw.Context
}

func (s *Span) SetLabels(args ...string) {
	s.mu.Lock()
	s.raw.Labels = label.Labels(args...)
	s.mu.Unlock()
}

func (s *Span) MergeLabels(args ...string) {
	ls := label.Labels(args...)
	s.mu.Lock()
	s.raw.Labels.Merge(ls)
	s.mu.Unlock()
}

func (s *Span) SetFieldSet(set field.FieldSet) {
	s.mu.Lock()
	s.raw.Fields = set
	s.mu.Unlock()
}

func (s *Span) MergeFieldSet(set field.FieldSet) {
	s.mu.Lock()
	s.raw.Fields.Merge(set)
	s.mu.Unlock()
}

func (s *Span) Finish() {
	s.mu.Lock()
	s.tracer.addRawSpan(s.raw)
	s.mu.Unlock()
}

func (s *Span) Tree() *TreeNode {
	return s.tracer.TreeFrom(s.raw.Context.SpanID)
}
