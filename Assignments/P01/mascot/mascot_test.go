package mascot_test

import (
	"testing"

	"example.com/go-demo-1/mascot"
)

func testMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot :<")
	}
}
