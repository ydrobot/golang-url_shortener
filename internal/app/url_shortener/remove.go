package url_shortener

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	url_shortener "github.com/ydrobot/golang-url_shortener/pkg/api/url_shortener"
)

func (i *URLShortenerServiceImplementation) Remove(context.Context, *url_shortener.RemoveRequest) (*url_shortener.RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
