package ale_test

import (
	"testing"

	"github.com/seanmmitchell/ale/v2"
	"github.com/seanmmitchell/ale/v2/alesyslog"
	"github.com/seanmmitchell/ale/v2/pconsole"
)

func TestMain(t *testing.T) {
	t.Log("====== Starting Core Test")
	t.Log("Creating log engine...")
	myLogEngine := ale.CreateLogEngine("ALE")

	t.Log("Creating pconsole CTX...")
	pCTX, err1 := pconsole.New(30, 30)
	if err1 != nil {
		t.Log("Failed to created pConsole CTX.")
		t.Fail()
	}

	t.Log("Adding WARN pipeline to engine...")
	myLogEngine.AddLogPipeline(ale.Warning, pCTX.Log)

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
	mySubLogEngine.AddLogPipeline(ale.Debug, pCTX.Log)

	t.Log("Sending logs to subengine...")
	mySubLogEngine.Log(ale.Critical, "Critical Log")
	mySubLogEngine.Log(ale.Error, "Error Log")
	mySubLogEngine.Log(ale.Warning, "Warning Log")
	mySubLogEngine.Log(ale.Info, "Info Log")
	mySubLogEngine.Log(ale.Verbose, "Verbose Log")
	mySubLogEngine.Log(ale.Debug, "Debug Log")

	t.Log("====== Starting Module Tests ")

	t.Log("=== Starting Syslog Module Tests ")
	t.Log("Creating log engine...")
	sysLogEngine := ale.CreateLogEngine("ALE > Syslog")

	t.Log("Creating Syslog Output...")
	as, err1 := alesyslog.New("udp", "10.16.6.4:12002")

	t.Log("Adding Syslog Pipeline...")
	sysLogEngine.AddLogPipeline(ale.Debug, as.Log)

	t.Log("Sending Syslog...")
	sysLogEngine.Log(ale.Debug, "Test for Syslog Module of ALE (https://github.com/seanmmitchell/ale)")
}
