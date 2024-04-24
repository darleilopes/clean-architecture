package usecase

import (
	"github.com/darleilopes/clean-architecture/internal/entity"
)

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	var ordersOutputDTO []OrderOutputDTO
	for _, order := range orders {
		ordersOutputDTO = append(ordersOutputDTO, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ListOrdersOutputDTO{Orders: ordersOutputDTO}, nil
}
