package log

import (
	"context"
	"fmt"
	"io"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
	"github.com/sirupsen/logrus"
)

// Define logrus alias
var (
	Tracef = logrus.Tracef
	Debugf = logrus.Debugf
	Debug  = logrus.Debug
	Info   = logrus.Info
	Infof  = logrus.Infof
	Warnf  = logrus.Warnf
	Warn   = logrus.Warn
	Errorf = logrus.Errorf
	Error  = logrus.Error
	Fatalf = logrus.Fatalf
	Fatal  = logrus.Fatal
	Panicf = logrus.Panicf
	Panic  = logrus.Panic
	Printf = logrus.Printf
	Print  = logrus.Print
)

// Define key
const (
	TraceIDKey = "traceID"
	UserIDKey  = "userID"
	TagKey     = "tag"
	StackKey   = "stack"
)

type (
	traceIDKey struct{}
	userIDKey  struct{}
	tagKey     struct{}
	stackKey   struct{}
)

// Logger Logrus
type Logger = logrus.Logger

// Entry Logrus entry
type Entry = logrus.Entry

// Hook 定义日志钩子别名
type Hook = logrus.Hook

// Fields Logrus
type Fields = logrus.Fields

// SetLevel 设定日志级别
func SetLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}

// SetFormatter 设定日志输出格式
func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		// ReportCaller 开启后性能降低
		// logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			// CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// 	// _, filename := path.Split(f.File)
			// 	return "", f.File
			// },
		})
	}
}

// WithField 设置字段
func WithField(key string, value any) *Entry {
	return logrus.WithField(key, value)
}

// WithFields 设置字段
func WithFields(fields Fields) *Entry {
	return logrus.WithFields(fields)
}

// SetOutput 设定日志输出
func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

// AddHook 增加日志钩子
func AddHook(hook Hook) {
	logrus.AddHook(hook)
}

// NewTraceIDContext 创建跟踪ID上下文
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// FromTraceIDContext 从上下文中获取跟踪ID
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewUserIDContext 创建用户ID上下文
func NewUserIDContext(ctx context.Context, userID types.ID) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// FromUserIDContext 从上下文中获取用户ID
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewTagContext 创建Tag上下文
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

// FromTagContext 从上下文中获取Tag
func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewStackContext 创建Stack上下文
func NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

// FromStackContext 从上下文中获取Stack
func FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

// WithContext 使用上下文创建Entry
func WithContext(ctx context.Context) *Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := map[string]any{}

	if v := FromTraceIDContext(ctx); v != "" {
		fields[TraceIDKey] = v
	}

	if v := FromUserIDContext(ctx); v != "" {
		fields[UserIDKey] = v
	}

	if v := FromTagContext(ctx); v != "" {
		fields[TagKey] = v
	}

	if v := FromStackContext(ctx); v != nil {
		fields[StackKey] = fmt.Sprintf("%+v", v)
	}

	return logrus.WithContext(ctx).WithFields(fields)
}
