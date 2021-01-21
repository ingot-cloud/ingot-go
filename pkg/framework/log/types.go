package log

// Config log 配置
type Config struct {
	Level         int    `yaml:"level"`
	Format        string `yaml:"format"`
	Output        string `yaml:"output"`
	OutputFileDir string `yaml:"outputFileDir"`
	LogSoftLink   string `yaml:"logSoftLink"`
}
