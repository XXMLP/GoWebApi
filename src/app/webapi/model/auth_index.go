package model

// AuthIndexResponse returns 200.
// swagger:response AuthIndexResponse
type AuthIndexResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data struct {
			UserId string `json:"userId"`
			// Required: true
			Token string `json:"token"`
		} `json:"data"`
		Message string `json:"message"`
	}
}
