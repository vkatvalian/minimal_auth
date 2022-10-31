package helpers

import (
    "context"
    "github.com/vkatvalian/auth/database"
)

type Helper struct{
    DB *database.Repository
}

func (h *Helper) Insert(ctx context.Context, username, email, password string) error {
    err := h.DB.InsertUsers(ctx, username, email, password)
    if err != nil {}
    return nil
}
