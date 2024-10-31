package models

import "github.com/uptrace/bun"

type Credential struct {
	bun.BaseModel `bun:"table:credentials,alias:c"`
	ID            int64  `bun:",pk,autoincrement"`
	Email         string `bun:",unique"`
	Password      string
}
