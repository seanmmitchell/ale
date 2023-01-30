package console

import (
	"fmt"
	"sync"

	"github.com/seanmmitchell/ale"
)

var ConsoleWriteLock sync.Mutex

func Log(log *ale.Log) error {
	ConsoleWriteLock.Lock()
	fmt.Println(log.ToString())
	ConsoleWriteLock.Unlock()
	return nil
}
