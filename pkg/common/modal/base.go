package common

import "time"

type BaseModal struct {
	Id        int32
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
