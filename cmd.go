package ffmpego

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func CombinedInputs(inputs []string) string {
	return "-i " + strings.Join(inputs, " -i ")
}

// CombinedOutput 获取输出结果
func CombinedOutput(s string) ([]byte, error) {
	return exec.Command("ffmpeg",
		strings.Split(s, " ")...).CombinedOutput()
}

// Command 命令执行入口，实时显示输出但不处理输出
func Command(s string) (err error) {

	// 必须分割命令参数，否则无法运行
	args := strings.Split("-y -hide_banner "+s, " ")
	c := exec.Command("ffmpeg", args...)

	stdoutIn, _ := c.StdoutPipe()
	stderrIn, _ := c.StderrPipe()

	// 在命令执行过程中显示输出
	go func() { _, _ = io.Copy(os.Stdout, stdoutIn) }()
	go func() { _, _ = io.Copy(os.Stderr, stderrIn) }()

	if err = c.Run(); err != nil {
		_, _ = io.Copy(os.Stderr, strings.NewReader(c.String()))
	}

	return err
}
