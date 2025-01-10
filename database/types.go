package database

import (
	"errors"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type (
	CHandle struct {
		gorm.Model
		Handle string `gorm:"column:handle;unique"`
		DID    string `gorm:"column:did;unique"`
		DHCode string `gorm:"column:dh_code"`
	}
)

// Makes sure the handle is:
// Does not start or end with a hypen
// 3-63 characters long
// lowercase a-z only
var (
	validationRegex     = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$`)
	errHandleNil        = errors.New("handle is empty")
	errDIDNil           = errors.New("did is empty")
	errValidationFailed = errors.New("handle validation failed")
)

// BeforeSave hook to make sure that the handle/did is not empty,
// and that the handle is valid
func (ch *CHandle) BeforeSave(*gorm.DB) error {
	if ch.Handle == "" {
		return errHandleNil
	}

	if ch.DID == "" {
		return errDIDNil
	}

	if ch.DHCode == "" {
		ch.DHCode = "reserved"
	}

	ch.Handle = strings.ToLower(ch.Handle)
	if !validationRegex.MatchString(ch.Handle) {
		return errValidationFailed
	}

	return nil
}
