package main

import (
	"context"
	"fmt"
	"time"
)

func sendData(ctx context.Context, num int) {
	timer := time.NewTimer(time.Duration(num) * time.Second)

	select {
	case <-ctx.Done():
		fmt.Printf("Процесс #%v отменён\n", num)
		return
	case <-timer.C:
		fmt.Printf("Дaнные процесса #%v успешно отправлены\n", num)
	}

}

func main() {
	parent := context.Background()
	ctx, _ := context.WithDeadline(parent, time.Now().Add(5*time.Second))

	for i := 0; i < 10; i++ {
		go sendData(ctx, i)
	}

	time.Sleep(6 * time.Second)
}
