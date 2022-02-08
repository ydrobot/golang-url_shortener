package url_shortener

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	url_shortener "github.com/ydrobot/golang-url_shortener/pkg/url_shotener_pb/api/url_shortener"
)

func (i *URLShortenerServiceImplementation) GetList(context.Context, *url_shortener.GetListRequest) (*url_shortener.GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
