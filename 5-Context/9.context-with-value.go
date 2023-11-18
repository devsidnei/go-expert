package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", "12345")
	bookHotel(ctx, "Astor")

}

func bookHotel(ctx context.Context, name string) {
	// Get the request ID from the context
	requestID := ctx.Value("requestID").(string)
	fmt.Println(requestID, name)
}
