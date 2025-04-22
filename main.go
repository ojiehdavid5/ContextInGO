package main

import (
	"context"
	"fmt"
	"time"
	"net/http"
)

func main() {

	fmt.Println("Hello, World!")
	ctx := context.Background()
	exampleTimeout(ctx)

	exampleWithValues()
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx,cancel:= context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	select {
	case <-time.After(1*time.Second):
		fmt.Println("Api response")
		case <-ctx.Done():
			fmt.Println("oh no my timeout expired",ctx.Err())
			http.Error(w, "Request timed out", http.StatusRequestTimeout)

}}
//USE CASE OF TIME OUT ARE LONG RUNNING API CALLS AND DATABASE QUERIES ALSO NETWORK CALLS,CONCURRENT REQUESTS
func exampleTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second);
	defer cancel()

	done:=make(chan struct{})


	go func() {
		time.Sleep(1*time.Second)
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
	const username string="chuks"
	ctx := context.Background()
	ctxWithValue:=context.WithValue(ctx,username,"chuks")
	if UserId,ok:=ctxWithValue.Value(username).(string);ok{
		fmt.Println("UserId is",UserId)
	

}else{
	fmt.Println("UserId not found")
}
}
