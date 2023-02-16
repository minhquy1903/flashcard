package modal

import (
	common "github.com/minhquy1903/flash-card-api/pkg/common/modal"
)

type User struct {
	common.BaseModal
	Email    string
	Name     string
	Password string
}
