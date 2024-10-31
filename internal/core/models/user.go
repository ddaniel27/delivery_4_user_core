package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:",pk,autoincrement"`
	Email         string `bun:",unique"`
	Name          string
	Institution   string
	City          string
	Birthdate     string    `bun:",type:date"`
	CredentialsID int64     `json:"-"`
	CreatedAt     time.Time `bun:",type:timestamp with time zone,default:now()"`
	UpdatedAt     time.Time `bun:",type:timestamp with time zone,default:now()"`
}
