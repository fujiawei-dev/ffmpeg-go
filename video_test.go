package ffmpego

import (
	"fmt"
	"testing"
)

// go test -run TestConcatVideos
func TestConcatVideos(t *testing.T) {
	err = ConcatVideos(dst+"cut_merge.mp4", dst+"long_cut.mp4", dst+"long_cut2.mp4")
	if err != nil {
		t.Error(err)
	}
}

// go test -run TestCrossFadeInOut
func TestCrossFadeInOut(t *testing.T) {
	var inputs []string
	for idx := range make([][]int, 5) {
		inputs = append(inputs, fmt.Sprintf("%s%d.mp4", dst, idx+1))
	}
	err = CrossFadeInOut(dst+"video_cross.mp4", inputs...)
	if err != nil {
		t.Error(err)
	}
}

// go test -run TestFadeInOut
func TestFadeInOut(t *testing.T) {
	err = FadeInOut(video, dst+"video_fade.mp4", 0, 3, 10, 3)
	if err != nil {
		t.Error(err)
	}
}

// go test -run TestCutVideo
func TestCutVideo(t *testing.T) {
	length := 3
	for idx := range make([][]int, 15) {
		if err = CutVideo(dst+"long.mp4", fmt.Sprintf("%spart%d.mp4",
			dst, idx), float32(idx*length), float32(length)); err != nil {
			t.Error(err)
		}
	}
}

// go test -run TestExtractVideo
func TestExtractVideo(t *testing.T) {
	if err = ExtractVideo(dst + "video.mp4"); err != nil {
		t.Error(err)
	}
}

// go test -run TestExtractAudio
func TestExtractAudio(t *testing.T) {
	if err = ExtractAudio(dst + "video.mp4"); err != nil {
		t.Error(err)
	}
}

// go test -run TestExtractVideoAudio
func TestExtractVideoAudio(t *testing.T) {
	if err = ExtractVideoAudio(dst + "video.mp4"); err != nil {
		t.Error(err)
	}
}
