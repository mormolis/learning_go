//go test pack -cover

//go test pack -coverprofile=cover.out
//go tool cover -html=cover.out

package pack

import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Log("Failed to add 1 + 2")
		t.Fail()
		//or t.Error("failed to ....")

	}
}

func TestCanAddMoreThanTwoNumbers(t *testing.T) {
	result := Add(1, 2, 3, 4)

	if result != 10 {
		t.Error("cannot add multiple numbers")
	}
}

func TestFutureFeature(t *testing.T) {
	t.Skip("Not implemented yet")
}
