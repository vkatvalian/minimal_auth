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

func (h *Helper) Fetch(ctx context.Context, _name string) (string, string, string, error) {
    username, email, password, err := h.DB.FetchUsers(ctx, _name)
    if err != nil {}
    return username, email, password, nil
}
