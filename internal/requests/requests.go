package requests

type UpdateMetricRequest struct {
	ID    string  `json:"id" binding:"required"`
	MType string  `json:"type" binding:"required,oneof=counter gauge"`
	Delta *int64   `json:"delta,omitempty" binding:"omitempty,required_if=MType counter"`
	Value *float64 `json:"value,omitempty" binding:"omitempty,required_if=MType gauge"`
}

type GetMetricValueRequest struct {
	ID    string `json:"id" binding:"required"`
	MType string `json:"type,omitempty" binding:"oneof=counter gauge,omitempty"`
}
