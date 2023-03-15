package sfuncs

import (
	"fmt"
	"strings"
)

// Log

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

func BCenteredText(str string, size int) string {
	if len(str) > size {
		return str[:size]
	}

	spaces := int(float64(size-len(str)) / 2)
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", size-(spaces+len(str)))
}

func BLeftAlignedText(str string, size int) string {
	if len(str) > size {
		return str[:size]
	}

	return str + strings.Repeat(" ", size-len(str))
}

func BRightAlignedText(str string, size int) string {
	if len(str) > size {
		return str[:size]
	}

	return strings.Repeat(" ", size-len(str)) + str
}
