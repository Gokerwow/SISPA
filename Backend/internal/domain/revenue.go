package domain

import "time"

type Revenue struct {
	ID         int        `json:"id,omitempty"`
	OmsetDirty int64      `json:"omset_dirty,omitempty"`
	OmsetClean int64      `json:"omset_clean,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}
