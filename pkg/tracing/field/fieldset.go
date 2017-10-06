package field

import "sort"

type FieldSet struct {
	list []Field
}

// Len returns the number of labels in the set.
func (fs FieldSet) Len() int { return len(fs.list) }

func (fs FieldSet) Fields() []Field { return fs.list }

// ForEach calls fn for each label in key order, passing the key and value as arguments.
func (fs FieldSet) ForEach(fn func(v Field)) {
	for _, l := range fs.list {
		fn(l)
	}
}

// Merge merges other with the current set, replacing any matching keys from other.
func (fs *FieldSet) Merge(other FieldSet) {
	var list []Field
	i, j := 0, 0
	for i < len(fs.list) && j < len(other.list) {
		if fs.list[i].key < other.list[j].key {
			list = append(list, fs.list[i])
			i++
		} else if fs.list[i].key > other.list[j].key {
			list = append(list, other.list[j])
			j++
		} else {
			// equal, then "other" replaces existing key
			list = append(list, other.list[j])
			i++
			j++
		}
	}

	if i < len(fs.list) {
		list = append(list, fs.list[i:]...)
	} else if j < len(other.list) {
		list = append(list, other.list[j:]...)
	}

	fs.list = list
}

// Labels takes an even number of strings representing key-value pairs
// and makes a LabelSet containing them.
// A label overwrites a prior label with the same key.
func Fields(args ...Field) FieldSet {
	fields := FieldSet{list: args}
	sort.Slice(fields.list, func(i, j int) bool {
		return fields.list[i].key < fields.list[j].key
	})

	return fields
}
