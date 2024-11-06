package test

import (
	"testing"

	"gihub.com/dyxgou/algorithims/sequences"
)

func TestDinamicArray(t *testing.T) {
	da := sequences.NewDinamicArray[int]()

	t.Run("The Dinamic array appends items to itself correctly", func(t *testing.T) {
		ts := NewTestingSuite(t)
		defer ts.Cancel()

		ts.AppendCleanUp(func() {
			da.Reset()
			ts.Cancel()
		})

		da.Append(0)
		da.Append(1)
		da.Append(2)
		da.Append(3)
		da.Append(4)
		da.Append(5)

		for i := 0; i < da.Len(); i++ {
			if da.GetAt(i) != i {
				t.Fail()
			}
		}
	})
}
