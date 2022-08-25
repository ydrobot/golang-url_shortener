package domain

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/ydrobot/golang-url_shortener/internal/dal"
)

type Shortener interface {
	Create(ctx context.Context, userID *int64, link string) (string, error)
	GetByShortLink(ctx context.Context, link string) (string, error)
}

type shortener struct {
	repository dal.Shortener
}

func NewShortenerService(repository dal.Shortener) Shortener {
	return &shortener{repository: repository}
}

func (s *shortener) Create(ctx context.Context, userID *int64, link string) (string, error) {
	short := s.generateRandomStr()
	if err := s.repository.Add(ctx, userID, link, short); err != nil {
		return "", errors.New("error added url")
	}
	return short, nil
}

func (s *shortener) GetByShortLink(ctx context.Context, link string) (string, error) {
	urlInfo, err := s.repository.GetByShort(ctx, link)
	if urlInfo == nil || err != nil {
		return "", err
	}

	return urlInfo.URL, s.repository.Followed(ctx, urlInfo.ID)
}

func (s *shortener) generateRandomStr() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	// TODO: config
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
