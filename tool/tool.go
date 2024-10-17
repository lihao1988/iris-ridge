package tool

import (
	"os"

	"net/http"

	"github.com/kataras/iris/v12"
)

// GetExecFilePath get the exec_file_path
func GetExecFilePath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path
}

// WriteFileContent 文件内容写入到浏览器.
func WriteFileContent(ctx iris.Context, filename string) {
	var b []byte

	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	if _, err = ctx.WriteString(string(b)); err != nil {
		ctx.StatusCode(http.StatusNotFound)
	}
}
