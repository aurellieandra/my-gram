package pkg

type Response struct {
	Status  uint64      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AuthResponse struct {
	Status  uint64      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}
