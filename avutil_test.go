package ffmpego

import "testing"

var (
	//dst     = "d:/mpeg/av/"
	dst     = "d:/mpeg/mp/"
	video   = dst + "video.mp4"
	noaudio = dst + "video_noaudio.mp4"
	audio   = dst + "audio.mp3"
	err     error
)

// go test -run TestGetDuration
func TestGetDuration(t *testing.T) {
	println(GetDuration(video))
}

// go test -run TestMixVideoBgm
func TestMixVideoBgm(t *testing.T) {
	if err = MixVideoBgm(video, audio, dst+"av_bgm.mp4"); err != nil {
		t.Error(err)
	}
}

// go test -run TestMixVideoAudio
func TestMixVideoAudio(t *testing.T) {
	if err = MixVideoAudio(noaudio, audio, dst+"av_merge.mp4"); err != nil {
		t.Error(err)
	}
}

// go test -run TestConvertFormat
func TestConvertFormat(t *testing.T) {
	if err = ConvertFormat(video, dst+"video_audio.aac"); err != nil {
		t.Error(err)
	}
	if err = ConvertFormat(video, dst+"video_convert.flv"); err != nil {
		t.Error(err)
	}
}

// go test -run TestShowStream
func TestShowStream(t *testing.T) {
	if err = ShowStream(video); err != nil {
		t.Error(err)
	}
}
