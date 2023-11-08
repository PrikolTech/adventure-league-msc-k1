package net

import "github.com/gofrs/uuid/v5"

type course struct {
}

func (c *course) Append(userID uuid.UUID, courseID uuid.UUID) error
