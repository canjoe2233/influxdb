package tracing

import (
	"time"

	"github.com/influxdata/influxdb/pkg/tracing/field"
	"github.com/influxdata/influxdb/pkg/tracing/label"
)

type RawSpan struct {
	Context      SpanContext
	ParentSpanID uint64
	Name         string
	Start        time.Time
	Labels       label.LabelSet
	Fields       field.FieldSet
}
