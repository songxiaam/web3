package static

import (
	"path"
	"runtime"
)

//| 调用                  | 含义                    |
//| ------------------- | --------------------- |
//| `runtime.Caller(0)` | 获取**当前函数**所在的源文件路径和行号 |
//| `runtime.Caller(1)` | 获取**调用当前函数的函数**的信息    |

func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
