package ffmpego

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// AddWatermark 添加水印
func AddWatermark(input, output, mark string, x, y int) (err error) {
	if err = IsAllExist(input, filepath.Dir(output), mark); err != nil {
		return err
	}
	if x < 0 || y < 0 {
		return errors.New("x y should >= 0")
	}
	return Command(fmt.Sprintf("%s %s -filter_complex "+
		"[1:v][logo];[0:v][logo]overlay=x=%d:y=%d %s", input, mark, x, y, output))
}

// EmbedSubtitle 内嵌字幕
func EmbedSubtitle(input, output, subtitle string) (err error) {
	if err = IsAllExist(input, filepath.Dir(output), subtitle); err != nil {
		return err
	}
	return Command(fmt.Sprintf("%s "+
		"-vf subtitles=%s %s", input, subtitle, output))
}

// CompressVideo 压缩视频
func CompressVideo(input, output string, scale string) (err error) {
	if err = IsAllExist(input, filepath.Dir(output)); err != nil {
		return err
	}
	if idx := strings.Split(scale, ":"); !(len(idx) == 2 && IsInt(idx[0]) && IsInt(idx[1])) {
		return errors.New("scale format should look like 1280:720")
	}
	return Command(fmt.Sprintf("-i %s -vf scale=%s "+
		"-c:v libx264 -preset veryslow -crf 24 -c:a aac -strict "+
		"-2 -b:a 128k %s", input, scale, output))
}

// CrossFadeInOut 设置交叉淡入淡出
func CrossFadeInOut(output string, inputs ...string) (err error) {
	if err = IsAllExist(append(inputs, filepath.Dir(output))...); err != nil {
		return err
	}
	if len(inputs) < 2 {
		return errors.New("at least two videos are required")
	}
	var (
		filter, overlay string
		tag             = "0"
		pts             float32
	)
	for idx, input := range inputs {
		d := ConvertStr2Sec(GetDuration(input))
		filter += fmt.Sprintf("[%d:v]format=pix_fmts=yuva420p,"+
			"fade=in:st=0:d=1:alpha=1,fade=out:st=%.2f:d=1:alpha=1,"+
			"setpts=PTS-STARTPTS+%.2f/TB[v%d];", idx, d-1, pts, idx)
		if idx < len(inputs)-2 {
			overlay += fmt.Sprintf("[v%s][v%d]overlay[v%s];",
				tag, idx+1, tag+strconv.Itoa(idx+1))
			tag += strconv.Itoa(idx + 1)
		}
		pts += d - 0.5
	}
	overlay += fmt.Sprintf("[v%s][v%d]overlay="+
		"format=yuv420[outv]", tag, len(inputs)-1)
	return Command(fmt.Sprintf("%s -filter_complex %s -vcodec libx264 "+
		"-map [outv] %s", CombinedInputs(inputs), filter+overlay, output))
}

// FadeInOut 设置淡入淡出
func FadeInOut(input, output string, in, ind, out, outd float32) (err error) {
	if err = IsAllExist(input, filepath.Dir(output)); err != nil {
		return err
	}
	var vf string
	if in >= 0 && ind > 0 && out >= 0 && outd > 0 {
		vf = fmt.Sprintf("fade=in:st=%f:d=%f,fade=out:st=%f:d=%f", in, ind, out, outd)
	} else if in >= 0 && ind > 0 {
		vf = fmt.Sprintf("fade=in:st=%f:d=%f", in, ind)
	} else if out >= 0 && outd > 0 {
		vf = fmt.Sprintf("fade=out:st=%f:d=%f", out, outd)
	} else {
		return errors.New("invalid parameters")
	}
	return Command(fmt.Sprintf("-i %s -vf %s %s", input, vf, output))
}

// TransposeVideo 旋转视频
func TransposeVideo(input, output string, angle int) (err error) {
	if err = IsAllExist(input, filepath.Dir(output)); err != nil {
		return err
	}
	var vf string
	switch angle {
	case 90:
		vf = "transpose=1" // 顺时针水平旋转 90 度
	case -90:
		vf = "transpose=2" // 逆时针水平旋转 90 度
	case 180:
		vf = "transpose=1,transpose=1"
	default:
		return errors.New("only support -90, 90 and 180")
	}
	return Command(fmt.Sprintf("-i %s -vf %s %s", input, vf, output))
}

// ConcatVideos 拼接视频
func ConcatVideos(output string, inputs ...string) (err error) {
	if err = IsAllExist(append(inputs, filepath.Dir(output))...); err != nil {
		return err
	}
	if len(inputs) < 2 {
		return errors.New("at least two videos are required")
	}
	tmpfile, _ := ioutil.TempFile("", "inputs*")
	for _, input := range inputs {
		_, _ = tmpfile.WriteString(fmt.Sprintf("file '%s'\n", input))
	}
	return Command(fmt.Sprintf("-f concat -i %s "+
		"-c copy %s", tmpfile.Name(), output))
}

// CutVideo 截取视频片段
func CutVideo(input, output string, start, length float32) (err error) {
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
	return Command(fmt.Sprintf("-i %s %s %s", input, ctrl, output))
}

// ReplaceOriginAudio 替换有声视频中的音频
func ReplaceOriginAudio(video, audio, output string) (err error) {
	if err = IsAllExist(video, audio, filepath.Dir(output)); err != nil {
		return err
	}
	return Command(fmt.Sprintf("-i %s -i %s  -shortest -c:v copy "+
		"-c:a copy -map 0:v:0 -map 1:a:0 %s", video, audio, output))
}

// ExtractVideo 分离视频流
func ExtractVideo(input string) (err error) {
	if !IsExist(input) {
		return IsNotExistError(input)
	}
	idx := strings.LastIndex(input, ".")
	output := input[:idx] + "_noaudio" + input[idx:]
	return Command(fmt.Sprintf("-i %s "+
		"-vcodec copy -an %s", input, output))
}

// ExtractAudio 分离音频流
func ExtractAudio(input string) (err error) {
	if !IsExist(input) {
		return IsNotExistError(input)
	}
	output := input[:strings.LastIndex(input, ".")] + "_audio.m4a"
	return Command(fmt.Sprintf("-i %s "+
		"-acodec copy -vn %s", input, output))
}

// ExtractVideoAudio 分离视频流和音频流
func ExtractVideoAudio(input string) (err error) {
	if err = ExtractVideo(input); err != nil {
		return err
	}
	return ExtractAudio(input)
}
