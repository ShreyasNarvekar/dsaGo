package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	st := new(Stack)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	t.Run("test Push", func(t *testing.T) {
		want := []interface{}{4, 3, 2, 1}
		got := []interface{}{}
		curr := st.Top
		for curr != nil {
			got = append(got, curr.Data)
			curr = curr.Next
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Pop", func(t *testing.T) {
		want := interface{}(4)
		got, _ := st.Pop()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Peek", func(t *testing.T) {
		want := interface{}(3)
		got, _ := st.Peek()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
