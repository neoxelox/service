package entity

import (
	"context"

	"github.com/rs/xid"
)

func (m *CookieMethods) Modify(ctx context.Context, cookieID xid.ID,
	title string, category *string) (*Cookie, error) {
	return &Cookie{}, nil
}
