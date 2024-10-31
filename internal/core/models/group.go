package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Group struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64 `bun:",pk,autoincrement"`
	Name          string
	CreatedAt     time.Time `bun:",type:timestamp with time zone,default:now()"`
	UpdatedAt     time.Time `bun:",type:timestamp with time zone,default:now()"`
}
