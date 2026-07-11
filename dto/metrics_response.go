package dto

type MetricsResponse struct {
	TotalOrders      int64 `json:"totalOrders"`
	PendingOrders    int64 `json:"pendingOrders"`
	ProcessingOrders int64 `json:"processingOrders"`
	CompletedOrders  int64 `json:"completedOrders"`
}