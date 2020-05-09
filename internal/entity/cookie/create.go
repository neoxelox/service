package entity

import "context"

func (m *CookieMethods) Create(ctx context.Context, title string,
	category *string) (*Cookie, error) {
	return &Cookie{}, nil
}
