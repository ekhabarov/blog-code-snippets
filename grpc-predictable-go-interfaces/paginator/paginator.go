package paginator

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

type PageLimiter interface {
	GetLimit() int32
	GetPage() int32
}

type Applier interface {
	Apply(squirrel.SelectBuilder) squirrel.SelectBuilder
}

type PageLimitApplier interface {
	PageLimiter
	Applier
}

type pageLimiter struct {
	limit, page int32
}

func (pl *pageLimiter) GetLimit() int32 { return pl.limit }
func (pl *pageLimiter) GetPage() int32  { return pl.page }

func (pl *pageLimiter) Apply(q squirrel.SelectBuilder) squirrel.SelectBuilder {
	if pl.GetLimit() < 1 {
		return q
	}
	return q.
		Limit(uint64(pl.GetLimit())).
		Offset((uint64(pl.GetPage()) - 1) * uint64(pl.GetLimit()))
}

func MustApply(a Applier, q *squirrel.SelectBuilder) error {
	if a == nil {
		return fmt.Errorf("applier is nil")
	}

	if q == nil {
		return fmt.Errorf("select builder is nil, nothing to apply to")
	}

	*q = a.Apply(*q)

	return nil
}

func FromRequest(pl PageLimiter) PageLimitApplier {
	page := pl.GetPage()
	if page < 1 {
		page = 1
	}

	limit := pl.GetLimit()
	if limit < 1 {
		limit = 50
	}

	return &pageLimiter{limit: limit, page: page}
}
