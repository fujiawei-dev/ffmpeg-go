package ffmpego

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IsFloat 判断字符串是否为浮点数
func IsFloat(s string) bool {
	if len(s) == 0 {
		return false
	}
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsInt 判断字符串是否为整数
func IsInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// ConvertSec2Str 将秒转换成时分秒
func ConvertSec2Str(second float32) string {
	hour := int(second) / 3600
	minute := int(second) % 3600 / 60
	second -= float32(hour*3600 + minute*60)
	return fmt.Sprintf("%02d:%02d:%05.2f", hour, minute, second)
}

// ConvertStr2Sec 将时分秒转换成秒
func ConvertStr2Sec(s string) float32 {
	time := strings.Split(s, ":")
	hour, _ := strconv.ParseFloat(time[0], 32)
	minute, _ := strconv.ParseFloat(time[1], 32)
	second, _ := strconv.ParseFloat(time[2], 32)
	return float32(hour*3600 + minute*60 + second)
}

// TruncateSecond 截断到秒，丢弃小数
func TruncateSecond(s string) string {
	return s[:strings.LastIndex(s, ".")]
}

// IsExist 判断文件/目录是否存在
func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil || os.IsExist(err)
}

// IsNotExistError 返回文件不存在的错误信息
func IsNotExistError(src string) (err error) {
	return errors.New(fmt.Sprintf("%s isn't exist", src))
}

// IsAllExist 判断文件是否都存在
func IsAllExist(ps ...string) (err error) {
	for _, p := range ps {
		if !IsExist(p) {
			return IsNotExistError(p)
		}
	}
	return nil
}
