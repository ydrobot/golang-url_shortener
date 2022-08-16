package url_shortener

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	url_shortener "github.com/ydrobot/golang-url_shortener/pkg/api/url_shortener"
)

func (i *URLShortenerServiceImplementation) GetFullURL(context.Context, *url_shortener.GetFullURLRequest) (*url_shortener.GetFullURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullURL not implemented")
}
