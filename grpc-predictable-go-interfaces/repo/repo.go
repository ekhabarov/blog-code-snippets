package repo

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/paginator"
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	// List runs query with plain page/limit arguments.
	List(page, limit int) ([]Entity, error)

	// ListWithApplier runs query with arguments applied by paginator.
	// It accepts minimal required argument paginator.Applier.
	ListWithApplier(a paginator.Applier) ([]Entity, error)
}

type Entity struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type repo struct {
	db *sqlx.DB
}

// New initializes in-memory sqlite database.
func New() (Repo, error) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize db: %w", err)
	}

	create := `CREATE TABLE IF NOT EXISTS entities (
  id INTEGER NOT NULL PRIMARY KEY,
  name TEXT
	);`

	if _, err := db.Exec(create); err != nil {
		return nil, fmt.Errorf("failed to create a table: %w", err)
	}

	for i := 1; i < 100; i++ {
		_, err := db.Exec("INSERT INTO entities VALUES(NULL,?);", fmt.Sprintf("entity_%d", i))
		if err != nil {
			return nil, fmt.Errorf("failed to insert a row %d: %w", i, err)
		}
	}

	return &repo{
		db: db,
	}, nil
}

func runQuery(q squirrel.SelectBuilder, db *sqlx.DB) ([]Entity, error) {
	sql, _, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	fmt.Printf("SQL: %s\n===\n", sql)

	var out []Entity
	if err := db.Select(&out, sql); err != nil {
		return nil, fmt.Errorf("select failed: %w", err)
	}

	return out, nil
}

func (r *repo) List(page int, limit int) ([]Entity, error) {
	q := squirrel.Select("*").From("entities")

	if page > 0 {
		if limit < 1 {
			limit = 5
		}

		q = q.Offset((uint64(page) - 1) * uint64(limit))
	}

	if limit > 0 {
		q = q.Limit(uint64(limit))
	}

	return runQuery(q, r.db)
}

func (r *repo) ListWithApplier(a paginator.Applier) ([]Entity, error) {
	q := squirrel.Select("*").From("entities")

	if err := paginator.MustApply(a, &q); err != nil {
		return nil, err
	}

	return runQuery(q, r.db)
}
