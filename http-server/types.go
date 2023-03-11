package httpserver

type Response[T any] struct {
	Status  int    `json:"status" default:"200" extensions:"x-order=01"`
	Message string `json:"message" default:"Success" extensions:"x-order=02"`
	Data    T      `json:"data" extensions:"x-order=03"`
}
