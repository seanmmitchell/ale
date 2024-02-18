package alesyslog

import (
	"log/syslog"

	"github.com/seanmmitchell/ale"
)

type SyslogWriter struct {
	writer *syslog.Writer
}

func New(proto string, remoteAddr string) (*SyslogWriter, error) {
	w, err := syslog.Dial(proto, remoteAddr, syslog.LOG_DEBUG, "ALE")
	if err != nil {
		return nil, err
	}

	return &SyslogWriter{writer: w}, nil
}

func (sw *SyslogWriter) Log(log *ale.Log) error {
	var err1 error
	switch log.Severity {
	case ale.Critical:
		{
			err1 = sw.writer.Crit(log.Message)
		}
	case ale.Error:
		{
			err1 = sw.writer.Err(log.Message)
		}
	case ale.Warning:
		{
			err1 = sw.writer.Warning(log.Message)
		}
	case ale.Info:
		{
			err1 = sw.writer.Info(log.Message)
		}
	case ale.Verbose:
		{
			err1 = sw.writer.Notice(log.Message)
		}
	case ale.Debug:
		{
			err1 = sw.writer.Debug(log.Message)
		}
	}

	if err1 != nil {
		return err1
	}

	return nil
}
