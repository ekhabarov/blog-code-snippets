package service

import (
	"log"

	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/paginator"
	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/repo"
)

type service struct {
	repo repo.Repo
}

type Service interface {
	List(page, limit int) ([]repo.Entity, error)
	ListWithApplier(a paginator.PageLimitApplier) ([]repo.Entity, error)
}

func New(r repo.Repo) Service {
	return &service{repo: r}
}

func (s *service) List(page int, limit int) ([]repo.Entity, error) {
	// some additional logic, like logger call
	log.Printf("List called with page: %d and limit: %d\n", page, limit)

	return s.repo.List(page, limit)
}

func (s *service) ListWithApplier(pla paginator.PageLimitApplier) ([]repo.Entity, error) {
	// some additional logic, like logger call
	log.Printf("ListWithApplier called with page: %d and limit: %d\n", pla.GetPage(), pla.GetLimit())

	return s.repo.ListWithApplier(pla)
}
