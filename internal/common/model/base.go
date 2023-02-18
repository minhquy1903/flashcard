package common

import "time"

type BaseModel struct {
	Id        int32
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
