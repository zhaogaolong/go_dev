package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 654321
	}
	fmt.Printf("ret:%d\n", ret)

	user_id, ok := ctx.Value("user_id").(string)
	if !ok {
		fmt.Println("Can't get user_id")
	}
	fmt.Println("user_id:", user_id)
	return
}
func main() {
	// context.Background() 可以理解为一个基类
	ctx := context.WithValue(context.Background(), "trace_id", 123456)

	// 之后再设置可以引用上面的ctx（继承）
	ctx = context.WithValue(ctx, "user_id", "9fd14d74-9fde-48ff-a060-b18d27cdf201")
	process(ctx)
}
