package ale // import "github.com/seanmmitchell/ale"

import (
	"fmt"
	"time"

	"github.com/seanmmitchell/ale/tfuncs"
)

// Log Types
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

// Log Engine
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
			outputError := pipeline.Output(&Log{time.Now().UTC(), getEnginePath(le), sev, message})

			if outputError != nil {
				// Log Output Failed!
				fmt.Println("ALE | Log Output Failed. Error: " + outputError.Error())
				//return outputError - We dont want to prevent other pipelines from being logged to. Best to just send to console and move on.
			}
		}
	}
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

func (log *Log) String() string {
	return fmt.Sprintf("%s | %s | %s | %s", tfuncs.TimeToString(log.Time), log.Source, TranslateSeverity(log.Severity), log.Message)
}
