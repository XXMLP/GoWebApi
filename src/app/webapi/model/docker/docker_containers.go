package docker

// DockerContainersResponse returns 200.
// swagger:response DockerContainersResponse
type DockerContainersResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data *DockerContainersData `json:"data"`
	}
}

// DockerContainersData is the docker data.
//type DockerContainersData struct {
//	DockerContainer []DockerContainer
//}
type DockerContainersData []struct {
	ID      string   `json:"Id"`
	Names   []string `json:"Names"`
	Image   string   `json:"Image"`
	ImageID string   `json:"ImageID"`
	Created int64    `json:"Created"`
	Ports   []Ports  `json:"Ports"`
	State   string   `json:"State"`
	Status  string   `json:"Status"`
	Mounts  []Mounts `json:"Mounts"`
}
type Ports struct {
	IP          string `json:"IP"`
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}
type Mounts struct {
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
}
