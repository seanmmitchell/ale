package sfuncs

import (
	"testing"
)

func TestCenteredText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res, err := CenteredText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Normal Case.")
		t.Fail()
	} else {
		t.Log("No Error On Normal Case.")
	}

	// Test Limit (Tecnically a Normal Case)
	t.Log("Running Limit Test...")
	res, err = CenteredText("test", 4)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Limit Case.")
		t.Fail()
	} else {
		t.Log("No Error On Limit Case.")
	}

	// Test Overflow
	t.Log("Running Overflow Test...")
	res, err = CenteredText("test1234", 4)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Log("Expected Error Occured on Overflow Case.")
	} else {
		t.Error("Unexpected Success Occured on Overflow Case.")
		t.Fail()
	}
}

func TestLeftAlignedText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res, err := LeftAlignedText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Normal Case.")
		t.Fail()
	} else {
		t.Log("No Error On Normal Case.")
	}

	// Test Limit (Tecnically a Normal Case)
	t.Log("Running Limit Test...")
	res, err = LeftAlignedText("test", 4)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Limit Case.")
		t.Fail()
	} else {
		t.Log("No Error On Limit Case.")
	}

	// Test Overflow
	t.Log("Running Overflow Test...")
	res, err = LeftAlignedText("test1234", 4)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Log("Expected Error Occured on Overflow Case.")
	} else {
		t.Error("Unexpected Success Occured on Overflow Case.")
		t.Fail()
	}
}

func TestRightAlignedText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res, err := RightAlignedText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Normal Case.")
		t.Fail()
	} else {
		t.Log("No Error On Normal Case.")
	}

	// Test Limit (Tecnically a Normal Case)
	t.Log("Running Limit Test...")
	res, err = RightAlignedText("test", 4)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Limit Case.")
		t.Fail()
	} else {
		t.Log("No Error On Limit Case.")
	}

	// Test Overflow
	t.Log("Running Overflow Test...")
	res, err = RightAlignedText("test1234", 4)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Log("Expected Error Occured on Overflow Case.")
	} else {
		t.Error("Unexpected Success Occured on Overflow Case.")
		t.Fail()
	}
}

func TestBCenteredText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res := BCenteredText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Limit
	t.Log("Running Limit Test...")
	res = BCenteredText("test", 4)
	t.Logf("Result: \"%s\"\n", res)

	// Test Overflow
	t.Log("Running Overflow Test...")
	res = BCenteredText("test1234", 4)
	t.Logf("Result: \"%s\"\n", res)
}

func TestBLeftAlignedText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res := BLeftAlignedText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Limit
	t.Log("Running Limit Test...")
	res = BLeftAlignedText("test", 4)
	t.Logf("Result: \"%s\"\n", res)

	// Test Overflow
	t.Log("Running Overflow Test...")
	res = BLeftAlignedText("test1234", 4)
	t.Logf("Result: \"%s\"\n", res)
}

func TestBRightAlignedText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res := BRightAlignedText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Limit
	t.Log("Running Limit Test...")
	res = BRightAlignedText("test", 4)
	t.Logf("Result: \"%s\"\n", res)

	// Test Overflow
	t.Log("Running Overflow Test...")
	res = BRightAlignedText("testtest", 4)
	t.Logf("Result: \"%s\"\n", res)
}
