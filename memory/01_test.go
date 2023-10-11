package memory_test

import (
	"gostudy/memory"
	"testing"
)

func TestSetFullName(t *testing.T) {
	u := memory.User{
		FirstName: "David",
		LastName:  "Klassen",
	}

	memory.SetFullName(&u)

	if want := "David Klassen"; u.FullName != want {
		t.Errorf("full name is incorrect: want %q, got %q", want, u.FullName)
	}
}
