package rbtree

import (
	"errors"
	"fmt"
)

const (
	RED   = 0
	BLACK = 1
)

type Node struct {
	Left, Right, parent *Node
	color               int
	Key                 int
	Value               interface{}
}

type Tree struct {
	root *Node
	size int
}

//NewTree return a new rbtree
func NewTree() *Tree {
	return &Tree{}
}

//Find find the node and return its value
func (t *Tree) Find(key int) interface{} {
	n := t.findnode(key)
	if n != nil {
		return n.Value
	}
	return nil
}

//FindIt find the node and return it as a iterator
func (t *Tree) FindIt(key int) *Node {
	return t.findnode(key)
}

//Empty check whether the rbtree is empty
func (t *Tree) Empty() bool {
	if t.root == nil {
		return true
	}
	return false
}

func (t *Tree) GetRoot() (*Node, error) {
	if t.Empty() {
		return nil, errors.New("Tree is empty")
	} else {
		return t.root, nil
	}
}

//Iterator create the rbtree's iterator that points to the minmum node
func (t *Tree) Iterator() *Node {
	return minimum(t.root)
}

//Size return the size of the rbtree
func (t *Tree) Size() int {
	return t.size
}

//Clear destroy the rbtree
func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

//Insert insert the key-value pair into the rbtree
func (t *Tree) Insert(key int, value interface{}) {
	x := t.root
	var y *Node

	for x != nil {
		y = x
		if key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z := &Node{parent: y, color: RED, Key: key, Value: value}
	t.size += 1

	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
	t.rb_insert_fixup(z)

}

//Delete delete the node by key
func (t *Tree) Delete(key int) {
	z := t.findnode(key)
	if z == nil {
		return
	}

	var x, y, parent *Node
	y = z
	y_original_color := y.color
	parent = z.parent
	if z.Left == nil {
		x = z.Right
		t.transplant(z, z.Right)
	} else if z.Right == nil {
		x = z.Left
		t.transplant(z, z.Left)
	} else {
		y = minimum(z.Right)
		y_original_color = y.color
		x = y.Right

		if y.parent == z {
			if x == nil {
				parent = y
			} else {
				x.parent = y
			}
		} else {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.parent = y
		}
		t.transplant(z, y)
		y.Left = z.Left
		y.Left.parent = y
		y.color = z.color
	}
	if y_original_color == BLACK {
		t.rb_delete_fixup(x, parent)
	}
	t.size -= 1
}

func (t *Tree) rb_insert_fixup(z *Node) {
	var y *Node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.Left {
			y = z.parent.parent.Right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = BLACK
				z = z.parent.parent
			} else {
				if z == z.parent.Right {
					z = z.parent
					t.left_rotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.right_rotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.Left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.Left {
					z = z.parent
					t.right_rotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.left_rotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *Tree) rb_delete_fixup(x, parent *Node) {
	var w *Node

	for x != t.root && getColor(x) == BLACK {
		if x != nil {
			parent = x.parent
		}
		if x == parent.Left {
			w = parent.Right
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.left_rotate(x.parent)
				w = parent.Right
			}
			if getColor(w.Left) == BLACK && getColor(w.Right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.Right) == BLACK {
					if w.Left != nil {
						w.Left.color = BLACK
					}
					w.color = RED
					t.right_rotate(w)
					w = parent.Right
				}
				w.color = parent.color
				parent.color = BLACK
				if w.Right != nil {
					w.Right.color = BLACK
				}
				t.left_rotate(parent)
				x = t.root
			}
		} else {
			w = parent.Left
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.right_rotate(parent)
				w = parent.Left
			}
			if getColor(w.Left) == BLACK && getColor(w.Right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.Left) == BLACK {
					if w.Right != nil {
						w.Right.color = BLACK
					}
					w.color = RED
					t.left_rotate(w)
					w = parent.Left
				}
				w.color = parent.color
				parent.color = BLACK
				if w.Left != nil {
					w.Left.color = BLACK
				}
				t.right_rotate(parent)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (t *Tree) left_rotate(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.Left {
		x.parent.Left = y
	} else {
		x.parent.Right = y
	}
	y.Left = x
	x.parent = y
}

func (t *Tree) right_rotate(x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != nil {
		y.Right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = x
	} else if x == x.parent.Left {
		x.parent.Left = y
	} else {
		x.parent.Right = y
	}
	y.Right = x
	x.parent = y
}
func (t *Tree) Preorder() {
	fmt.Println("preorder begin!")
	if t.root != nil {
		t.root.preorder()
	}
	fmt.Println("preorder end!")
}

//findnode find the node by key and return it,if not exists return nil
func (t *Tree) findnode(key int) *Node {
	x := t.root
	for x != nil {
		if key < x.Key {
			x = x.Left
		} else {
			if key == x.Key {
				return x
			} else {
				x = x.Right
			}
		}
	}
	return nil
}

//transplant transplant the subtree u and v
func (t *Tree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.Left {
		u.parent.Left = v
	} else {
		u.parent.Right = v
	}
	if v == nil {
		return
	}
	v.parent = u.parent
}

//Next return the node's successor as an iterator
func (n *Node) Next() *Node {
	return successor(n)
}

func (n *Node) preorder() {
	fmt.Printf("(%v %v)", n.Key, n.Value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf("whose parent is %v", n.parent.Key)
	}
	if n.color == RED {
		fmt.Println(" and color RED")
	} else {
		fmt.Println(" and color BLACK")
	}
	if n.Left != nil {
		fmt.Printf("%v's left child is ", n.Key)
		n.Left.preorder()
	}
	if n.Right != nil {
		fmt.Printf("%v's right child is ", n.Key)
		n.Right.preorder()
	}
}

//successor return the successor of the node
func successor(x *Node) *Node {
	if x.Right != nil {
		return minimum(x.Right)
	}
	y := x.parent
	for y != nil && x == y.Right {
		x = y
		y = x.parent
	}
	return y
}

//getColor get color of the node
func getColor(n *Node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

//minimum find the minimum node of subtree n.
func minimum(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

//maximum find the maximum node of subtree n.
func maximum(n *Node) *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}
