package url_shortener

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	url_shotener_pb "github.com/ydrobot/golang-url_shotener/pkg/url_shotener_pb/api/url_shortener"
)

func (i *URLShortenerServiceImplementation) Create(context.Context, *url_shotener_pb.CreateRequest) (*url_shotener_pb.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
