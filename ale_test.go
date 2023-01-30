package ale_test

import (
	"testing"

	"github.com/seanmmitchell/ale"
	"github.com/seanmmitchell/ale/pconsole"
)

func TestMain(t *testing.T) {
	t.Log("Creating log engine...")
	myLogEngine := ale.CreateLogEngine("ALE")

	t.Log("Adding WARN pipeline to engine...")
	myLogEngine.AddLogPipeline(ale.Warning, pconsole.Log)

	t.Log("Sending logs to engine...")
	myLogEngine.Log(ale.Critical, "Critical Log")
	myLogEngine.Log(ale.Error, "Error Log")
	myLogEngine.Log(ale.Warning, "Warning Log")
	myLogEngine.Log(ale.Info, "Info Log")
	myLogEngine.Log(ale.Verbose, "Verbose Log")
	myLogEngine.Log(ale.Debug, "Debug Log")

	t.Log("Adding subengine to engine...")
	mySubLogEngine := myLogEngine.CreateSubEngine("SUB")

	t.Log("Adding DEBUG pipeline to subengine...")
	mySubLogEngine.AddLogPipeline(ale.Debug, pconsole.Log)

	t.Log("Sending log to subengine...")
	mySubLogEngine.Log(ale.Critical, "Critical Log")
	mySubLogEngine.Log(ale.Error, "Error Log")
	mySubLogEngine.Log(ale.Warning, "Warning Log")
	mySubLogEngine.Log(ale.Info, "Info Log")
	mySubLogEngine.Log(ale.Verbose, "Verbose Log")
	mySubLogEngine.Log(ale.Debug, "Debug Log")
}
