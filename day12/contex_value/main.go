package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	res, ok := ctx.Value("trace_id").(int)

	if !ok {
		fmt.Println("cat't get trace_id, err:", ok)
		res = 23892
	}
	fmt.Printf("res:%d\n", res)

	s, _ := ctx.Value("session_id").(string)
	fmt.Printf("s:%s\n", s)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 2379723)
	ctx = context.WithValue(ctx, "session_id", "sessionhfoanofdandfidsnfoiadmflamfd")
	process(ctx)

}
