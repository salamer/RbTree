package rbtree

import (
	"testing"
)

type key int

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

func Test_Preorder(t *testing.T) {
	tree := NewTree()

	tree.Insert(int(1), "123")
	tree.Insert(int(3), "234")
	tree.Insert(int(4), "dfa3")
	tree.Insert(int(6), "sd4")
	tree.Insert(int(5), "jcd4")
	tree.Insert(int(2), "bcd4")
	if tree.Size() != 6 {
		t.Error("Error size")
		return
	}
	tree.Preorder()
}

func Test_Find(t *testing.T) {

	tree := NewTree()

	tree.Insert(int(1), "123")
	tree.Insert(int(3), "234")
	tree.Insert(int(4), "dfa3")
	tree.Insert(int(6), "sd4")
	tree.Insert(int(5), "jcd4")
	tree.Insert(int(2), "bcd4")

	n := tree.FindIt(int(4))
	if n.Value != "dfa3" {
		t.Error("Error value")
		return
	}
	n.Value = "bdsf"
	if n.Value != "bdsf" {
		t.Error("Error value modify")
		return
	}
	value := tree.Find(int(5)).(string)
	if value != "jcd4" {
		t.Error("Error value after modifyed other node")
		return
	}
}
func Test_Iterator(t *testing.T) {
	tree := NewTree()

	tree.Insert(int(1), "123")
	tree.Insert(int(3), "234")
	tree.Insert(int(4), "dfa3")
	tree.Insert(int(6), "sd4")
	tree.Insert(int(5), "jcd4")
	tree.Insert(int(2), "bcd4")

	it := tree.Iterator()

	for it != nil {
		it = it.Next()
	}

}

func Test_Delete(t *testing.T) {
	tree := NewTree()

	tree.Insert(int(1), "123")
	tree.Insert(int(3), "234")
	tree.Insert(int(4), "dfa3")
	tree.Insert(int(6), "sd4")
	tree.Insert(int(5), "jcd4")
	tree.Insert(int(2), "bcd4")
	for i := 1; i <= 6; i++ {
		tree.Delete(int(i))
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(int(1), "bcd4")
	tree.Clear()
	tree.Preorder()
	if tree.Find(int(1)) != nil {
		t.Error("Can't clear")
		return
	}
}
