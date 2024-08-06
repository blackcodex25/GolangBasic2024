package main

import (
	"testing"
)

func listToSlice(list *LinkedList) []int {
	var result []int
	current := list.head
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}

func TestLinkedList_Add(t *testing.T) {
	intlist := &LinkedList{}
	intlist.Add(10)
	intlist.Add(20)
	intlist.Add(30)

	expected := []int{10, 20, 30}
	result := listToSlice(intlist)

	if len(result) != len(expected) {
		t.Fatalf("ความยาวที่คาดหวัง %d, ได้รับ %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("คาดหวัง %d ที่ดัชนี %d, ได้รับ %d", v, i, result[i])
		}
	}
}

func TestLinkedList_Print(t *testing.T) {
	tests := []struct {
		name string
		list *LinkedList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Print()
		})
	}
}

func TestLinkedList_Find(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		list *LinkedList
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Find(tt.args.value); got != tt.want {
				t.Errorf("LinkedList.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		list *LinkedList
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Delete(tt.args.value); got != tt.want {
				t.Errorf("LinkedList.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func BenchmarkLinkedList(b *testing.B) {

}
