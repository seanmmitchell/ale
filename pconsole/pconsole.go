package pconsole

import (
	"fmt"
	"github.com/seanmmitchell/ale/tfuncs"

	"github.com/seanmmitchell/ale"

	"github.com/wzshiming/ctc"
)

func Log(log *ale.Log) error {
	source, err1 := ale.LeftAlignedText(log.Source, 20)
	if err1 != nil {
		return fmt.Errorf("failed to format source field")
	}
	severity, err2 := ale.CenteredText(ale.TranslateSeverity(log.Severity), 20)
	if err2 != nil {
		return fmt.Errorf("failed to format severity field")
	}

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
