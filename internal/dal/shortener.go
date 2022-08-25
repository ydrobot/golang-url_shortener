package dal

import (
	"context"
	"github.com/ydrobot/golang-url_shortener/internal/dal/model"
)

type Shortener interface {
	Add(ctx context.Context, userID *int64, url string, short string) error
	GetByShort(ctx context.Context, short string) (info *model.URLInfo, err error)
	Followed(ctx context.Context, id int64) error
}
