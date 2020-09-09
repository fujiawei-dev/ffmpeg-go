package ffmpego

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)



// MixVideoBgm 混合视频原声和背景音乐
func MixVideoBgm(video, bgm, output string) (err error) {
	if err = IsAllExist(video, bgm, filepath.Dir(output)); err != nil {
		return err
	}
	if err = ExtractVideoAudio(video); err != nil {
		return err
	}
	idx := strings.LastIndex(video, ".")
	noaudio := video[:idx] + "_noaudio" + video[idx:]
	audio := video[:idx] + "_audio.m4a"
	mixed := video[:idx] + "_mixed.mp3"
	if err = MixAudios(mixed, audio, bgm); err != nil {
		return err
	}

	defer func() {
		_ = os.Remove(noaudio)
		_ = os.Remove(audio)
		_ = os.Remove(mixed)
	}()

	return MixVideoAudio(noaudio, mixed, output)
}

// MixVideoAudio 混合无声视频和音频，取二者中的最短时长
func MixVideoAudio(video, audio, output string) (err error) {
	if err = IsAllExist(video, audio, filepath.Dir(output)); err != nil {
		return err
	}
	return Command(fmt.Sprintf("-i %s -i %s -shortest "+
		"-vcodec copy -acodec copy %s", video, audio, output))
}

// ConvertFormat 音视频格式转换，以文件名后缀区分格式
func ConvertFormat(input, output string) (err error) {
	if input == output {
		return errors.New("input and output are the same file")
	}
	if err = IsAllExist(input, filepath.Dir(output)); err != nil {
		return err
	}
	return Command(fmt.Sprintf("-i %s %s", input, output))
}

// GetDuration 获取音视频时长
func GetDuration(input string) string {
	if !IsExist(input) {
		return "00:00:00:00"
	}
	buf, _ := CombinedOutput(fmt.Sprintf("-i %s", input))
	reg, _ := regexp.Compile(`Duration: (\d{2}:\d{2}:\d{2}.\d{2})`)
	return reg.FindStringSubmatch(string(buf))[1]
}

// ShowStream 显示流信息
func ShowStream(input string) (err error) {
	if !IsExist(input) {
		return IsNotExistError(input)
	}
	return Command(fmt.Sprintf("-i %s", input))
}
