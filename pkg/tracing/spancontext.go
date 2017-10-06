package tracing

import (
	"github.com/gogo/protobuf/proto"
	"github.com/influxdata/influxdb/pkg/tracing/wire"
)

type SpanContext struct {
	TraceID uint64
	SpanID  uint64
}

func (s SpanContext) MarshalBinary() ([]byte, error) {
	ws := wire.SpanContext(s)
	return proto.Marshal(&ws)
}

func (s *SpanContext) UnmarshalBinary(data []byte) error {
	var ws wire.SpanContext
	err := proto.Unmarshal(data, &ws)
	if err == nil {
		*s = SpanContext(ws)
	}
	return err
}
