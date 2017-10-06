package label

import "sort"

type Label struct {
	Key, Value string
}

type LabelSet struct {
	list []Label
}

// Len returns the number of labels in the set.
func (ls LabelSet) Len() int { return len(ls.list) }

func (ls LabelSet) Labels() []Label { return ls.list }

// ForEach calls fn for each label in key order, passing the key and value as arguments.
func (ls LabelSet) ForEach(fn func(k, v string)) {
	for _, l := range ls.list {
		fn(l.Key, l.Value)
	}
}

// Merge merges other with the current set, replacing any matching keys from other.
func (ls *LabelSet) Merge(other LabelSet) {
	var list []Label
	i, j := 0, 0
	for i < len(ls.list) && j < len(other.list) {
		if ls.list[i].Key < other.list[j].Key {
			list = append(list, ls.list[i])
			i++
		} else if ls.list[i].Key > other.list[j].Key {
			list = append(list, other.list[j])
			j++
		} else {
			// equal, then "other" replaces existing key
			list = append(list, other.list[j])
			i++
			j++
		}
	}

	if i < len(ls.list) {
		list = append(list, ls.list[i:]...)
	} else if j < len(other.list) {
		list = append(list, other.list[j:]...)
	}

	ls.list = list
}

// Labels takes an even number of strings representing key-value pairs
// and makes a LabelSet containing them.
// A label overwrites a prior label with the same key.
func Labels(args ...string) LabelSet {
	if len(args)%2 != 0 {
		panic("uneven number of arguments to label.Labels")
	}
	labels := LabelSet{}
	for i := 0; i+1 < len(args); i += 2 {
		labels.list = append(labels.list, Label{Key: args[i], Value: args[i+1]})
	}

	sort.Slice(labels.list, func(i, j int) bool {
		return labels.list[i].Key < labels.list[j].Key
	})

	return labels
}
