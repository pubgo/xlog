// +build !windows

package xlog_journal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"go.uber.org/zap/zapcore"
)

// NewJournalWriter wraps "io.Writer" to redirect log output
// to the local systemd journal. If journald send fails, it fails
// back to writing to the original writer.
// The decode overhead is only <30µs per write.
// Reference: https://github.com/coreos/pkg/blob/master/capnslog/journald_formatter.go
func NewJournalWriter(wr io.Writer) (io.Writer, error) {
	return &journalWriter{Writer: wr}, DialJournal()
}

type journalWriter struct {
	io.Writer
}

// WARN: assume that etcd uses default field names in zap encoder config
// make sure to keep this up-to-date!
type logLine struct {
	Level  string `json:"level"`
	Caller string `json:"caller"`
}

func (w *journalWriter) Write(p []byte) (int, error) {
	line := &logLine{}
	if err := json.NewDecoder(bytes.NewReader(p)).Decode(line); err != nil {
		return 0, err
	}

	var pri Priority
	switch line.Level {
	case zapcore.DebugLevel.String():
		pri = PriDebug
	case zapcore.InfoLevel.String():
		pri = PriInfo

	case zapcore.WarnLevel.String():
		pri = PriWarning
	case zapcore.ErrorLevel.String():
		pri = PriErr

	case zapcore.DPanicLevel.String():
		pri = PriCrit
	case zapcore.PanicLevel.String():
		pri = PriCrit
	case zapcore.FatalLevel.String():
		pri = PriCrit

	default:
		panic(fmt.Errorf("unknown log level: %q", line.Level))
	}

	err := Send(string(p), pri, map[string]string{
		"PACKAGE":           filepath.Dir(line.Caller),
		"SYSLOG_IDENTIFIER": filepath.Base(os.Args[0]),
	})
	if err != nil {
		// "journal" also falls back to stderr
		// "fmt.Fprintln(os.Stderr, s)"
		return w.Writer.Write(p)
	}
	return 0, nil
}

// DialJournal returns no error if the process can dial journal socket.
// Returns an error if dial failed, whichi indicates journald is not available
// (e.g. run embedded etcd as docker daemon).
// Reference: https://github.com/coreos/go-systemd/blob/master/journal/journal.go.
func DialJournal() error {
	conn, err := net.Dial("unixgram", "/run/systemd/journal/socket")
	if conn != nil {
		defer conn.Close()
	}
	return err
}
