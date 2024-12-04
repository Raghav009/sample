package models

import "github.com/google/uuid"

type Preferences struct {
	UserId          uuid.NullUUID `json:"UserId"`
	PageName        string        `json:"PageName"`
	ViewPreferences string        `json:"ViewPreferences"`
}
