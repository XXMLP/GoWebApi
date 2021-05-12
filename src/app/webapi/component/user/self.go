package user

import (
	"errors"
	"net/http"

	"app/webapi/model"
	"app/webapi/pkg/structcopy"
	"app/webapi/store"
)

// *****************************************************************************
// Read
// *****************************************************************************

// userSelf .
// swagger:route GET /v1/userSelf user UserShow
//
// Return one user.
//
// Security:
//   token:
//
// Responses:
//   200: UserShowResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *Endpoint) ShowSelf(w http.ResponseWriter, r *http.Request) (int, error) {

	// Create the DB store.
	u := store.NewUser(p.DB, p.Q)
	//解析token
	bearer := r.Header.Get("Authorization")
	userId, _ := ParseToken(bearer[7:])
	//Get an item by ID.
	exists, err := u.FindOneByID(u, userId)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("user not found")
	}

	// Copy the items to the JSON model.
	arr := make([]model.UserShowResponseData, 0)
	item := new(model.UserShowResponseData)
	err = structcopy.ByTag(u, "db", item, "json")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	arr = append(arr, *item)

	// Send the response.
	resp := new(model.UserShowResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data = arr
	return p.Response.JSON(w, resp.Body)
}
