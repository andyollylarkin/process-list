package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	for range 20 {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			panic(err)
		}
		_ = l
	}

	for range 1000 {
		f, _ := os.Open("/dev/null")
		_ = f
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
}
