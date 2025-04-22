package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello, World!")
	ctx := context.Background()
	exampleTimeout(ctx)
}

func exampleTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second);
	defer cancel()

	done:=make(chan struct{})


	go func() {
		time.Sleep(6*time.Second)
		close(done)
	}()
select {
case <-done:
	fmt.Println("call the api")
	case <-ctxWithTimeout.Done():
		fmt.Println("oh no my timeout expired",ctxWithTimeout.Err())
		//do some logic here
}
}


func exampleWithValues(){
	
}
