package users

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"ms-users/internal/handlers"
	"ms-users/internal/middlewares"
	userDto "ms-users/internal/users/dto"
	userDtoOut "ms-users/internal/users/dto/out"
	"ms-users/pkg/logging"
	"net/http"
	"time"
)

type handler struct {
	service   *Service
	validator *validator.Validate
	logger    *logging.Logger
}

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

func NewHandler(service *Service, logger *logging.Logger) handlers.Handler {
	v := validator.New()

	return &handler{service: service, validator: v, logger: logger}
}

func (h *handler) Register(router *httprouter.Router) {
	//router.GET(usersURL, h.GetAll)
	//router.GET(userURL, h.GetById)
	//router.POST(usersURL, h.Create)
	//router.PUT(userURL, h.UpdateOrCreate)
	//router.PATCH(userURL, h.Update)
	//router.DELETE(userURL, h.Delete)
	router.Handle(http.MethodGet, usersURL, middlewares.ErrorHandlingMiddleware(h.GetAll))
	router.Handle(http.MethodGet, userURL, middlewares.ErrorHandlingMiddleware(h.GetById))
	router.Handle(http.MethodPost, usersURL, middlewares.ErrorHandlingMiddleware(h.Create))
	router.Handle(http.MethodPut, userURL, middlewares.ErrorHandlingMiddleware(h.UpdateOrCreate))
	router.Handle(http.MethodPatch, userURL, middlewares.ErrorHandlingMiddleware(h.Update))
	router.Handle(http.MethodDelete, userURL, middlewares.ErrorHandlingMiddleware(h.Delete))
	h.logger.Infoln("User's handler is initialized.")
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var payload userDto.GetUsersDto
	err := json.NewDecoder(r.Body).Decode(&payload) // get data from body. if there is no structure, need to use: body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = h.validator.Struct(payload)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest) // already handle in midlware
		return err
	}

	users, err := h.service.GetAll(ctx, payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json") // should be before WriteHeader
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := params.ByName("id")
	payload := userDto.GetUserByIdDto{Id: id}

	err := h.validator.Struct(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	user, err := h.service.GetById(ctx, payload)
	if err != nil {
		h.logger.Errorln(err)
		return err
	}

	w.Header().Set("Content-Type", "application/json") // should be before WriteHeader
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var payload userDto.CreateUserDto
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		//http.Error(w, "Bad Request", http.StatusBadRequest)
		return err
	}

	err = h.validator.Struct(payload)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	id, err := h.service.Create(ctx, payload)
	if err != nil {
		h.logger.Errorln(err)
		return err
	}

	result := userDtoOut.CreateUserDto{Id: id}

	w.Header().Set("Content-Type", "application/json") // should be before WriteHeader
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) UpdateOrCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Not implemented yet."))
	return nil
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var payload userDto.UpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return err
	}

	payload.Id = params.ByName("id")

	err = h.validator.Struct(payload)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = h.service.Update(ctx, payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := params.ByName("id")
	payload := userDto.DeleteUserDto{Id: id}

	err := h.validator.Struct(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = h.service.Delete(ctx, payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
