package rotate

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	rotateLogs "github.com/pubgo/xlog/internal/writer/rotate/file-rotatelogs"
)

const (
	DefaultRotateMaxAge               = 7 * 24 * time.Hour
	DefaultRotateDuration             = 24 * time.Hour
	DefaultRotatePattern              = ".%Y%m%d%H%M"
	DefaultLogDir                     = "/data/logr"
	DefaultLogSubDir                  = "info"
	DefaultFileMode       os.FileMode = 0755
)

var (
	DefaultFilename = filepath.Base(os.Args[0])
	DefaultLocation = Location()
)

func Location() *time.Location {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.Local
	}
	return loc
}

type Config struct {
	Dir      string
	Sub      string
	Filename string
	Perm     os.FileMode
	Age      time.Duration
	Duration time.Duration
	Pattern  string
	Count    uint
	Loc      *time.Location
}

func NewWriter(opts ...Option) (io.Writer, error) {
	cfg := NewWriterConfig()
	for _, opt := range opts {
		opt.apply(cfg)
	}
	return NewRotateLogger(cfg)
}

func NewWriterConfig() *Config {
	return &Config{
		Dir:      DefaultLogDir,
		Sub:      DefaultLogSubDir,
		Filename: DefaultFilename,
		Perm:     DefaultFileMode,
		Loc:      DefaultLocation,
		Age:      DefaultRotateMaxAge,
		Duration: DefaultRotateDuration,
		Pattern:  DefaultRotatePattern,
	}
}

func NewRotateLogger(config *Config) (*rotateLogs.RotateLogs, error) {

	base := path.Join(config.Dir, config.Sub, config.Filename)
	p, _ := filepath.Split(base)

	// check dir
	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(p, config.Perm); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// create Rotate Logs
	return rotateLogs.New(
		base+config.Pattern,
		rotateLogs.WithLocation(config.Loc),
		rotateLogs.WithRotationCount(config.Count),
		rotateLogs.WithLinkName(base),                // 生成软链，指向最新日志文件
		rotateLogs.WithMaxAge(config.Age),            // 文件最大保存时间
		rotateLogs.WithRotationTime(config.Duration), // 日志切割时间间隔
	)
}
