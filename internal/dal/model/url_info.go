package model

import "time"

type URLInfo struct {
	ID        int64     `db:"url_info_id"`
	UserID    *int64    `db:"user_id"`
	Short     string    `db:"short"`
	URL       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
}
