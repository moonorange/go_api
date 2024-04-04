package services

import (
	"context"

	"github.com/moonorange/go_api/domain"
)

type TagService interface {
	ListTags(ctx context.Context) ([]*domain.Tag, error)
}
