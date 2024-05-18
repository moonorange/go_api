package services

import (
	"context"
	"testing"

	"github.com/moonorange/go_api/domain"
	"github.com/moonorange/go_api/infra/mysql_test"
)

type TagService interface {
	ListTags(ctx context.Context) ([]*domain.Tag, error)
}


