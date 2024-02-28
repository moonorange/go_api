package services

import (
	"context"

	"github.com/moonorange/go_api/models"
)

type TagService interface {
	ListTags(ctx context.Context) ([]*models.Tag, error)
}
