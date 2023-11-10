package console

import (
	"fmt"
	"sync"

	"github.com/seanmmitchell/ale"
)

var ConsoleWriteLock sync.Mutex

func Log(log *ale.Log) error {
	ConsoleWriteLock.Lock()
	_, err := fmt.Println(log)
	if err != nil {
		return err
	}
	ConsoleWriteLock.Unlock()
	return nil
}
