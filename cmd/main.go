package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/ydrobot/golang-url_shortener/internal/app"
	"github.com/ydrobot/golang-url_shortener/internal/app/url_shortener"
	desc "github.com/ydrobot/golang-url_shortener/pkg/api/url_shortener"
)

func main() {
	// Interceptors
	grpcOpts := app.GrpcInterceptor()
	httpOpts := app.HttpInterceptor()

	go runGRPC(grpcOpts)
	runHTTP(httpOpts)
}
func runGRPC(grpcOpts grpc.ServerOption) {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(grpcOpts)
	desc.RegisterUrlShortenerServiceServer(grpcServer, url_shortener.NewURLShortenerService())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}

func runHTTP(httpOpts runtime.ServeMuxOption) {
	mux := runtime.NewServeMux(httpOpts, app.HttpMarshalerOption())
	err := desc.RegisterUrlShortenerServiceHandlerServer(context.Background(), mux, url_shortener.NewURLShortenerService())
	if err != nil {
		log.Println("cannot register this service")
	}

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
