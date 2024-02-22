package console

import (
	"fmt"
	"sync"

	"github.com/seanmmitchell/ale/v2"
)

var ConsoleWriteLock sync.Mutex

func Log(log *ale.Log) error {
	ConsoleWriteLock.Lock()
	fmt.Println(log)
	ConsoleWriteLock.Unlock()
	return nil
}
