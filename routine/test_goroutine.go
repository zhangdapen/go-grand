package routine

import (
	"fmt"
	"time"
)

func ShowMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 100)
	}
}
