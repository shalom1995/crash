package util

import (
	"bytes"
	"os"
	"runtime"
)

/**
 * @parameter:
 * @return: 存在返回false，不存在返回true
 * @Description: 判断指定路径文件是否存在
 * @author: shalom
 * @date: 2020/12/9 14:32
 */
func FileNotExist(path string) bool {
	_, err := os.Lstat(path)
	return os.IsNotExist(err)
}

/**
  * @parameter:
  * @return:
  * @Description: 捕获goroutine出现panic时的错误信息
  * @author: shalom
  * @date: 2020/12/17 9:37
  */
func PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}
