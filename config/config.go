package config

import (
	"github.com/BurntSushi/toml"
)

const Path = "/etc/auditlog/auditlog.conf"

var Config = struct {
	Log LogConf
}{
	// default config
	LogConf{
		Path:       "/var/log/auditlog/",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   true,
	},
}

type LogConf struct {
	Path string // Directory path for storing logs, e.g. `/var/log/auditlog`.

	// Following attributes are used for log rotating, which is exactly the
	// same as lumberjack.Logger. See `gopkg.in/natefinch/lumberjack.v2` for
	// more details.
	//
	// The maximum size in MB of single log file before it gets rotated,
	// The zero-value means 100 MB.
	MaxSize int
	// The maximum number of days to retain old log files,
	// The zero-value means never delete old log based on age.
	MaxAge int
	// The maximum number of old log files to retain (MaxAge would still
	// cause them to get deleted).
	// The zero-value means to retain all old log files.
	MaxBackups int
	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.
	// The zero-value means to use UTC time.
	LocalTime bool
	// Compress determines if the rotated log files should be compressed using
	// gzip.
	// The zero-value means not to perform compression.
	Compress bool
}

func Init() error {
	_, err := toml.DecodeFile(Path, &Config)
	return err
}

// Overwrite current config to config file.
func OverWrite() error {
	// TODO
	return nil
}
