package structures

import (
	"time"
)

type Temp struct {
	ID          uint
	IDBook      string
	Title       string
	Author      string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
