package user

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"ms-gateway/internal/handlers"
	userDto "ms-gateway/internal/user/dto"
	"ms-gateway/pkg/logging"
	"net/http"
	"time"
)

type handler struct {
	service *Service
	logger  *logging.Logger
}

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

func NewHandler(service *Service, logger *logging.Logger) handlers.Handler {
	return &handler{service: service, logger: logger}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetAll)
	router.GET(userURL, h.GetById)
	router.POST(usersURL, h.Create)
	router.PUT(userURL, h.UpdateOrCreate)
	router.PATCH(userURL, h.Update)
	router.DELETE(userURL, h.Delete)
	h.logger.Infoln("User's handler is initialized.")
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetUsers"))
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetUserById"))
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var payload userDto.CreateUserDto
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(ctx, payload)
	if err != nil {
		h.logger.Errorln(err)
		http.Error(w, "Something was wrong.", http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

func (h *handler) UpdateOrCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("UpdateOrCreate"))
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update"))
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete"))
}
