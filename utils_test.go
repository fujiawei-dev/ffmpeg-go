package ffmpego

import (
	"testing"
)

// go test -run TestTruncateSecond
func TestTruncateSecond(t *testing.T) {
	if ret := TruncateSecond("01:02:03.04"); ret != "01:02:03" {
		t.Error(ret)
	}
}

// go test -run TestConvertStr2Sec
func TestConvertStr2Sec(t *testing.T) {
	if ret := ConvertStr2Sec("01:02:03.04"); ret != 3723.04 {
		t.Error(ret)
	}
}

// go test -run TestConvertSec2Str
func TestConvertSec2Str(t *testing.T) {
	if ret := ConvertSec2Str(3723.04); ret != "01:02:03.04" {
		t.Error(ret)
	}
}
