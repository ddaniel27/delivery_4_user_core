package models

type UsersGroups struct {
	ID      int64 `bun:",pk,autoincrement"`
	UserID  int64
	GroupID int64
}
