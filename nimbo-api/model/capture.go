package model

import (
	"time"

	"github.com/google/uuid"
)

type Capture struct {
	ID          uuid.UUID
	File_path   string
	Source      string
	Captured_at time.Time
	Created_at  time.Time
}
