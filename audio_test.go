package ffmpego

import (
	"testing"
)

// go test -run TestAdjustVolume
func TestAdjustVolume(t *testing.T) {
	if err = AdjustVolume(dst+"1.mp3", dst+"1_vol.mp3", 0.3); err != nil {
		t.Error(err)
	}
}

// go test -run TestMixAudios
func TestMixAudios(t *testing.T) {
	if err = MixAudios(dst+"audios_mixed.mp3", dst+"1.mp3",
		dst+"2.mp3", dst+"3.mp3"); err != nil {
		t.Error(err)
	}
}

// go test -run TestCutAudio
func TestCutAudio(t *testing.T) {
	if err = CutAudio(dst+"1.mp3", dst+"1_cut.mp3", 10, 20); err != nil {
		t.Error(err)
	}
}
