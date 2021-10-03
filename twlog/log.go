package twlog

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/twwch/gin-sdk/constant"
	"path"
	"time"
)

type LogConf struct {
	Path         string        `json:"path" mapstructure:"path" toml:"path"`
	ApiLog       string        `jsos:"api" mapstructure:"api" toml:"api"`
	PanicLog     string        `jsos:"panic" mapstructure:"panic" toml:"panic"`
	DebugLog     string        `json:"debug" mapstructure:"debug" toml:"debug"`
	FatalLog     string        `json:"fatal" mapstructure:"fatal" toml:"fatal"`
	WarnLog      string        `json:"warn"  mapstructure:"warn" toml:"warn"`
	TraceLog     string        `json:"trace"  mapstructure:"trace" toml:"trace"`
	ErrorLog     string        `json:"error" mapstructure:"error" toml:"error"`
	MaxAge       time.Duration `json:"max_age" mapstructure:"max" toml:"max_age"`
	RotationTime time.Duration `json:"rotation_time" mapstructure:"rotation_time" toml:"rotation_time"`
}

func NewLog(l *LogConf) {
	configLocalFilesystemLogger(l)
}

// config logrus twlog to local filesystem, with file rotation
func configLocalFilesystemLogger(l *LogConf) {
	witerMap := lfshook.WriterMap{}
	logs := []string{l.ApiLog, l.DebugLog, l.ErrorLog, l.WarnLog, l.PanicLog, l.FatalLog, l.TraceLog}
	logsMap := map[string]log.Level{
		l.ApiLog:   log.InfoLevel,
		l.DebugLog: log.DebugLevel,
		l.ErrorLog: log.ErrorLevel,
		l.WarnLog:  log.WarnLevel,
		l.FatalLog: log.FatalLevel,
		l.PanicLog: log.PanicLevel,
		l.TraceLog: log.TraceLevel,
	}
	if l.Path != "" {
		for _, logItem := range logs {
			if logItem != "" {
				if l.MaxAge == 0 {
					l.MaxAge = time.Hour * 15
				}
				if l.RotationTime == 0 {
					l.RotationTime = time.Hour * 24
				}
				file := path.Join(l.Path, logItem)
				apiWaiter, err := rotatelogs.New(
					file+".%Y%m%d",
					rotatelogs.WithLinkName(file),               // 生成软链，指向最新日志文件
					rotatelogs.WithMaxAge(l.MaxAge),             // 文件最大保存时间
					rotatelogs.WithRotationTime(l.RotationTime), // 日志切割时间间隔
				)
				if err != nil {
					log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
				}
				witerMap[logsMap[logItem]] = apiWaiter
			}
		}
	}
	lfHook := lfshook.NewHook(witerMap, &log.JSONFormatter{TimestampFormat: constant.FullDateTimeFormat, DisableHTMLEscape: true})
	log.AddHook(lfHook)
}
