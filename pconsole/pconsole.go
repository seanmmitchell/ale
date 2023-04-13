package pconsole

import (
	"fmt"

	"github.com/seanmmitchell/ale/sfuncs"
	"github.com/seanmmitchell/ale/tfuncs"

	"github.com/seanmmitchell/ale"

	"github.com/wzshiming/ctc"
)

func Log(log *ale.Log) error {
	source := sfuncs.BLeftAlignedText(log.Source, 20)
	severity := sfuncs.BCenteredText(ale.TranslateSeverity(log.Severity), 20)

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

	fmt.Printf("%s | %s %s | %s\r\n", tfuncs.TimeToString(log.Time), (fmt.Sprint(foregroundColor|backgroundColor) + severity + fmt.Sprint(ctc.Reset)), source, log.Message)
	return nil
}
