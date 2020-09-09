package ffmpego

import (
	"errors"
	"fmt"
	"path/filepath"
)

// AdjustVolume 倍数调节音量
func AdjustVolume(input, output string, vol float32) (err error) {
	if err = IsAllExist(input, filepath.Dir(output)); err != nil {
		return err
	}
	if vol < 0 || vol > 5 {
		return errors.New("a reasonable multiple should be 0 ~ 5")
	}
	return Command(fmt.Sprintf("-i %s "+
		"-filter:a volume=%f %s", input, vol, output))
}

// AdjustVolumedB 加减分贝调节音量
func AdjustVolumedB(input, output string, dB float32) (err error) {
	if err = IsAllExist(input, filepath.Dir(output)); err != nil {
		return err
	}
	if dB < -100 || dB > 100 {
		return errors.New("a reasonable multiple should be -100 ~ 100")
	}
	return Command(fmt.Sprintf("-i %s "+
		"-filter:a volume=%fdB %s", input, dB, output))
}

// MixAudios 将多个音频混音
func MixAudios(output string, inputs ...string) (err error) {
	if err = IsAllExist(append(inputs, filepath.Dir(output))...); err != nil {
		return err
	}
	if len(inputs) == 0 {
		return errors.New("input files list is empty")
	}
	return Command(fmt.Sprintf("%s -filter_complex "+
		"amix=inputs=%d:duration=longest %s",
		CombinedInputs(inputs), len(inputs), output))
}

// CutAudio 截取音频片段
func CutAudio(input, output string, start, length float32) (err error) {
	if err = IsAllExist(input, filepath.Dir(output)); err != nil {
		return err
	}
	if length <= 0 {
		return errors.New("length <=0 isn't supported")
	}

	ctrl := "-ss " + ConvertSec2Str(start)
	if length > 0 {
		ctrl += " -t " + ConvertSec2Str(length)
	}
	return Command(fmt.Sprintf("-i %s %s "+
		"-acodec copy %s", input, ctrl, output))
}
