package log

import (
	"fmt"
	"ingot/common/utils"
	"ingot/config"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// InitLogger 初始化日志
func InitLogger() (func(), error) {
	cfg := config.CONFIG.Log

	SetLevel(cfg.Level)
	SetFormatter(cfg.Format)

	var file *os.File
	if cfg.Output != "" {
		switch cfg.Output {
		case "stdout":
			SetOutput(os.Stdout)
		case "stderr":
			SetOutput(os.Stderr)
		case "file":
			if dir := cfg.OutputFileDir; dir != "" {
				if fileWriter := rotate(cfg); fileWriter != nil {
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

func rotate(c config.Log) io.Writer {
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
