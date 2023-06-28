package ticker

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	checkInerval = time.Millisecond * 500
	timeout      = time.Second * 5
)

func ExecTicker() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	count := 0
	ticker := time.NewTicker(checkInerval)
	checkOk := false
	for {
		count++
		select {
		case <-ticker.C:
			fmt.Printf("tick %d\n", count)
			ok, err := RandomCheck()
			if err != nil {
				fmt.Println("err")
				return
			}
			if ok {
				fmt.Println("ok")
				checkOk = true
			}
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		}
		if checkOk {
			break
		}
	}
}

func RandomCheck() (bool, error) {
	return rand.Intn(10) == 0, nil
}