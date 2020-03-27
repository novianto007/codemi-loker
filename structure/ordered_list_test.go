package structure

import "testing"

func TestPopPush(t *testing.T) {
	val := 3
	list := OrderedList{}
	list.Push(val)
	res, err := list.Pop()
	if err != nil {
		t.Fatal("This opertion should not give an error")
	}
	if res != val {
		t.Fatal("pop give wrong number")
	}
}

func TestOredered(t *testing.T) {
	data := []int{2, 3, 1}
	expect := []int{1, 2, 3}
	list := OrderedList{}
	for _, row := range data {
		list.Push(row)
	}
	for _, row := range expect {
		val, _ := list.Pop()
		if val != row {
			t.Fatalf("ordered was wrong, expect %d, give %d", row, val)
		}
	}
}
