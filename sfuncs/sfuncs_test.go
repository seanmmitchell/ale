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

	// Test Overflow
	t.Log("Running Overflow Test...")
	res, err = CenteredText("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 10)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Log("Expected Error Occured on Overflow Case.")
	} else {
		t.Error("Unexpected Success Occured on Overflow Case.")
		t.Fail()
	}

	// Test Underflow (Tecnically a Normal Case)
	t.Log("Running Underflow Test...")
	res, err = CenteredText("test", 6)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Underflow Case.")
		t.Fail()
	} else {
		t.Log("No Error On Underflow Case.")
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

	// Test Overflow
	t.Log("Running Overflow Test...")
	res, err = LeftAlignedText("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 10)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Log("Expected Error Occured on Overflow Case.")
	} else {
		t.Error("Unexpected Success Occured on Overflow Case.")
		t.Fail()
	}

	// Test Underflow (Tecnically a Normal Case)
	t.Log("Running Underflow Test...")
	res, err = LeftAlignedText("test", 6)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Underflow Case.")
		t.Fail()
	} else {
		t.Log("No Error On Underflow Case.")
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

	// Test Overflow
	t.Log("Running Overflow Test...")
	res, err = RightAlignedText("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 10)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Log("Expected Error Occured on Overflow Case.")
	} else {
		t.Error("Unexpected Success Occured on Overflow Case.")
		t.Fail()
	}

	// Test Underflow (Tecnically a Normal Case)
	t.Log("Running Underflow Test...")
	res, err = RightAlignedText("test", 6)
	t.Logf("Result: \"%s\"\n", res)
	if err != nil {
		t.Error("Unexpected Error Occured on Underflow Case.")
		t.Fail()
	} else {
		t.Log("No Error On Underflow Case.")
	}
}

func TestBCenteredText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res := BCenteredText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Overflow
	t.Log("Running Overflow Test...")
	res = BCenteredText("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Underflow (Tecnically a Normal Case)
	t.Log("Running Underflow Test...")
	res = BCenteredText("test", 6)
	t.Logf("Result: \"%s\"\n", res)
}

func TestBLeftAlignedText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res := BLeftAlignedText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Overflow
	t.Log("Running Overflow Test...")
	res = BLeftAlignedText("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Underflow (Tecnically a Normal Case)
	t.Log("Running Underflow Test...")
	res = BLeftAlignedText("test", 6)
	t.Logf("Result: \"%s\"\n", res)
}

func TestBRightAlignedText(t *testing.T) {
	// Test Normal
	t.Log("Running Normal Test...")
	res := BRightAlignedText("normal", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Overflow
	t.Log("Running Overflow Test...")
	res = BRightAlignedText("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 10)
	t.Logf("Result: \"%s\"\n", res)

	// Test Underflow (Tecnically a Normal Case)
	t.Log("Running Underflow Test...")
	res = BRightAlignedText("test", 6)
	t.Logf("Result: \"%s\"\n", res)
}
