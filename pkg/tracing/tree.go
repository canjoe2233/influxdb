package tracing

import (
	"github.com/influxdata/influxdb/pkg/tracing/field"
	"github.com/xlab/treeprint"
)

// A Visitor's Visit method is invoked for each node encountered by Walk.
// If the result of Visit is not nil, Walk visits each of the children.
type Visitor interface {
	Visit(*TreeNode) Visitor
}

type TreeNode struct {
	Raw      RawSpan
	Children []*TreeNode
}

func (t *TreeNode) String() string {
	tv := newTreeVisitor()
	Walk(tv, t)
	return tv.root.String()
}

func Walk(v Visitor, node *TreeNode) {
	if v = v.Visit(node); v == nil {
		return
	}

	for _, c := range node.Children {
		Walk(v, c)
	}
}

type treeVisitor struct {
	root  treeprint.Tree
	trees []treeprint.Tree
}

func newTreeVisitor() *treeVisitor {
	t := treeprint.New()
	return &treeVisitor{root: t, trees: []treeprint.Tree{t}}
}

func (v *treeVisitor) Visit(n *TreeNode) Visitor {
	t := v.trees[len(v.trees)-1].AddBranch(n.Raw.Name)
	v.trees = append(v.trees, t)

	if labels := n.Raw.Labels; labels.Len() > 0 {
		l := t.AddBranch("labels")
		n.Raw.Labels.ForEach(func(k, v string) {
			l.AddNode(k + ": " + v)
		})
	}

	n.Raw.Fields.ForEach(func(v field.Field) {
		t.AddNode(v.String())
	})

	for _, cn := range n.Children {
		Walk(v, cn)
	}

	v.trees[len(v.trees)-1] = nil
	v.trees = v.trees[:len(v.trees)-1]

	return nil
}
