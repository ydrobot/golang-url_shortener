package url_shortener

import url_shortener "github.com/ydrobot/golang-url_shortener/pkg/url_shotener_pb/api/url_shortener"

type URLShortenerServiceImplementation struct {
	url_shortener.UrlShortenerServiceServer
}

func NewURLShortenerService() url_shortener.UrlShortenerServiceServer {
	return &URLShortenerServiceImplementation{}
}
