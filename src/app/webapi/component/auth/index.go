package auth

import (
	"app/webapi/store"
	"errors"
	"net/http"
	"time"

	"app/webapi/model"
)

// Index .
// swagger:route GET /v1/auth auth AuthIndex
//
// Get an access token.
//
// Responses:
//   200: AuthIndexResponse
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters UserCheck
	type request struct {
		// in: formData
		// Required: true
		LastName string `json:"last_name" validate:"required"`
		// in: formData
		// Required: true
		Password string `json:"password" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.FormUnmarshal(req, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(req); err != nil {
		return http.StatusBadRequest, err
	}
	// Create the DB store.
	u := store.NewUser(p.DB, p.Q)
	password, err := p.Password.HashString(req.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check for existing item.
	exists, err := u.FindByLastNameAndPassword(u, "last_name", req.LastName, "password", password)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("用户名或密码错误")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("用户名或密码错误")
	}

	item := new(model.UserShowResponseData)

	t, err := p.Token.Generate(item.ID, 8*time.Hour)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp := new(model.AuthIndexResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data.User = item.LastName
	resp.Body.Data.Token = t
	return p.Response.JSON(w, resp.Body)
}
