package cmd

import (
	"context"
	"net"

	"github.com/golang/glog"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := net.Listen("tcp", ":8842")
	if err != nil {
		glog.Fatal(err)
	}

	// TODO: register URLShortenerServiceImplementation
}
