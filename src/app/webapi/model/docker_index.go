package model

// UserIndexResponse returns 200.
// swagger:response DockerInfoResponse
type DockerInfoResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data *DockerInfoResponseData `json:"data"`
	}
}

// DockerInfoResponseData is the docker data.
type DockerInfoResponseData struct {
	ID                string `json:"ID"`
	Containers        int    `json:"first_name"`
	ContainersRunning int    `json:"ContainersRunning"`
	ContainersPaused  int    `json:"ContainersPaused"`
	ContainersStopped int    `json:"ContainersStopped"`
	Images            int    `json:"Images"`
}
