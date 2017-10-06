package metrics

type groupDesc struct {
	Name string
	id   GID
}

type metricType int

const (
	counterMetricType metricType = iota
	timerMetricType
)

type desc struct {
	Name string
	mt   metricType
	gid  GID
	id   ID
}

type descOption func(*desc)

// WithGroup assigns the associated measurement to group identified by gid.
func WithGroup(gid GID) descOption {
	return func(d *desc) {
		d.gid = gid
	}
}

func newDesc(name string, opts ...descOption) *desc {
	desc := &desc{Name: name}
	for _, o := range opts {
		o(desc)
	}
	return desc
}

const (
	idMask   = (1 << 32) - 1
	gidShift = 32
)

type (
	GID uint32
	ID  uint64
)

func newID(id int, gid GID) ID {
	return ID(gid)<<gidShift + (ID(id) & idMask)
}

func (id ID) id() int {
	return int(id & idMask)
}

func (id ID) gid() int {
	return int(id >> gidShift)
}

func (id *ID) setGID(gid GID) {
	*id |= ID(gid) << gidShift
}