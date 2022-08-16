package url_shortener

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	url_shortener "github.com/ydrobot/golang-url_shortener/pkg/api/url_shortener"
)

func (i *URLShortenerServiceImplementation) Create(context.Context, *url_shortener.CreateRequest) (*url_shortener.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
