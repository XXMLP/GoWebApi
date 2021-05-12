package docker

import (
	"app/webapi/model/docker"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
)

// Index .
// swagger:route GET /v1/docker docker DockerContainers
//
// Return all docker.
//
// Security:
//   token:
//
// Responses:
//   200: DockerContainersData
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) Containers(w http.ResponseWriter, r *http.Request) (int, error) {

	type request struct {
		// in: formData
		// Required: true
		IP string `json:"ip" validate:"required"`
		// in: formData
		// Required: true
		Port string `json:"port" validate:"required"`
	}

	// Request validation.
	reqBody := new(request)
	if err := p.Bind.FormUnmarshal(reqBody, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(reqBody); err != nil {
		return http.StatusBadRequest, err
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://" + reqBody.IP + ":" + reqBody.Port + "/containers/json?all=true")
	req.Header.SetMethod("GET")
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)
	if err := HTTPClient.Do(req, resp); err != nil {
		return http.StatusBadRequest, err
	}
	//格式化返回参数
	body := new(docker.DockerContainersData)
	if err := json.Unmarshal(resp.Body(), body); err != nil {
		return http.StatusBadRequest, err
	}

	// Send the response.
	response := new(docker.DockerContainersResponse)
	response.Body.Status = http.StatusText(http.StatusOK)
	response.Body.Data = body
	return p.Response.JSON(w, response.Body)
}
