package net

import "github.com/gofrs/uuid/v5"

type role struct {
}

func (r *role) Append(userID uuid.UUID, roleID uuid.UUID) error
