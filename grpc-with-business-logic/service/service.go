package service

import (
	"context"
	"errors"
	"strconv"
)

type (
	OrderItem struct {
		ID    int
		Name  string
		Qty   int
		State string
	}

	Order struct {
		ID      int
		Num     string
		Status  string
		Items   []OrderItem
		Comment string
	}

	// OrderSvc is an interface which represents business logic.
	OrderSvc interface {
		Get(context.Context, int) (*Order, error)
		Add(context.Context, *Order) (int, error)
	}
)

type osvc struct{}

func New() OrderSvc {
	return &osvc{}
}

// Get returns an Order by ID with OrderItems.
//
// Actually this function have to call repo, but for simplicity it returns faked
// data.
func (s *osvc) Get(ctx context.Context, id int) (*Order, error) {
	switch id {
	case 1, 2, 3:
		items := []OrderItem{}

		for i := 0; i < id; i++ {
			items = append(items, OrderItem{
				ID:   (i + 1) * 10,
				Name: "order_item_" + strconv.Itoa(i),
				Qty:  (i + 1) * 2,
			})
		}

		return &Order{
			ID:      id,
			Num:     "order_" + strconv.Itoa(id),
			Status:  "created",
			Items:   items,
			Comment: "none",
		}, nil
	}

	return nil, errors.New("order not found")
}

// Add returns faked added ID.
func (s *osvc) Add(ctx context.Context, o *Order) (int, error) {
	return 5, nil
}
