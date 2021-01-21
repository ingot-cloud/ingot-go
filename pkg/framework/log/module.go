package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// InitLogger 初始化日志
func InitLogger(config Config) (func(), error) {

	SetLevel(config.Level)
	SetFormatter(config.Format)

	var file *os.File
	if config.Output != "" {
		switch config.Output {
		case "stdout":
			SetOutput(os.Stdout)
		case "stderr":
			SetOutput(os.Stderr)
		case "file":
			if dir := config.OutputFileDir; dir != "" {
				if fileWriter := rotate(config); fileWriter != nil {
					gin.DefaultWriter = io.MultiWriter(fileWriter, os.Stdout)
					SetOutput(fileWriter)
				}
			}
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}

func rotate(c Config) io.Writer {
	if ok, _ := utils.PathExists(c.OutputFileDir); !ok {
		// directory not exist
		fmt.Println("create log directory")
		err := os.Mkdir(c.OutputFileDir, os.ModePerm)
		if err != nil {
			fmt.Println("mkdir error - ", err)
		}
	}
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s%s%s", c.OutputFileDir, string(os.PathSeparator), "%Y-%m-%d-%H-%M.log"),
		// generate soft link, point to latest log file
		rotatelogs.WithLinkName(c.LogSoftLink),
		// maximum time to save log files
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// time period of log file switching
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println(err)
	}
	return writer
}
