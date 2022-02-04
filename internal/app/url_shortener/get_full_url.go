package url_shortener

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	url_shotener_pb "github.com/ydrobot/golang-url_shotener/pkg/url_shotener_pb/api/url_shortener"
)

func (i *URLShortenerServiceImplementation) GetFullURL(context.Context, *url_shotener_pb.GetFullURLRequest) (*url_shotener_pb.GetFullURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullURL not implemented")
}
