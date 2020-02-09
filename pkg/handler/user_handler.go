package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rahul-golang/mongocrud/log"

	"github.com/gorilla/mux"
	"github.com/rahul-golang/mongocrud/pkg/models"
	"github.com/rahul-golang/mongocrud/pkg/service"
)

//UserHandlersImpl for handler Functions
type UserHandlersImpl struct {
	userSvc service.UserService
}

// CreateUser handler Function
func (userHandlersImpl UserHandlersImpl) CreateUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user := models.User{}
	log.Logger(ctx).Info("in request")
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		HTTPError(w, http.StatusBadRequest)
		return
	}

	resp, err := userHandlersImpl.userSvc.CreateUser(ctx, user)
	if err != nil {
		CustomHTTPError(w, http.StatusConflict, err.Error())
		return
	}

	endpointResp, err := json.MarshalIndent(resp, " ", " ")
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	w.Write(endpointResp)
}

// GetUsers handler Function
func (userHandlersImpl UserHandlersImpl) GetUsers(w http.ResponseWriter, req *http.Request) {
	var strpage, strlimit string

	keys := req.URL.Query()
	pages, ok := keys["page"]
	if ok {
		strpage = pages[0]
	} else {
		strpage = "0"
	}
	limits, ok := keys["limit"]
	if ok {
		strlimit = limits[0]
	} else {
		strlimit = "10"
	}

	page, err := strconv.Atoi(strpage)

	if err != nil {
		CustomHTTPError(w, http.StatusBadRequest, "page query parameter should be an int")
		return
	}

	limit, err := strconv.Atoi(strlimit)

	if err != nil {
		CustomHTTPError(w, http.StatusBadRequest, "page limit parameter should be an int")
		return
	}

	ctx := req.Context()
	resp, err := userHandlersImpl.userSvc.GetAllUser(ctx, page, limit)

	if err != nil {

		CustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(resp) <= 0 {
		HTTPError(w, http.StatusNotFound)
		return
	}

	endpointResp, err := json.MarshalIndent(resp, "  ", " ")

	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	w.Write(endpointResp)

}

// GetUserByID handler Function
func (userHandlersImpl UserHandlersImpl) GetUserByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	resp, err := userHandlersImpl.userSvc.GetUser(ctx, id)

	if err != nil {
		CustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	endpointResp, err := json.MarshalIndent(resp, "  ", " ")
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	w.Write(endpointResp)

}

// UpdateUser handler Function
func (userHandlersImpl UserHandlersImpl) UpdateUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]

	user := models.User{}
	log.Logger(ctx).Info("in request")
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	objectID, err := GetHexID(ctx, id)
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	user.ID = objectID

	resp, err := userHandlersImpl.userSvc.UpdateUser(ctx, user)
	if err != nil {
		CustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	endpointResp, err := json.Marshal(resp)
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	w.Write(endpointResp)
}

// DeleteUser handler Function
func (userHandlersImpl UserHandlersImpl) DeleteUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]

	resp, err := userHandlersImpl.userSvc.DeleteUser(ctx, id)
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	endpointResp, err := json.MarshalIndent(resp, "  ", " ")
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}
	w.Write(endpointResp)

}

//NewUserHandlerImpl inits dependancies for graphQL and Handlers
func NewUserHandlerImpl(userService service.UserService) *UserHandlersImpl {
	return &UserHandlersImpl{userSvc: userService}

}
