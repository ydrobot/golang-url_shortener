package url_shortener

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	url_shortener "github.com/ydrobot/golang-url_shortener/pkg/api/url_shortener"
)

func (i *URLShortenerServiceImplementation) Redirect(ctx context.Context, req *url_shortener.RedirectRequest) (*emptypb.Empty, error) {
	if len(req.Url) == 0 {
		return nil, nil
	}

	url, err := i.service.GetByShortLink(ctx, req.Url)
	if err != nil {
		return nil, err
	}

	if len(url) == 0 {
		return nil, nil
	}

	header := metadata.Pairs("Location", url)
	return nil, grpc.SendHeader(ctx, header)
}
