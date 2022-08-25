package url_shortener

import (
	"github.com/ydrobot/golang-url_shortener/internal/domain"
	url_shortener "github.com/ydrobot/golang-url_shortener/pkg/api/url_shortener"
)

type URLShortenerServiceImplementation struct {
	url_shortener.UrlShortenerServiceServer
	service domain.Shortener
}

func NewURLShortenerService(service domain.Shortener) url_shortener.UrlShortenerServiceServer {
	return &URLShortenerServiceImplementation{
		service: service,
	}
}
