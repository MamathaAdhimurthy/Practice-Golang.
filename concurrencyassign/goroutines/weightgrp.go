package channelassign

import (
	"fmt"
	"sync"
	"time"
)

func Sleep(wg *sync.WaitGroup, t time.Duration) {
	defer wg.Done()
	time.Sleep(t)
	fmt.Println("finished execution")

}
