package model

import "time"

type FollowInfo struct {
	URLInfoID int64     `db:"url_info_id"`
	CreatedAt time.Time `db:"created_at"`
}
