package gostl

import (
	"errors"
	"fmt"

	"github.com/RajjjAryan/gostl/utils/visitor"
)

// RbTreeIterator is an iterator implementation of RbTree
type RbTreeIterator[K, V any] struct {
	node *Node[K, V]
}

// NewIterator creates a RbTreeIterator from the passed node
func NewIterator[K, V any](node *Node[K, V]) *RbTreeIterator[K, V] {
	return &RbTreeIterator[K, V]{node: node}
}

// IsValid returns true if the iterator is valid, otherwise returns false
func (iter *RbTreeIterator[K, V]) IsValid() bool {
	return iter.node != nil
}

// Next moves the pointer of the iterator to the next node, and returns itself
func (iter *RbTreeIterator[K, V]) Next() ConstIterator[V] {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev moves the pointer of the iterator to the previous node, and returns itself
func (iter *RbTreeIterator[K, V]) Prev() ConstBidIterator[V] {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Key returns the node's key of the iterator point to
func (iter *RbTreeIterator[K, V]) Key() K {
	return iter.node.Key()
}

// Value returns the node's value of the iterator point to
func (iter *RbTreeIterator[K, V]) Value() V {
	return iter.node.Value()
}

// SetValue sets the node's value of the iterator point to
func (iter *RbTreeIterator[K, V]) SetValue(val V) error {
	iter.node.SetValue(val)
	return nil
}

// Clone clones the iterator into a new RbTreeIterator
func (iter *RbTreeIterator[K, V]) Clone() ConstIterator[V] {
	return NewIterator(iter.node)
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *RbTreeIterator[K, V]) Equal(other ConstIterator[V]) bool {
	otherIter, ok := other.(*RbTreeIterator[K, V])
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}

// Color defines node color type
type Color bool

// Define node 's colors
const (
	RED   = false
	BLACK = true
)

// Node is a tree node
type Node[K, V any] struct {
	parent *Node[K, V]
	left   *Node[K, V]
	right  *Node[K, V]
	color  Color
	key    K
	value  V
}

// Key returns node's key
func (n *Node[K, V]) Key() K {
	return n.key
}

// Value returns node's value
func (n *Node[K, V]) Value() V {
	return n.value
}

// SetValue sets node's value
func (n *Node[K, V]) SetValue(val V) {
	n.value = val
}

// Next returns the Node's successor as an iterator.
func (n *Node[K, V]) Next() *Node[K, V] {
	return successor(n)
}

// Prev returns the Node's predecessor as an iterator.
func (n *Node[K, V]) Prev() *Node[K, V] {
	return presuccessor(n)
}

// successor returns the successor of the Node
func successor[K, V any](x *Node[K, V]) *Node[K, V] {
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

// presuccessor returns the presuccessor of the Node
func presuccessor[K, V any](x *Node[K, V]) *Node[K, V] {
	if x.left != nil {
		return maximum(x.left)
	}
	if x.parent != nil {
		if x.parent.right == x {
			return x.parent
		}
		for x.parent != nil && x.parent.left == x {
			x = x.parent
		}
		return x.parent
	}
	return nil
}

// minimum finds the minimum Node of subtree n.
func minimum[K any, V any](n *Node[K, V]) *Node[K, V] {
	for n.left != nil {
		n = n.left
	}
	return n
}

// maximum finds the maximum Node of subtree n.
func maximum[K any, V any](n *Node[K, V]) *Node[K, V] {
	for n.right != nil {
		n = n.right
	}
	return n
}

var ErrorNotFound = errors.New("not found")

// RbTree is a kind of self-balancing binary search tree in computer science.
// Each node of the binary tree has an extra bit, and that bit is often interpreted
// as the color (red or black) of the node. These color bits are used to ensure the tree
// remains approximately balanced during insertions and deletions.
type RbTree[K, V any] struct {
	root   *Node[K, V]
	size   int
	keyCmp Comparator[K]
}

// New creates a new RbTree
func New[K, V any](cmp Comparator[K]) *RbTree[K, V] {
	return &RbTree[K, V]{keyCmp: cmp}
}

// Clear clears the RbTree
func (t *RbTree[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

// Find finds the first node that the key is equal to the passed key, and returns its value
func (t *RbTree[K, V]) Find(key K) (V, error) {
	n := t.findFirstNode(key)
	if n != nil {
		return n.value, nil
	}
	return *new(V), ErrorNotFound
}

// FindNode the first node that the key is equal to the passed key and return it
func (t *RbTree[K, V]) FindNode(key K) *Node[K, V] {
	return t.findFirstNode(key)
}

// Compare compares two keys wrt the RbTree's key comparator
func (t *RbTree[K, V]) Compare(key1, key2 K) int {
	return t.keyCmp(key1, key2)
}

// Begin returns the node with minimum key in the RbTree
func (t *RbTree[K, V]) Begin() *Node[K, V] {
	return t.First()
}

// First returns the node with minimum key in the RbTree
func (t *RbTree[K, V]) First() *Node[K, V] {
	if t.root == nil {
		return nil
	}
	return minimum(t.root)
}

// RBegin returns the Node with maximum key in the RbTree
func (t *RbTree[K, V]) RBegin() *Node[K, V] {
	return t.Last()
}

// Last returns the Node with maximum key in the RbTree
func (t *RbTree[K, V]) Last() *Node[K, V] {
	if t.root == nil {
		return nil
	}
	return maximum(t.root)
}

// IterFirst returns the iterator of first node
func (t *RbTree[K, V]) IterFirst() *RbTreeIterator[K, V] {
	return NewIterator(t.First())
}

// IterLast returns the iterator of first node
func (t *RbTree[K, V]) IterLast() *RbTreeIterator[K, V] {
	return NewIterator(t.Last())
}

// Empty returns true if Tree is empty,otherwise returns false.
func (t *RbTree[K, V]) Empty() bool {
	return t.size == 0
}

// Size returns the size of the rbtree.
func (t *RbTree[K, V]) Size() int {
	return t.size
}

// Insert inserts a key-value pair into the RbTree.
func (t *RbTree[K, V]) Insert(key K, value V) {
	x := t.root
	var y *Node[K, V]

	for x != nil {
		y = x
		if t.keyCmp(key, x.key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node[K, V]{parent: y, color: RED, key: key, value: value}
	t.size++

	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if t.keyCmp(z.key, y.key) < 0 {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixup(z)
}

func (t *RbTree[K, V]) rbInsertFixup(z *Node[K, V]) {
	var y *Node[K, V]
	for z.parent != nil && !z.parent.color {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && !y.color {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && !y.color {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

// Delete deletes node from the RbTree
func (t *RbTree[K, V]) Delete(node *Node[K, V]) {
	z := node
	if z == nil {
		return
	}

	var x, y *Node[K, V]
	if z.left != nil && z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	xparent := y.parent
	if x != nil {
		x.parent = xparent
	}
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.key = y.key
		z.value = y.value
	}

	if y.color {
		t.rbDeleteFixup(x, xparent)
	}
	t.size--
}

func (t *RbTree[K, V]) rbDeleteFixup(x, parent *Node[K, V]) {
	var w *Node[K, V]
	for x != t.root && getColor(x) {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			x, w = t.rbFixupLeft(x, parent, w)
		} else {
			x, w = t.rbFixupRight(x, parent, w)
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (t *RbTree[K, V]) rbFixupLeft(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	w = parent.right
	if !w.color {
		w.color = BLACK
		parent.color = RED
		t.leftRotate(parent)
		w = parent.right
	}
	if getColor(w.left) && getColor(w.right) {
		w.color = RED
		x = parent
	} else {
		if getColor(w.right) {
			if w.left != nil {
				w.left.color = BLACK
			}
			w.color = RED
			t.rightRotate(w)
			w = parent.right
		}
		w.color = parent.color
		parent.color = BLACK
		if w.right != nil {
			w.right.color = BLACK
		}
		t.leftRotate(parent)
		x = t.root
	}
	return x, w
}

func (t *RbTree[K, V]) rbFixupRight(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	w = parent.left
	if !w.color {
		w.color = BLACK
		parent.color = RED
		t.rightRotate(parent)
		w = parent.left
	}
	if getColor(w.left) && getColor(w.right) {
		w.color = RED
		x = parent
	} else {
		if getColor(w.left) {
			if w.right != nil {
				w.right.color = BLACK
			}
			w.color = RED
			t.leftRotate(w)
			w = parent.left
		}
		w.color = parent.color
		parent.color = BLACK
		if w.left != nil {
			w.left.color = BLACK
		}
		t.rightRotate(parent)
		x = t.root
	}
	return x, w
}

func (t *RbTree[K, V]) leftRotate(x *Node[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RbTree[K, V]) rightRotate(x *Node[K, V]) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// findNode finds the node that its key is equal to the passed key, and returns it.
func (t *RbTree[K, V]) findNode(key K) *Node[K, V] {
	x := t.root
	for x != nil {
		if t.keyCmp(key, x.key) < 0 {
			x = x.left
		} else {
			if t.keyCmp(key, x.key) == 0 {
				return x
			}
			x = x.right
		}
	}
	return nil
}

// findNode finds the first node that its key is equal to the passed key, and returns it
func (t *RbTree[K, V]) findFirstNode(key K) *Node[K, V] {
	node := t.FindLowerBoundNode(key)
	if node == nil {
		return nil
	}
	if t.keyCmp(node.key, key) == 0 {
		return node
	}
	return nil
}

// FindLowerBoundNode finds the first node that its key is equal or greater than the passed key, and returns it
func (t *RbTree[K, V]) FindLowerBoundNode(key K) *Node[K, V] {
	return t.findLowerBoundNode(t.root, key)
}

func (t *RbTree[K, V]) findLowerBoundNode(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	if t.keyCmp(key, x.key) <= 0 {
		ret := t.findLowerBoundNode(x.left, key)
		if ret == nil {
			return x
		}
		if t.keyCmp(ret.key, x.key) <= 0 {
			return ret
		}
		return x
	}
	return t.findLowerBoundNode(x.right, key)
}

// FindUpperBoundNode finds the first node that its key is greater than the passed key, and returns it
func (t *RbTree[K, V]) FindUpperBoundNode(key K) *Node[K, V] {
	return t.findUpperBoundNode(t.root, key)
}

func (t *RbTree[K, V]) findUpperBoundNode(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	if t.keyCmp(key, x.key) >= 0 {
		return t.findUpperBoundNode(x.right, key)
	}
	ret := t.findUpperBoundNode(x.left, key)
	if ret == nil {
		return x
	}
	if t.keyCmp(ret.key, x.key) <= 0 {
		return ret
	}
	return x
}

// Traversal traversals elements in the RbTree, it will not stop until to the end of RbTree or the visitor returns false
func (t *RbTree[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	for node := t.First(); node != nil; node = node.Next() {
		if !visitor(node.key, node.value) {
			break
		}
	}
}

// IsRbTree is a function use to test whether t is a RbTree or not
func (t *RbTree[K, V]) IsRbTree() (bool, error) {
	// Properties:
	// 1. Each node is either red or black.
	// 2. The root is black.
	// 3. All leaves (NIL) are black.
	// 4. If a node is red, then both its children are black.
	// 5. Every path from a given node to any of its descendant NIL nodes contains the same number of black nodes.
	_, property, ok := t.test(t.root)
	if !ok {
		return false, fmt.Errorf("violate property %v", property)
	}
	return true, nil
}

func (t *RbTree[K, V]) test(n *Node[K, V]) (int, int, bool) {

	if n == nil { // property 3:
		return 1, 0, true
	}

	if n == t.root && !n.color { // property 2:
		return 1, 2, false
	}
	leftBlackCount, property, ok := t.test(n.left)
	if !ok {
		return leftBlackCount, property, ok
	}
	rightBlackCount, property, ok := t.test(n.right)
	if !ok {
		return rightBlackCount, property, ok
	}

	if rightBlackCount != leftBlackCount { // property 5:
		return leftBlackCount, 5, false
	}
	blackCount := leftBlackCount

	if !n.color {
		if !getColor(n.left) || !getColor(n.right) { // property 4:
			return 0, 4, false
		}
	} else {
		blackCount++
	}

	// if n == t.root {
	// 	fmt.Printf("blackCount:%v \n", blackCount)
	// }
	return blackCount, 0, true
}

// getColor returns the node's color
func getColor[K, V any](n *Node[K, V]) Color {
	if n == nil {
		return BLACK
	}
	return n.color
}
