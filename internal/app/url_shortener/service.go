package url_shortener

import url_shotener_pb "github.com/ydrobot/golang-url_shotener/pkg/url_shotener_pb/api/url_shortener"

type URLShortenerServiceImplementation struct {
	url_shotener_pb.UrlShortenerServiceServer
}

func NewURLShortenerService() url_shotener_pb.UrlShortenerServiceServer {
	return &URLShortenerServiceImplementation{}
}
