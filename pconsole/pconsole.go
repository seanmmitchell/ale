package pconsole

import (
	"fmt"
	"sync"

	"github.com/seanmmitchell/ale/v2/sfuncs"
	"github.com/seanmmitchell/ale/v2/tfuncs"

	"github.com/seanmmitchell/ale/v2"

	"github.com/wzshiming/ctc"
)

type PConsoleCTX struct {
	logLock        sync.Mutex
	sourceWidth    int
	serverityWidth int
}

func New(sourceWidth int, serverityWidth int) (*PConsoleCTX, error) {
	if sourceWidth < 0 || serverityWidth < 0 {
		return nil, fmt.Errorf("console widths can not be negative")
	}

	return &PConsoleCTX{
		logLock:        sync.Mutex{},
		sourceWidth:    sourceWidth,
		serverityWidth: serverityWidth,
	}, nil
}

func (ctx *PConsoleCTX) Log(log *ale.Log) error {
	source := sfuncs.BLeftAlignedText(log.Source, ctx.sourceWidth)
	severity := sfuncs.BCenteredText(ale.TranslateSeverity(log.Severity), ctx.serverityWidth)

	var foregroundColor ctc.Color
	var backgroundColor ctc.Color
	switch log.Severity {
	case ale.Critical:
		{
			backgroundColor = ctc.BackgroundRed
		}
	case ale.Error:
		{
			backgroundColor = ctc.BackgroundBrightRed
			foregroundColor = ctc.ForegroundBlack
		}
	case ale.Warning:
		{
			foregroundColor = ctc.ForegroundBlack
			backgroundColor = ctc.BackgroundYellow
		}
	case ale.Info:
		{
			foregroundColor = ctc.ForegroundBlack
			backgroundColor = ctc.BackgroundWhite
		}
	case ale.Verbose:
		{
			foregroundColor = ctc.ForegroundBlack
			backgroundColor = ctc.BackgroundBrightGreen
		}
	case ale.Debug:
		{
			backgroundColor = ctc.BackgroundBrightBlue
		}
	}

	ctx.logLock.Lock()
	fmt.Printf("%s | %s %s | %s\r\n", tfuncs.TimeToString(log.Time), (fmt.Sprint(foregroundColor|backgroundColor) + severity + fmt.Sprint(ctc.Reset)), source, log.Message)
	ctx.logLock.Unlock()
	return nil
}
