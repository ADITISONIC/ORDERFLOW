package services

import (
	"orderflow/dto"
	"orderflow/repositories"
)

func GetMetrics() (*dto.MetricsResponse, error) {

	total, err := repositories.CountOrders()
	if err != nil {
		return nil, err
	}

	pending, err := repositories.CountOrdersByStatus("PENDING")
	if err != nil {
		return nil, err
	}

	processing, err := repositories.CountOrdersByStatus("PROCESSING")
	if err != nil {
		return nil, err
	}

	completed, err := repositories.CountOrdersByStatus("COMPLETED")
	if err != nil {
		return nil, err
	}

	return &dto.MetricsResponse{
		TotalOrders:      total,
		PendingOrders:    pending,
		ProcessingOrders: processing,
		CompletedOrders:  completed,
	}, nil
}