package ffmpego

import (
	"testing"
)

// go test -run TestCombinedInputs
func TestCombinedInputs(t *testing.T) {
	if ret := CombinedInputs([]string{"1.mp3", "2.mp3", "3.mp3"});
		ret != "-i 1.mp3 -i 2.mp3 -i 3.mp3" {
		t.Error(ret)
	}
}
