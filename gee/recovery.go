package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//recovery中间件，错误处理机制

// 获取panic堆栈信息
func trace(message string) string {
	var pcs [32]uintptr
	//Callers的第一个参数的3代币哦从第3个开始获取，第0个是Callers 本身，
	//第 1 个是上一层 trace，第 2 个是再上一层的 defer func
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			//当发生panic时就会执行当前协程的defer中的recovery
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
