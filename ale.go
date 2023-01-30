package ale // import "github.com/seanmmitchell/ale"

import (
	"fmt"
	"strings"
	"time"

	"github.com/seanmmitchell/ale/tfuncs"
)

const (
	Critical = iota
	Error    = iota
	Warning  = iota
	Info     = iota
	Verbose  = iota
	Debug    = iota
)

type Log struct {
	Time     time.Time
	Source   string
	Severity int
	Message  string
}

type LogEngine struct {
	parentEngine *LogEngine
	Name         string
	pipeLines    []LogPipeline
	subEngines   []*LogEngine
}

type LogPipeline struct {
	Severity int
	Output   func(log *Log) error
}

// Engine

func CreateLogEngine(name string) *LogEngine {
	return &LogEngine{Name: name}
}

func (le *LogEngine) AddLogPipeline(sev int, out func(log *Log) error) {
	le.pipeLines = append(le.pipeLines, LogPipeline{sev, out})
}

func (le *LogEngine) CreateSubEngine(name string) *LogEngine {
	newEngine := &LogEngine{parentEngine: le, Name: name}
	le.subEngines = append(le.subEngines, newEngine)
	return newEngine
}

func (le *LogEngine) Log(sev int, message string) {
	for _, pipeline := range le.pipeLines {
		if pipeline.Severity >= sev {
			pipeline.Output(&Log{time.Now().UTC(), getEnginePath(le), sev, message})
		}
	}
}

// Log

func (log *Log) ToString() string {
	return fmt.Sprintf("%s | %s | %s | %s", tfuncs.TimeToString(log.Time), log.Source, TranslateSeverity(log.Severity), log.Message)
}

func CenteredText(str string, size int) (string, error) {
	if len(str) > size {
		return "", fmt.Errorf("string too long for size")
	}

	spaces := int(float64(size-len(str)) / 2)
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", size-(spaces+len(str))), nil
}

func LeftAlignedText(str string, size int) (string, error) {
	if len(str) > size {
		return "", fmt.Errorf("string too long for size")
	}

	return str + strings.Repeat(" ", size-len(str)), nil
}

func RightAlignedText(str string, size int) (string, error) {
	if len(str) > size {
		return "", fmt.Errorf("string too long for size")
	}

	return strings.Repeat(" ", size-len(str)) + str, nil
}

// Util

var severityTranslationMap = map[int]string{
	0: "Critical",
	1: "Error",
	2: "Warning",
	3: "Info",
	4: "Verbose",
	5: "Debug",
}

func TranslateSeverity(sev int) string {
	val := severityTranslationMap[sev]
	if val != "" {
		return val
	} else {
		return "UNKNOWN"
	}
}

func getEnginePath(le *LogEngine) string {
	path := le.Name
	pEngine := le.parentEngine
	for {
		if pEngine == nil {
			return path
		} else {
			path = pEngine.Name + "/" + path
			pEngine = pEngine.parentEngine
		}
	}
}
